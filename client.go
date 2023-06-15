package goverlandcorewebsdk

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

type Client struct {
	client  *http.Client
	baseURL string

	subscriberID     string
	webhookURL       string
	subscriberIDLock sync.Locker
}

func NewClient(baseURL, subscriberID, webhookURL string, client ...*http.Client) *Client {
	c := http.DefaultClient
	if len(client) > 0 {
		c = client[0]
	}

	return &Client{
		client:           c,
		baseURL:          strings.TrimRight(baseURL, "/"),
		subscriberID:     subscriberID,
		webhookURL:       webhookURL,
		subscriberIDLock: &sync.Mutex{},
	}
}

// Body should be already added to req
func (c *Client) sendRequest(ctx context.Context, req *http.Request, result interface{}) (http.Header, error) {
	subscriberID, err := c.getSubscriberID(ctx)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", subscriberID)

	return c.sendRaw(ctx, req, result)
}

func (c *Client) sendRaw(ctx context.Context, req *http.Request, result interface{}) (http.Header, error) {
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

func (c *Client) getSubscriberID(ctx context.Context) (string, error) {
	c.subscriberIDLock.Lock()
	defer c.subscriberIDLock.Unlock()

	if c.subscriberID != "" {
		return c.subscriberID, nil
	}

	response, err := c.CreateSubscriber(ctx, c.webhookURL)
	if err != nil {
		return "", err
	}

	c.subscriberID = response.SubscriberID

	return c.subscriberID, nil
}

func (c *Client) RestoreAccessToken(token string, expiresAt time.Time) {
	c.subscriberIDLock.Lock()
	defer c.subscriberIDLock.Unlock()

	c.subscriberID = token
}
