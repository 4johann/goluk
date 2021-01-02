package http

import (
	"bytes"
	"net/http"
	"strings"
	"time"
)

type RequestMessage struct {
	Method string
	Path   string
	Header map[string]string
	Body   []byte
}

type HTTPClient struct {
	host                    string
	requestTimeoutInSeconds int
	request                 *http.Request
	response                *http.Response
	client                  *http.Client
}

func (c *HTTPClient) SendRequest(message RequestMessage) error {
	var errorToReturn error = nil

	c.createClient()
	setError := c.setRequest(message)
	if setError == nil {
		errorToReturn = c.sendRequestToService()
	} else {
		errorToReturn = setError
	}
	return errorToReturn
}

func (c *HTTPClient) GetResponse() *http.Response {
	return c.response
}

func (c *HTTPClient) SetHost(_host string) {
	c.host = _host
}

func (c *HTTPClient) GetHost() string {
	return c.host
}

func (c *HTTPClient) SetRequestTimeoutInSeconds(_requestTimeoutInSeconds int) {
	c.requestTimeoutInSeconds = _requestTimeoutInSeconds
}

func (c *HTTPClient) GetRequestTimeoutInSeconds() int {
	return c.requestTimeoutInSeconds
}

func (c *HTTPClient) createClient() {
	c.client = &http.Client{
		Timeout: time.Duration(c.requestTimeoutInSeconds) * time.Second,
	}
}

func (c *HTTPClient) setRequest(message RequestMessage) error {
	var errorToReturn error = nil
	body := bytes.NewBuffer(message.Body)
	c.request, errorToReturn = http.NewRequest(message.Method, c.getHostToSendRequest(message.Path), body)
	if errorToReturn == nil {
		c.setHeaders(message)
	}

	return errorToReturn
}

func (c *HTTPClient) sendRequestToService() error {
	var responseError error
	c.response, responseError = c.client.Do(c.request)

	return responseError
}

func (c *HTTPClient) setHeaders(message RequestMessage) {
	if message.Header == nil {
	} else {
		for key, value := range message.Header {
			c.request.Header.Set(key, value)
		}
	}
}

func (c *HTTPClient) getHostToSendRequest(path string) string {
	var host strings.Builder

	host.WriteString(c.host)
	host.WriteString(path)
	return host.String()
}
