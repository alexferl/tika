package tika

import (
	"errors"
	"fmt"
	"net/http"
)

// DetectResource represents the structure of our resource
type DetectResource struct {
	client   *Client
	endpoint string
}

// Detect is the entry point for interacting with the Detect resource
func (c *Client) Detect() *DetectResource {
	return &DetectResource{client: c,
		endpoint: c.Url + "/detect/stream",
	}
}

// Raw returns the identified MIME/media type as bytes
func (dr *DetectResource) Raw() ([]byte, error) {
	req, err := dr.newRequest()
	if err != nil {
		return nil, err
	}

	return dr.client.raw(req)
}

// Text returns the identified MIME/media type as plain text
func (dr *DetectResource) Text() (string, error) {
	req, err := dr.newRequest()
	if err != nil {
		return "", err
	}

	return dr.client.text(req)
}

func (dr *DetectResource) newRequest() (*http.Request, error) {
	if dr.client.Document == nil {
		return nil, errors.New("need a document")
	}

	req, err := dr.client.NewRequest(http.MethodPut, dr.endpoint, nil)
	if err != nil {
		return nil, err
	}

	if len(dr.client.DocumentName) > 0 {
		v := fmt.Sprintf("attachment; filename=%s", dr.client.DocumentName)
		req.Header.Set("Content-Disposition", v)
	}

	return req, nil
}
