package goverlandcorewebsdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/goverland-labs/goverland-core-sdk-go/subscriber"
)

type CreateSubscriberRequest struct {
	WebhookURL string `json:"webhook_url"`
}

func (c *Client) CreateSubscriber(ctx context.Context, webhookURL string) (*subscriber.Subscriber, error) {
	body, err := json.Marshal(CreateSubscriberRequest{WebhookURL: webhookURL})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/subscribe", c.baseURL), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	var sub subscriber.Subscriber
	if _, err = c.sendRequest(ctx, req, &sub); err != nil {
		return nil, err
	}

	return &sub, nil
}

type UpdateSubscriberRequest struct {
	WebhookURL string `json:"webhook_url"`
}

func (c *Client) UpdateSubscriber(ctx context.Context, subscriberID uuid.UUID, webhookURL string) error {
	body, err := json.Marshal(UpdateSubscriberRequest{WebhookURL: webhookURL})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/subscribe", c.baseURL), bytes.NewReader(body))
	if err != nil {
		return err
	}

	c.prepareAuthRequest(req, subscriberID)

	if _, err = c.sendRequest(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) prepareAuthRequest(req *http.Request, subscriberID uuid.UUID) {
	req.Header.Set("Authorization", subscriberID.String())
}

type SubscribeOnDaoRequest struct {
	DaoID uuid.UUID `json:"dao"`
}

func (c *Client) SubscribeOnDao(ctx context.Context, subscriberID, daoID uuid.UUID) error {
	body, err := json.Marshal(SubscribeOnDaoRequest{DaoID: daoID})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/subscriptions", c.baseURL), bytes.NewReader(body))
	if err != nil {
		return err
	}

	c.prepareAuthRequest(req, subscriberID)

	if _, err = c.sendRequest(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

type UnsubscribeFromDaoRequest struct {
	DaoID uuid.UUID `json:"dao"`
}

func (c *Client) UnsubscribeFromDao(ctx context.Context, subscriberID, daoID uuid.UUID) error {
	body, err := json.Marshal(UnsubscribeFromDaoRequest{DaoID: daoID})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/subscriptions", c.baseURL), bytes.NewReader(body))
	if err != nil {
		return err
	}

	c.prepareAuthRequest(req, subscriberID)

	if _, err = c.sendRequest(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
