package app

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/project-n-oss/sidekick-router/app/aws"
)

func statusCodeIs2xx(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}

// DoRequest makes a request to the cloud platform
// Does a request to the source bucket and if it returns 404, tries the crunched bucket
// Returns the response and a boolean indicating if the response is from the crunched bucket
func (sess *Session) DoRequest(req *http.Request) (*http.Response, bool, error) {
	switch sess.app.cfg.CloudPlatform {
	case AwsCloudPlatform.String():
		return sess.DoAwsRequest(req)
	default:
		return nil, false, fmt.Errorf("CloudPlatform %s not supported", sess.app.cfg.CloudPlatform)
	}
}

const crunchFileFoundErr = "Src file not found, but crunched file found."

// DoAwsRequest makes a request to AWS
// Does a request to the source bucket and if it returns 404, tries the crunched bucket
// Returns the response and a boolean indicating if the response is from the crunched bucket
func (sess *Session) DoAwsRequest(req *http.Request) (*http.Response, bool, error) {
	sourceBucket, err := aws.ExtractSourceBucket(req)
	if err != nil {
		return nil, false, fmt.Errorf("failed to extract source bucket from request: %w", err)
	}

	cloudRequest, err := aws.NewRequest(sess.Context(), sess.Logger(), req, sourceBucket)
	if err != nil {
		return nil, false, fmt.Errorf("failed to make aws request: %w", err)
	}

	resp, err := http.DefaultClient.Do(cloudRequest)
	if err != nil {
		return nil, false, fmt.Errorf("failed to do aws request: %w", err)
	}

	statusCode := -1
	if resp != nil {
		statusCode = resp.StatusCode
	}

	if statusCode == 404 && !sess.app.cfg.NoCrunchErr {
		path := req.URL.Path
		newPath := path
		print("path: ", path, "\n")

		crunchedRequest, err := aws.NewRequest(sess.Context(), sess.Logger(), req, sourceBucket, aws.WithPath(newPath))
		if err != nil {
			return nil, false, fmt.Errorf("failed to make aws request: %w", err)
		}

		resp, err := http.DefaultClient.Do(crunchedRequest)
		if err != nil {
			return nil, false, fmt.Errorf("failed to do crunched aws request: %w", err)
		}
		crunchedStatusCode := -1
		if resp != nil {
			crunchedStatusCode = resp.StatusCode
		}

		// return 500 to client if there is a crunch version of the file
		if statusCodeIs2xx(crunchedStatusCode) {
			resp.StatusCode = 500
			resp.Body = io.NopCloser(bytes.NewReader([]byte(crunchFileFoundErr)))
		}

		return resp, true, err
	}

	return resp, false, err
}
