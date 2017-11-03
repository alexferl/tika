package tika

import (
	"errors"
	"net/http"
)

// RecursiveMetaResource represents the structure of our Recursive Meta resource
type RecursiveMetaResource struct {
	client   *Client
	endpoint string
}

// RMeta is the entry point for interacting with the Recursive Meta resource
func (c *Client) RMeta() *RecursiveMetaResource {
	endpoint := c.Url + "/rmeta"

	return &RecursiveMetaResource{client: c,
		endpoint: endpoint,
	}
}

// Html returns a list of Metadata objects for the container document and all embedded documents as HTML
func (rmr *RecursiveMetaResource) Html() (string, error) {
	rmr.endpoint += "/html"
	req, err := rmr.newRequest()
	if err != nil {
		return "", err
	}

	res, err := rmr.client.raw(req)

	return string(res), err
}

// Ignore returns the metadata only
func (rmr *RecursiveMetaResource) Ignore() (string, error) {
	rmr.endpoint += "/ignore"
	req, err := rmr.newRequest()
	if err != nil {
		return "", err
	}

	res, err := rmr.client.raw(req)

	return string(res), err
}

// Json returns a list of Metadata objects for the container document and all embedded documents as JSON
func (rmr *RecursiveMetaResource) Json() ([]byte, error) {
	req, err := rmr.newRequest()
	if err != nil {
		return nil, err
	}

	return rmr.client.raw(req)
}

// Raw returns a list of Metadata objects for the container document and all embedded documents as bytes
func (rmr *RecursiveMetaResource) Raw() ([]byte, error) {
	req, err := rmr.newRequest()
	if err != nil {
		return nil, err
	}

	return rmr.client.raw(req)
}

// Text returns a list of Metadata objects for the container document and all embedded documents as plain text
func (rmr *RecursiveMetaResource) Text() (string, error) {
	rmr.endpoint += "/text"
	req, err := rmr.newRequest()
	if err != nil {
		return "", err
	}

	res, err := rmr.client.raw(req)

	return string(res), err
}

func (rmr *RecursiveMetaResource) newRequest() (*http.Request, error) {
	if rmr.client.Document == nil {
		return nil, errors.New("need a document")
	}

	req, err := rmr.client.NewRequest(http.MethodPut, rmr.endpoint, rmr.client.Document)
	if err != nil {
		return nil, err
	}

	return req, nil
}
