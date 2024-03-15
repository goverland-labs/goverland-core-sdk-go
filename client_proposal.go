package goverlandcorewebsdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/goverland-labs/goverland-core-sdk-go/proposal"
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
	Offset      int
	Limit       int
	Dao         string
	Category    string
	Title       string
	ProposalIDs []string
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
	if len(params.ProposalIDs) != 0 {
		q.Add("proposals", strings.Join(params.ProposalIDs, ","))
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
	OrderByVoter string
	Offset       int
	Limit        int
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

	if params.OrderByVoter != "" {
		q.Add("voter", params.OrderByVoter)
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
		TotalVp:  GetTotalVpFromHeaders(headers),
	}, nil
}

type ValidateVoteRequest struct {
	Voter string `json:"voter"`
}

func (c *Client) ValidateVote(ctx context.Context, proposalID string, params ValidateVoteRequest) (proposal.VoteValidation, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return proposal.VoteValidation{}, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/proposals/%s/votes/validate", c.baseURL, proposalID), bytes.NewReader(data))
	if err != nil {
		return proposal.VoteValidation{}, err
	}

	var result proposal.VoteValidation
	_, err = c.sendRequest(ctx, req, &result)
	if err != nil {
		return proposal.VoteValidation{}, err
	}

	return result, nil
}

type PrepareVoteRequest struct {
	Voter  string          `json:"voter"`
	Choice json.RawMessage `json:"choice"`
	Reason *string         `json:"reason,omitempty"`
}

func (c *Client) PrepareVote(ctx context.Context, proposalID string, params PrepareVoteRequest) (proposal.VotePreparation, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return proposal.VotePreparation{}, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/proposals/%s/votes/prepare", c.baseURL, proposalID), bytes.NewReader(data))
	if err != nil {
		return proposal.VotePreparation{}, err
	}

	var result proposal.VotePreparation
	_, err = c.sendRequest(ctx, req, &result)
	if err != nil {
		return proposal.VotePreparation{}, err
	}

	return result, nil
}

type VoteRequest struct {
	ID  string `json:"id"`
	Sig string `json:"sig"`
}

func (c *Client) Vote(ctx context.Context, params VoteRequest) (proposal.SuccessfulVote, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return proposal.SuccessfulVote{}, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/proposals/votes", c.baseURL), bytes.NewReader(data))
	if err != nil {
		return proposal.SuccessfulVote{}, err
	}

	var result proposal.SuccessfulVote
	_, err = c.sendRequest(ctx, req, &result)
	if err != nil {
		return proposal.SuccessfulVote{}, err
	}

	return result, nil
}
