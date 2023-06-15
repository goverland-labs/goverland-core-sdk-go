package goverlandcorewebsdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type SubscribeOnDaoRequest struct {
	DaoID string `json:"dao"`
}

func (c *Client) SubscribeOnDao(ctx context.Context, id string) error {
	body, err := json.Marshal(SubscribeOnDaoRequest{DaoID: id})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/subscriptions", c.baseURL), bytes.NewReader(body))
	if err != nil {
		return err
	}

	if _, err = c.sendRequest(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

type UnsubscribeFromDaoRequest struct {
	DaoID string `json:"dao"`
}

func (c *Client) UnsubscribeFromDao(ctx context.Context, id string) error {
	body, err := json.Marshal(UnsubscribeFromDaoRequest{DaoID: id})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/subscriptions", c.baseURL), bytes.NewReader(body))
	if err != nil {
		return err
	}

	if _, err = c.sendRequest(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
