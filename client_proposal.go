package goverlandcorewebsdk

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/goverland-labs/core-web-sdk/proposal"
)

func (c *Client) GetProposal(ctx context.Context, id string) (*proposal.Proposal, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/proposals/%s", c.baseURL, id), nil)
	if err != nil {
		return nil, err
	}

	var result proposal.Proposal
	if _, err = c.sendRequest(ctx, req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

type GetProposalListRequest struct {
	Offset   int
	Limit    int
	Dao      string
	Category string
	Title    string
}

func (c *Client) GetProposalList(ctx context.Context, params GetProposalListRequest) (*proposal.List, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/proposals", c.baseURL), nil)
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
	if params.Dao != "" {
		q.Add("dao", params.Dao)
	}
	if params.Category != "" {
		q.Add("category", params.Category)
	}
	if params.Title != "" {
		q.Add("title", params.Title)
	}
	req.URL.RawQuery = q.Encode()

	var result []proposal.Proposal
	headers, err := c.sendRequest(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return &proposal.List{
		Items:    result,
		Offset:   GetOffsetFromHeaders(headers),
		Limit:    GetLimitFromHeaders(headers),
		TotalCnt: GetTotalCntFromHeaders(headers),
	}, nil
}

type GetProposalVotesRequest struct {
	Offset int
	Limit  int
}

type GetProposalTopRequest struct {
	Offset int
	Limit  int
}

func (c *Client) GetProposalTop(ctx context.Context, params GetProposalTopRequest) (*proposal.List, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/proposals/top", c.baseURL), nil)
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

	var result []proposal.Proposal
	headers, err := c.sendRequest(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return &proposal.List{
		Items:    result,
		Offset:   GetOffsetFromHeaders(headers),
		Limit:    GetLimitFromHeaders(headers),
		TotalCnt: GetTotalCntFromHeaders(headers),
	}, nil
}

func (c *Client) GetProposalVotes(ctx context.Context, id string, params GetProposalVotesRequest) (*proposal.VoteList, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/proposals/%s/votes", c.baseURL, id), nil)
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

	var result []proposal.Vote
	headers, err := c.sendRequest(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return &proposal.VoteList{
		Items:    result,
		Offset:   GetOffsetFromHeaders(headers),
		Limit:    GetLimitFromHeaders(headers),
		TotalCnt: GetTotalCntFromHeaders(headers),
	}, nil
}
