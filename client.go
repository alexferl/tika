package tika

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
)

const userAgent = "go-tika-client"

// Client represents the structure of the client for interacting with the Tika Rest API
type Client struct {
	httpClient   *http.Client
	Document     io.Reader
	DocumentName string
	Key          string
	Options
}

// Options represents the structure of the options for the Client
type Options struct {
	Url string
}

// NewRequest returns a new http.Request
func (c *Client) NewRequest(method, endpoint string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// Do makes an HTTP request and returns bytes
func (c *Client) Do(req *http.Request) ([]byte, error) {
	req.Header.Set("User-Agent", userAgent)

	res, getErr := c.httpClient.Do(req)
	if getErr != nil {
		return nil, getErr
	}

	if res.StatusCode != 200 {
		logrus.Fatalf("Status: %s", res.Status)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	return body, nil
}

// NewClient takes Options and returns a Client with the
func NewClient(options *Options) *Client {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	var netClient = &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}

	return &Client{
		httpClient: netClient,
		Options: *options,
	}
}

func (c *Client) csv(req *http.Request) (string, error) {
	req.Header.Set("Accept", "text/csv")
	res, err := c.Do(req)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

func (c *Client) html(req *http.Request) (string, error) {
	req.Header.Set("Accept", "text/html")
	res, err := c.Do(req)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

func (c *Client) json(req *http.Request) ([]byte, error) {
	req.Header.Set("Accept", "application/json")
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	var raw map[string]interface{}
	json.Unmarshal(res, &raw)
	out, err := json.Marshal(raw)
	if err != nil {
		return nil, err
	}

	return out, err
}

func (c *Client) raw(req *http.Request) ([]byte, error) {
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) text(req *http.Request) (string, error) {
	req.Header.Set("Accept", "text/plain")
	res, err := c.Do(req)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
