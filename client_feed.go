package goverlandcorewebsdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/goverland-labs/core-web-sdk/feed"
)

type FeedByFiltersRequest struct {
	DaoList  []string `json:"dao_list"`
	IsActive *bool    `json:"is_active,omitempty"`
	Types    []string `json:"types"`
	Actions  []string `json:"actions"`
	Offset   int
	Limit    int
}

func (c *Client) GetFeedByFilters(ctx context.Context, params FeedByFiltersRequest) (*feed.Feed, error) {
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/feed", c.baseURL), bytes.NewReader(body))
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
