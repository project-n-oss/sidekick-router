package app

import (
	"fmt"
	"net/http"

	"github.com/project-n-oss/interchange/app/aws"
)

// DoRequest makes a request to the cloud platform
// Does a request to the source bucket and if it returns 404, tries the crunched bucket
// Returns the response and a boolean indicating if the response is from the crunched bucket
func (sess *Session) DoRequest(req *http.Request) (*http.Response, bool, error) {
	switch sess.app.cfg.CloudPlatform {
	case AwsCloudPlatform:
		return sess.DoAwsRequest(req)
	default:
		return nil, false, fmt.Errorf("CloudPlatform %s not supported", sess.app.cfg.CloudPlatform)
	}
}

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

	crunchedBucket := aws.SourceBucket{
		Bucket: fmt.Sprintf("project-n-%s", sourceBucket.Bucket),
		Region: sourceBucket.Region,
		Style:  sourceBucket.Style,
	}

	crunchedRequest, err := aws.NewRequest(sess.Context(), sess.Logger(), req, crunchedBucket)
	if err != nil {
		return nil, false, fmt.Errorf("failed to make aws request: %w", err)
	}

	resp, err := http.DefaultClient.Do(cloudRequest)

	statusCode := -1
	if resp != nil {
		statusCode = resp.StatusCode
	}

	// check crunched version of file
	if statusCode == 404 {
		sess.logger.Info("Trying crunched version of file")
		resp, err := http.DefaultClient.Do(crunchedRequest)
		return resp, true, err
	}

	return resp, false, err
}
