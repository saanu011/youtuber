package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var (
	ErrClient = fmt.Sprintf("HTTP client error")
	ErrServer = fmt.Sprintf("HTTP server error")

	ErrBadRequest    = fmt.Sprintf("HTTP bad request error")
	ErrNotFound      = fmt.Sprintf("HTTP not found error")
	ErrConflict      = fmt.Sprintf("HTTP conflict error")
	ErrUnprocessable = fmt.Sprintf("HTTP unprocessable entity error")
)

type Client struct {
	commandName string
	client      *http.Client
}

func New(commandName string) *Client {
	return &Client{
		commandName: commandName,
		client: &http.Client{
			Timeout: 10,
		},
	}
}

type Request struct {
	ctx context.Context

	method  string
	url     string
	query   url.Values
	headers http.Header
	body    interface{}

	target    interface{}
	targetErr interface{}
}

func NewRequest(ctx context.Context, method, url string, body interface{}, target interface{}) *Request {
	return &Request{
		ctx:     ctx,
		method:  method,
		url:     url,
		body:    body,
		headers: make(http.Header),
		target:  target,
	}
}

func (r *Request) WithQuery(query url.Values) *Request {
	r.query = query
	return r
}

func (r *Request) WithHeaders(headers map[string]string) *Request {
	for key, value := range headers {
		r.headers.Add(key, value)
	}

	return r
}

func (r *Request) WithErrorTarget(targetErr interface{}) *Request {
	r.targetErr = targetErr
	return r
}

func (c *Client) Do(req *Request) error {
	httpReq, err := req.buildHttpRequest(c.commandName)
	if err != nil {
		return err
	}

	res, err := c.client.Do(httpReq)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if !isSuccess(res.StatusCode) {
		errResponse, _ := readErrorBody(res, req.targetErr)

		return fmt.Errorf("err: %v with code: %s", errResponse, mapError(res.StatusCode))
	}

	if req.target == nil {
		return nil
	}

	err = json.NewDecoder(res.Body).Decode(req.target)

	if err != nil {
		return fmt.Errorf("err: %v with code: %s", err, mapError(res.StatusCode))
	}

	return nil
}

func (r *Request) buildHttpRequest(command string) (*http.Request, error) {
	var data io.Reader

	var jsonData []byte

	if r.body != nil {
		var err error
		jsonData, err = json.Marshal(r.body)

		if err != nil {
			return nil, err
		}

		data = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequestWithContext(r.ctx, r.method, r.url, data)
	if err != nil {
		return nil, err
	}

	if r.query != nil {
		req.URL.RawQuery = r.query.Encode()
	}

	if r.headers != nil {
		req.Header = r.headers
	}

	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

func readErrorBody(res *http.Response, target interface{}) (interface{}, error) {
	if target != nil {
		if err := json.NewDecoder(res.Body).Decode(target); err == nil {
			return target, nil
		}
	}

	errRes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return errRes, nil
}

func isSuccess(statusCode int) bool {
	return statusCode >= http.StatusOK && statusCode <= http.StatusAlreadyReported
}

func mapError(statusCode int) string {
	switch statusCode {
	case http.StatusBadRequest:
		return ErrBadRequest
	case http.StatusNotFound:
		return ErrNotFound
	case http.StatusConflict:
		return ErrConflict
	case http.StatusUnprocessableEntity:
		return ErrUnprocessable
	default:
		return ErrServer
	}
}

func ExecuteHTTP(ctx context.Context, client *Client, method, url string, query url.Values,
	data interface{}, target interface{}, targetErr interface{}, headers map[string]string) error {
	req := NewRequest(ctx, method, url, data, target).
		WithQuery(query).
		WithHeaders(headers).
		WithErrorTarget(targetErr)

	return client.Do(req)
}
