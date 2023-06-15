package goverlandcorewebsdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateSubscriberRequest struct {
	WebhookURL string `json:"webhook_url"`
}

type Subscriber struct {
	SubscriberID string `json:"subscriber_id"`
}

func (c *Client) CreateSubscriber(ctx context.Context, webhookURL string) (*Subscriber, error) {
	body, err := json.Marshal(CreateSubscriberRequest{WebhookURL: webhookURL})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/subscribe", c.baseURL), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	var sub Subscriber
	if _, err = c.sendRaw(ctx, req, &sub); err != nil {
		return nil, err
	}

	return &sub, nil
}
