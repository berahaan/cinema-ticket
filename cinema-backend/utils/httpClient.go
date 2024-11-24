package utils

import (
	"net/http"
)

type CustomTransport struct {
	Transport http.RoundTripper
	Headers   map[string]string
}

func (c *CustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for key, value := range c.Headers {
		req.Header.Set(key, value)
	}
	return c.Transport.RoundTrip(req)
}
