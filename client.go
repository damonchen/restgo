package restgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Client struct {
	req *Request
}

func NewClient(req *Request) *Client {
	return &Client{
		req: req,
	}
}

func (c Client) Do() (*Response, error) {
	switch c.req.Method {
	case "GET":
		return c.get()
	case "PUT":
		return c.put()
	case "POST":
		return c.post()
	case "PATCH":
		return c.patch()
	case "HEAD":
		return c.head()
	case "DELETE":
		return c.delete()
	default:
		return nil, errors.New("Unsupport")
	}
}

func (c Client) get() (*Response, error) {
	req, err := http.NewRequest("GET", c.req.URL, nil)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	value := map[string]interface{}{}
	err = json.Unmarshal(data, &value)
	if err != nil {
		return nil, err
	}

	headers := map[string]string{}
	for key := range resp.Header {
		value := resp.Header.Get(key)
		headers[key] = value
	}
	statusCode := resp.StatusCode

	return &Response{
		StatusCode: statusCode,
		JSON:       value,
		Headers:    headers,
	}, nil
}

func (c Client) post() (*Response, error) {
	data, err := json.Marshal(c.req.JSON)
	if err != nil {
		return nil, err
	}
	buff := bytes.NewBuffer(data)

	req, err := http.NewRequest("POST", c.req.URL, buff)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	value := map[string]interface{}{}
	err = json.Unmarshal(data, &value)
	if err != nil {
		return nil, err
	}

	headers := map[string]string{}
	for key := range resp.Header {
		value := resp.Header.Get(key)
		headers[key] = value
	}
	statusCode := resp.StatusCode

	return &Response{
		StatusCode: statusCode,
		JSON:       value,
		Headers:    headers,
	}, nil
}

func (c Client) delete() (*Response, error) {
	req, err := http.NewRequest("DELETE", c.req.URL, nil)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	value := map[string]interface{}{}
	err = json.Unmarshal(data, &value)
	if err != nil {
		return nil, err
	}

	headers := map[string]string{}
	for key := range resp.Header {
		value := resp.Header.Get(key)
		headers[key] = value
	}
	statusCode := resp.StatusCode

	return &Response{
		StatusCode: statusCode,
		JSON:       value,
		Headers:    headers,
	}, nil

}

func (c Client) put() (*Response, error) {
	data, err := json.Marshal(c.req.JSON)
	if err != nil {
		return nil, err
	}
	buff := bytes.NewBuffer(data)

	req, err := http.NewRequest("PUT", c.req.URL, buff)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	value := map[string]interface{}{}
	err = json.Unmarshal(data, &value)
	if err != nil {
		return nil, err
	}

	headers := map[string]string{}
	for key := range resp.Header {
		value := resp.Header.Get(key)
		headers[key] = value
	}
	statusCode := resp.StatusCode

	return &Response{
		StatusCode: statusCode,
		JSON:       value,
		Headers:    headers,
	}, nil
}

func (c Client) patch() (*Response, error) {
	data, err := json.Marshal(c.req.JSON)
	if err != nil {
		return nil, err
	}
	buff := bytes.NewBuffer(data)

	req, err := http.NewRequest("PATCH", c.req.URL, buff)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	value := map[string]interface{}{}
	err = json.Unmarshal(data, &value)
	if err != nil {
		return nil, err
	}

	headers := map[string]string{}
	for key := range resp.Header {
		value := resp.Header.Get(key)
		headers[key] = value
	}
	statusCode := resp.StatusCode

	return &Response{
		StatusCode: statusCode,
		JSON:       value,
		Headers:    headers,
	}, nil
}

func (c Client) head() (*Response, error) {
	req, err := http.NewRequest("HEAD", c.req.URL, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	value := map[string]interface{}{}
	err = json.Unmarshal(data, &value)
	if err != nil {
		return nil, err
	}

	headers := map[string]string{}
	for key := range resp.Header {
		value := resp.Header.Get(key)
		headers[key] = value
	}
	statusCode := resp.StatusCode

	return &Response{
		StatusCode: statusCode,
		JSON:       value,
		Headers:    headers,
	}, nil

}
