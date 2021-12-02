package orka

import (
	"context"
	"fmt"
	"net/http"
)

type Authentication struct {
	BearerToken string
	LicenseKey  string
}

// NewOrkaClient creates a client for the Orka API using the provided
// authentication credentials.
//
// You do not have to set both the BearerToken and LicenseKey fields, only
// the ones needed by the API calls you plan on using.
//
// For example, if you only plan on performing non-administrative requests,
// you only need to set the BearerToken field.
func NewOrkaClient(serverURL string, auth Authentication, options ...ClientOption) (*ClientWithResponses, error) {
	finalOptions := append(
		options,
		WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			if auth.BearerToken != "" {
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", auth.BearerToken))
			}
			if auth.LicenseKey != "" {
				req.Header.Set("orka-licensekey", auth.LicenseKey)
			}
			return nil
		}),
	)
	client, err := NewClientWithResponses(serverURL, finalOptions...)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// ExtractError parses an HTTP response body from a failed API response and returns
// the corresponding error.
func ExtractError(body []byte) error {
	if body == nil {
		return fmt.Errorf("orka API failed but unexpectedly did not return an error message")
	}
	return fmt.Errorf("orka API returned an error: %+v", string(body))
}
