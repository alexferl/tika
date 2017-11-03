package tika

import (
	"errors"
	"fmt"
	"net/http"
)

// TranslateResource represents the structure of our Translate resource
type TranslateResource struct {
	client   *Client
	endpoint string
}

// TranslateOptions represents the structure of our Translate options
type TranslateOptions struct {
	Translator string
	Src        string
	Dest       string
}

// Translate is the entry point for interacting with the Translate resource
func (c *Client) Translate(options *TranslateOptions) (string, error) {
	endpoint := fmt.Sprintf("/translate/all/%s/%s/%s", options.Translator, options.Src, options.Dest)
	tr := &TranslateResource{client: c,
		endpoint: c.Url + endpoint,
	}

	req, err := tr.newRequest()
	if err != nil {
		return "", err
	}

	return tr.client.text(req)
}

func (tr *TranslateResource) newRequest() (*http.Request, error) {
	if tr.client.Document == nil {
		return nil, errors.New("need a document")
	}

	req, err := tr.client.NewRequest(http.MethodPut, tr.endpoint, tr.client.Document)
	if err != nil {
		return nil, err
	}

	return req, nil
}
