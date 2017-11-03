package tika

import (
	"errors"
	"net/http"
)

// MetaResource represents the structure of our Meta resource
type MetaResource struct {
	client   *Client
	endpoint string
}

// Meta is the entry point for interacting with the Meta resource
func (c *Client) Meta() *MetaResource {
	endpoint := c.Url + "/meta"

	if len(c.Key) > 0 {
		endpoint += "/" + c.Key
	}

	return &MetaResource{client: c,
		endpoint: endpoint,
	}
}

// Raw returns the metadata as bytes
func (mr *MetaResource) Raw() ([]byte, error) {
	req, err := mr.newRequest()
	if err != nil {
		return nil, err
	}

	return mr.client.raw(req)
}

// Csv returns the metadata as comma-separated values
func (mr *MetaResource) Csv() (string, error) {
	req, err := mr.newRequest()
	if err != nil {
		return "", err
	}

	return mr.client.csv(req)
}

// Text returns the metadata as plain text
func (mr *MetaResource) Text() (string, error) {
	if len(mr.client.Key) < 1 {
		return "", errors.New("text requires a key")
	}

	req, err := mr.newRequest()
	if err != nil {
		return "", err
	}

	return mr.client.text(req)
}

// Json returns the metadata as JSON
func (mr *MetaResource) Json() ([]byte, error) {
	req, err := mr.newRequest()
	if err != nil {
		return nil, err
	}

	return mr.client.json(req)
}

func (mr *MetaResource) newRequest() (*http.Request, error) {
	if mr.client.Document == nil {
		return nil, errors.New("need a document")
	}

	req, err := mr.client.NewRequest(http.MethodPut, mr.endpoint, mr.client.Document)
	if err != nil {
		return nil, err
	}

	return req, nil
}
