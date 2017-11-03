package tika

import (
	"errors"
	"net/http"
)

// LanguageService represents the structure of our Language service resource
type LanguageService struct {
	client   *Client
	endpoint string
}

// Language is the entry point for interacting with the Language service
func (c *Client) Language() *LanguageService {
	return &LanguageService{client: c,
		endpoint: c.Url + "/language/stream",
	}
}

// Raw returns the identified language as bytes
func (ls *LanguageService) Raw() ([]byte, error) {
	req, err := ls.newRequest()
	if err != nil {
		return nil, err
	}

	return ls.client.raw(req)
}

// Text returns the identified language as plain text
func (ls *LanguageService) Text() (string, error) {
	req, err := ls.newRequest()
	if err != nil {
		return "", err
	}

	return ls.client.text(req)
}

func (ls *LanguageService) newRequest() (*http.Request, error) {
	if ls.client.Document == nil {
		return nil, errors.New("need a document")
	}

	req, err := ls.client.NewRequest(http.MethodPut, ls.endpoint, ls.client.Document)

	if err != nil {
		return nil, err
	}

	return req, nil
}
