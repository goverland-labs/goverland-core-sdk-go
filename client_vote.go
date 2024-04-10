package goverlandcorewebsdk

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"strings"

	"github.com/goverland-labs/goverland-core-sdk-go/dao"
	"github.com/goverland-labs/goverland-core-sdk-go/proposal"
)

type GetUserVotesRequest struct {
	ProposalIDs []string
	Offset      int
	Limit       int
}

func (c *Client) GetUserVotes(ctx context.Context, address string, params GetUserVotesRequest) (*proposal.VoteList, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/user/%s/votes", c.baseURL, address), nil)
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
	if len(params.ProposalIDs) != 0 {
		q.Add("proposals", strings.Join(params.ProposalIDs, ","))
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

func (c *Client) GetUserParticipatedDaos(ctx context.Context, voter string) (*dao.DaoIds, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/user/%s/participated-daos", c.baseURL, voter), nil)
	if err != nil {
		return nil, err
	}

	var result []uuid.UUID
	headers, err := c.sendRequest(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return &dao.DaoIds{
		Ids:      result,
		TotalCnt: GetTotalCntFromHeaders(headers),
	}, nil
}
