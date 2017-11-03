package tika

import (
	"errors"
	"net/http"
)

// TikaResource represents the structure of our resource
type TikaResource struct {
	client   *Client
	endpoint string
}

// Tika is the entry point for interacting with the Tika resource
func (c *Client) Tika() *TikaResource {
	return &TikaResource{client: c,
		endpoint: c.Url + "/tika",
	}
}

// Hello returns a greeting
func (c *Client) Hello() (string, error) {
	req, err := c.NewRequest(http.MethodGet, c.Tika().endpoint, nil)
	if err != nil {
		return "", err
	}

	return c.text(req)
}

// Html returns extracted text as HTML
func (tr *TikaResource) Html() (string, error) {
	req, err := tr.newRequest()
	if err != nil {
		return "", err
	}

	return tr.client.html(req)
}

// Raw returns extracted text as bytes
func (tr *TikaResource) Raw() ([]byte, error) {
	req, err := tr.newRequest()
	if err != nil {
		return nil, err
	}

	return tr.client.raw(req)
}

// Text returns extracted text as plain text
func (tr *TikaResource) Text() (string, error) {
	req, err := tr.newRequest()
	if err != nil {
		return "", err
	}

	return tr.client.text(req)
}

func (tr *TikaResource) newRequest() (*http.Request, error) {
	if tr.client.Document == nil {
		return nil, errors.New("need a document")
	}

	req, err := tr.client.NewRequest(http.MethodPut, tr.endpoint, tr.client.Document)
	if err != nil {
		return nil, err
	}

	return req, nil
}
