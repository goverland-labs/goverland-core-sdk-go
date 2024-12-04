package goverlandcorewebsdk

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/goverland-labs/goverland-core-sdk-go/delegate"
)

func (c *Client) GetTopDelegatesByAddress(ctx context.Context, address string) (*delegate.TopDelegates, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/user/%s/delegates/top", c.baseURL, address), nil)
	if err != nil {
		return nil, err
	}

	var result delegate.TopDelegates
	if _, err = c.sendRequest(ctx, req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

type GetDelegatesListRequest struct {
	Address string
	DaoID   string
	Offset  int
	Limit   int
}

func (c *Client) GetDelegatesList(ctx context.Context, params GetDelegatesListRequest) (*delegate.DelegatesList, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/user/%s/delegates/%s/list", c.baseURL, params.Address, params.DaoID), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if params.Offset != 0 {
		q.Add("offset", strconv.Itoa(params.Offset))
	}
	if params.Limit != 0 {
		q.Add("limit", strconv.Itoa(params.Limit))
	}

	req.URL.RawQuery = q.Encode()

	var result delegate.DelegatesList
	if _, err = c.sendRequest(ctx, req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetTopDelegatorsByAddress(ctx context.Context, address string) (*delegate.TopDelegators, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/user/%s/delegators/top", c.baseURL, address), nil)
	if err != nil {
		return nil, err
	}

	var result delegate.TopDelegators
	if _, err = c.sendRequest(ctx, req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetTotalDelegationsByAddress(ctx context.Context, address string) (*delegate.TotalDelegations, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/user/%s/delegations/total", c.baseURL, address), nil)
	if err != nil {
		return nil, err
	}

	var result delegate.TotalDelegations
	if _, err = c.sendRequest(ctx, req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

type GetDelegatorsListRequest struct {
	Address string
	DaoID   string
	Offset  int
	Limit   int
}

func (c *Client) GetDelegatorsList(ctx context.Context, params GetDelegatorsListRequest) (*delegate.DelegatorsList, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/user/%s/delegators/%s/list", c.baseURL, params.Address, params.DaoID), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if params.Offset != 0 {
		q.Add("offset", strconv.Itoa(params.Offset))
	}
	if params.Limit != 0 {
		q.Add("limit", strconv.Itoa(params.Limit))
	}

	req.URL.RawQuery = q.Encode()

	var result delegate.DelegatorsList
	if _, err = c.sendRequest(ctx, req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
