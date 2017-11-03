package tika

func NewTestClient() *Client {
	return NewClient(&Options{Url: "http://127.0.0.1:9998"})
}
