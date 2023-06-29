package goverlandcorewebsdk

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/goverland-labs/core-web-sdk/dao"
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
	Offset   int
	Limit    int
	Query    string
	Category string
	DaoIDS   []string
}

func (c *Client) GetDaoList(ctx context.Context, params GetDaoListRequest) (*dao.List, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/daos", c.baseURL), nil)
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
	if params.Query != "" {
		q.Add("query", params.Query)
	}
	if params.Category != "" {
		q.Add("category", params.Category)
	}
	if len(params.DaoIDS) != 0 {
		q.Add("daos", strings.Join(params.DaoIDS, ","))
	}
	req.URL.RawQuery = q.Encode()

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
