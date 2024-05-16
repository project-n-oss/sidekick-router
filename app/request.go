package app

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/project-n-oss/sidekick-router/app/aws"
)

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

	statusCode := -1
	if resp != nil {
		statusCode = resp.StatusCode
	}

	if statusCode == 404 && !sess.app.cfg.NoCrunchFailover {
		crunchedBucket := aws.SourceBucket{
			Bucket: fmt.Sprintf("project-n-%s", sourceBucket.Bucket),
			Region: sourceBucket.Region,
			Style:  sourceBucket.Style,
		}

		if sourceBucket.Style == aws.PathStyle {
			req.URL.Path = strings.Replace(req.URL.Path, sourceBucket.Bucket, crunchedBucket.Bucket, 1)
		}

		crunchedRequest, err := aws.NewRequest(sess.Context(), sess.Logger(), req, crunchedBucket)
		if err != nil {
			return nil, false, fmt.Errorf("failed to make aws request: %w", err)
		}

		resp, err := http.DefaultClient.Do(crunchedRequest)
		return resp, true, err
	}

	return resp, false, err
}
