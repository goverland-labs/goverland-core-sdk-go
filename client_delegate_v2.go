package goverlandcorewebsdk

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"

	"github.com/goverland-labs/goverland-core-sdk-go/delegate"
)

func (c *Client) GetDelegatesV2(
	ctx context.Context,
	id uuid.UUID,
	params delegate.GetDelegatesV2Request,
) (delegate.GetDelegatesV2Response, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/daos/%s/delegates", c.baseURL, id.String()), nil)
	if err != nil {
		return delegate.GetDelegatesV2Response{}, err
	}

	q := req.URL.Query()
	if params.Offset != 0 {
		q.Add("offset", strconv.Itoa(params.Offset))
	}
	if params.Limit != 0 {
		q.Add("limit", strconv.Itoa(params.Limit))
	}
	if params.Query != "" {
		q.Add("query", params.Query)
	}
	if params.By != "" {
		q.Add("by", params.By)
	}
	if params.DelegationType != "" {
		q.Add("delegation_type", params.DelegationType)
	}
	if params.ChainID != "" {
		q.Add("chain_id", params.ChainID)
	}
	req.URL.RawQuery = q.Encode()

	var result delegate.GetDelegatesV2Response
	if _, err = c.sendRequest(ctx, req, &result); err != nil {
		return delegate.GetDelegatesV2Response{}, err
	}

	return result, nil
}

func (c *Client) GetUserDelegatesTopV2(
	ctx context.Context,
	address string,
) (delegate.GetUserDelegatesTopV2Response, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/users/%s/delegates/top", c.baseURL, address), nil)
	if err != nil {
		return delegate.GetUserDelegatesTopV2Response{}, err
	}

	var result delegate.GetUserDelegatesTopV2Response
	if _, err = c.sendRequest(ctx, req, &result); err != nil {
		return delegate.GetUserDelegatesTopV2Response{}, err
	}

	return result, nil
}

func (c *Client) GetUserTopDelegatorsByDaoV2(
	ctx context.Context,
	address string,
	daoID uuid.UUID,
) (delegate.GetUserDelegatorsTopV2Response, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/daos/%s/delegates/%s/delegators", c.baseURL, daoID.String(), address), nil)
	if err != nil {
		return delegate.GetUserDelegatorsTopV2Response{}, err
	}

	var result delegate.GetUserDelegatorsTopV2Response
	if _, err = c.sendRequest(ctx, req, &result); err != nil {
		return delegate.GetUserDelegatorsTopV2Response{}, err
	}

	return result, nil
}

func (c *Client) GetUserDelegatesListV2(
	ctx context.Context,
	id uuid.UUID,
	address string,
	params delegate.GetUserDelegatesListV2Request,
) (delegate.GetUserDelegatesListV2Request, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/user/%s/delegates/%s/list", c.baseURL, address, id.String()), nil)
	if err != nil {
		return delegate.GetUserDelegatesListV2Request{}, err
	}

	q := req.URL.Query()
	if params.Offset != 0 {
		q.Add("offset", strconv.Itoa(params.Offset))
	}
	if params.Limit != 0 {
		q.Add("limit", strconv.Itoa(params.Limit))
	}
	if params.DelegationType != "" {
		q.Add("delegation_type", params.DelegationType)
	}
	if params.ChainID != "" {
		q.Add("chain_id", params.ChainID)
	}
	req.URL.RawQuery = q.Encode()

	var result delegate.GetUserDelegatesListV2Request
	if _, err = c.sendRequest(ctx, req, &result); err != nil {
		return delegate.GetUserDelegatesListV2Request{}, err
	}

	return result, nil
}

func (c *Client) GetUserDelegatorsV2(
	ctx context.Context,
	id uuid.UUID,
	address string,
	params delegate.GetUserDelegatorsV2Request,
) (delegate.GetDelegatorsV2Response, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/user/%s/delegators/%s/list", c.baseURL, address, id.String()), nil)
	if err != nil {
		return delegate.GetDelegatorsV2Response{}, err
	}

	q := req.URL.Query()
	if params.Offset != 0 {
		q.Add("offset", strconv.Itoa(params.Offset))
	}
	if params.Limit != 0 {
		q.Add("limit", strconv.Itoa(params.Limit))
	}
	if params.Query != "" {
		q.Add("query", params.Query)
	}
	if params.DelegationType != "" {
		q.Add("delegation_type", params.DelegationType)
	}
	if params.ChainID != "" {
		q.Add("chain_id", params.ChainID)
	}
	req.URL.RawQuery = q.Encode()

	var result delegate.GetDelegatorsV2Response
	if _, err = c.sendRequest(ctx, req, &result); err != nil {
		return delegate.GetDelegatorsV2Response{}, err
	}

	return result, nil
}

func (c *Client) GetUserDelegatorsTopV2(
	ctx context.Context,
	address string,
) (delegate.GetUserDelegatorsTopV2Response, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/users/%s/delegators/top", c.baseURL, address), nil)
	if err != nil {
		return delegate.GetUserDelegatorsTopV2Response{}, err
	}

	var result delegate.GetUserDelegatorsTopV2Response
	if _, err = c.sendRequest(ctx, req, &result); err != nil {
		return delegate.GetUserDelegatorsTopV2Response{}, err
	}

	return result, nil
}
