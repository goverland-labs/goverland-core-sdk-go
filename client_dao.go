package goverlandcorewebsdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"

	"github.com/goverland-labs/goverland-core-sdk-go/dao"
	"github.com/goverland-labs/goverland-core-sdk-go/feed"
)

func (c *Client) GetDao(ctx context.Context, id string) (*dao.Dao, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/daos/%s", c.baseURL, id), nil)
	if err != nil {
		return nil, err
	}

	var result dao.Dao
	if _, err = c.sendRequest(ctx, req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

type GetDaoListRequest struct {
	Offset      int      `json:"offset"`
	Limit       int      `json:"limit"`
	Query       string   `url:"query"`
	Category    string   `url:"category"`
	DaoIDS      []string `url:"daos"`
	FungibleIDs []string `json:"fungible_ids"`
}

func (c *Client) GetDaoList(ctx context.Context, params GetDaoListRequest) (*dao.List, error) {
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/daos", c.baseURL), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	var result []dao.Dao
	headers, err := c.sendRequest(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return &dao.List{
		Items:    result,
		Offset:   GetOffsetFromHeaders(headers),
		Limit:    GetLimitFromHeaders(headers),
		TotalCnt: GetTotalCntFromHeaders(headers),
	}, nil
}

type GetDaoTopRequest struct {
	Limit int
}

func (c *Client) GetDaoTop(ctx context.Context, params GetDaoTopRequest) (*dao.TopCategories, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/daos/top", c.baseURL), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if params.Limit != 0 {
		q.Add("limit", strconv.Itoa(params.Limit))
	}
	req.URL.RawQuery = q.Encode()

	var result dao.TopCategories
	if _, err = c.sendRequest(ctx, req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

type GetDaoFeedRequest struct {
	Offset int
	Limit  int
}

func (c *Client) GetDaoFeed(ctx context.Context, id uuid.UUID, params GetDaoFeedRequest) (*feed.Feed, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/daos/%s/feed", c.baseURL, id.String()), nil)
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

	var result []feed.Item
	headers, err := c.sendRequest(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return &feed.Feed{
		Items:    result,
		Offset:   GetOffsetFromHeaders(headers),
		Limit:    GetLimitFromHeaders(headers),
		TotalCnt: GetTotalCntFromHeaders(headers),
	}, nil
}

func (c *Client) GetDaoRecommendations(ctx context.Context) (dao.Recommendations, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/daos/recommendations", c.baseURL), nil)
	if err != nil {
		return nil, err
	}

	var result dao.Recommendations
	if _, err = c.sendRequest(ctx, req, &result); err != nil {
		return nil, err
	}

	return result, nil
}

type GetDelegatesRequest struct {
	Query          string
	By             string
	DelegationType string
	ChainID        string
	Offset         int
	Limit          int
}

func (c *Client) GetDelegates(ctx context.Context, id uuid.UUID, params GetDelegatesRequest) (dao.DelegatesResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/daos/%s/delegates", c.baseURL, id.String()), nil)
	if err != nil {
		return dao.DelegatesResponse{}, err
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

	var result dao.DelegatesResponse
	if _, err = c.sendRequest(ctx, req, &result); err != nil {
		return dao.DelegatesResponse{}, err
	}

	return result, nil
}

type GetDelegateProfileRequest struct {
	Address        string
	DelegationType string
	ChainID        string
}

func (c *Client) GetDelegateProfile(ctx context.Context, id uuid.UUID, params GetDelegateProfileRequest) (dao.DelegateProfile, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/daos/%s/delegate-profile", c.baseURL, id.String()), nil)
	if err != nil {
		return dao.DelegateProfile{}, err
	}

	q := req.URL.Query()
	q.Add("address", params.Address)
	if params.DelegationType != "" {
		q.Add("delegation_type", params.DelegationType)
	}
	if params.ChainID != "" {
		q.Add("chain_id", params.ChainID)
	}
	req.URL.RawQuery = q.Encode()

	var result dao.DelegateProfile
	_, err = c.sendRequest(ctx, req, &result)

	return result, err
}

func (c *Client) GetDaoTokenInfo(ctx context.Context, id uuid.UUID) (dao.TokenInfo, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/daos/%s/token-info", c.baseURL, id.String()), nil)
	if err != nil {
		return dao.TokenInfo{}, err
	}

	var result dao.TokenInfo
	_, err = c.sendRequest(ctx, req, &result)

	return result, err
}

func (c *Client) GetDaoTokenChart(ctx context.Context, id uuid.UUID, period string) (dao.TokenChart, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/daos/%s/token-chart", c.baseURL, id.String()), nil)
	if err != nil {
		return dao.TokenChart{}, err
	}

	q := req.URL.Query()
	q.Add("period", period)
	req.URL.RawQuery = q.Encode()

	var result dao.TokenChart
	_, err = c.sendRequest(ctx, req, &result)

	return result, err
}

type GetDelegatorsRequest struct {
	Address        string
	DelegationType string
	ChainID        string
	Offset         int
	Limit          int
}

func (c *Client) GetDelegators(ctx context.Context, id uuid.UUID, params GetDelegatorsRequest) (*dao.Delegators, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/daos/%s/delegates/%s/delegators", c.baseURL, id.String(), params.Address), nil)
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
	if params.ChainID != "" {
		q.Add("chain_id", params.ChainID)
	}
	if params.DelegationType != "" {
		q.Add("delegation_type", params.DelegationType)
	}

	req.URL.RawQuery = q.Encode()

	var result []dao.Delegator
	headers, err := c.sendRequest(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return &dao.Delegators{
		Items:    result,
		Offset:   GetOffsetFromHeaders(headers),
		Limit:    GetLimitFromHeaders(headers),
		TotalCnt: GetTotalCntFromHeaders(headers),
	}, nil
}
