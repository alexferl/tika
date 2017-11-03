package tika

import "net/http"

// Version returns the Tika server version
func (c *Client) Version() (string, error) {
	req, err := c.NewRequest(http.MethodGet, c.Url+"/version", nil)
	if err != nil {
		return "", err
	}

	return c.text(req)
}
