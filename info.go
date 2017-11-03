package tika

import "net/http"

// InfoServices represents the structure of our Info services resource
type InfoServices struct {
	client   *Client
	endpoint string
}

// Info is the entry point for interacting with the Info services
func (c *Client) Info() *InfoServices {
	return &InfoServices{client: c,
		endpoint: c.Url,
	}
}

// Detectors returns the top level Detector to be used, and any child detectors within it.
// Available as plain text, json or human readable HTML
func (is *InfoServices) Detectors() *InfoServices {
	is.endpoint += "/detectors"
	return is
}

// MimeTypes returns Mime Types, their aliases, their supertype, and the parser.
// Available as plain text, json or human readable HTML
func (is *InfoServices) MimeTypes() *InfoServices {
	is.endpoint += "/mime-types"
	return is
}

// Parsers returns all the available parsers, along with what mime types they support.
func (is *InfoServices) Parsers() *InfoServices {
	is.endpoint += "/parsers/details"
	return is
}

// Text returns the info as HTML
func (is *InfoServices) Html() (string, error) {
	req, err := is.newRequest()
	if err != nil {
		return "", err
	}

	return is.client.html(req)
}

// Json returns the info as JSON
func (is *InfoServices) Json() ([]byte, error) {
	req, err := is.newRequest()
	if err != nil {
		return nil, err
	}

	return is.client.json(req)
}

// Raw returns the info as bytes
func (is *InfoServices) Raw() ([]byte, error) {
	req, err := is.newRequest()
	if err != nil {
		return nil, err
	}

	return is.client.raw(req)
}

// Text returns the info as plain text
func (is *InfoServices) Text() (string, error) {
	req, err := is.newRequest()
	if err != nil {
		return "", err
	}

	return is.client.text(req)
}

func (is *InfoServices) newRequest() (*http.Request, error) {
	req, err := is.client.NewRequest(http.MethodGet, is.endpoint, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
