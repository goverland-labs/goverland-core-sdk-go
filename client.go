package goverlandcorewebsdk

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

type Client struct {
	client  *http.Client
	baseURL string
}

func NewClient(baseURL string, client ...*http.Client) *Client {
	c := http.DefaultClient
	if len(client) > 0 {
		c = client[0]
	}

	return &Client{
		client:  c,
		baseURL: strings.TrimRight(baseURL, "/"),
	}
}

// Body should be already added to req
func (c *Client) sendRequest(ctx context.Context, req *http.Request, result interface{}) (http.Header, error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	log.Debug().Fields(map[string]interface{}{
		"status_code": res.StatusCode,
		"url":         req.URL.String(),
	}).Msg(string(body))

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, c.parseError(res.StatusCode, body)
	}

	if result == nil {
		return res.Header, nil
	}

	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}

	return res.Header, nil
}

func (c *Client) parseError(status int, body []byte) error {
	switch status {
	case http.StatusNotFound:
		return ErrNotFound
	case http.StatusUnauthorized:
		return ErrUnauthorized
	case http.StatusForbidden:
		return ErrForbidden
	case http.StatusTooManyRequests:
		var resp map[string]interface{}

		if err := json.Unmarshal(body, &resp); err != nil {
			return NewTooManyRequestsError(0)
		}

		seconds, ok := resp["retry_after"].(int)
		if !ok {
			return NewTooManyRequestsError(0)
		}

		return NewTooManyRequestsError(time.Duration(seconds) * time.Second)
	case http.StatusBadRequest:
		var errors map[string]interface{}

		if err := json.Unmarshal(body, &errors); err != nil {
			return NewValidationError(err.Error(), nil)
		}

		return NewValidationError("validation error", errors)
	}
	return ErrInternalServer
}
