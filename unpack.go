package tika

import (
	"errors"
	"net/http"
)

// UnpackResource represents the structure of our Unpack resource
type UnpackResource struct {
	client   *Client
	endpoint string
}

// Unpack is the entry point for interacting with the Unpack resource
func (c *Client) Unpack() *UnpackResource {
	return &UnpackResource{client: c,
		endpoint: c.Url + "/unpack",
	}
}

// All returns the content and metadata. Text is stored in __TEXT__ file, metadata csv in __METADATA__
func (ur *UnpackResource) All() *UnpackResource {
	ur.endpoint += "/all"
	return ur
}

// Raw returns the metadata as *http.Response
func (ur *UnpackResource) Raw() ([]byte, error) {
	req, err := ur.newRequest()
	if err != nil {
		return nil, err
	}

	return ur.client.raw(req)
}

// Tar returns the metadata (and content if using All) as a tarball
func (ur *UnpackResource) Tar() ([]byte, error) {
	req, err := ur.newRequest()
	req.Header.Set("Accept", "application/x-tar")
	if err != nil {
		return nil, err
	}

	return ur.client.raw(req)
}

// Zip returns the metadata (and content if using All) as a zip file
func (ur *UnpackResource) Zip() ([]byte, error) {
	req, err := ur.newRequest()
	req.Header.Set("Accept", "application/zip")
	if err != nil {
		return nil, err
	}

	return ur.client.raw(req)
}

func (ur *UnpackResource) newRequest() (*http.Request, error) {
	if ur.client.Document == nil {
		return nil, errors.New("need a document")
	}

	req, err := ur.client.NewRequest(http.MethodPut, ur.endpoint, ur.client.Document)
	if err != nil {
		return nil, err
	}

	return req, nil
}
