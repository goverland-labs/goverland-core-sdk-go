package goverlandcorewebsdk

import (
	"context"
	"fmt"
	"net/http"

	"github.com/goverland-labs/goverland-core-sdk-go/delegate"
)

func (c *Client) GetAllDelegationsByAddress(ctx context.Context, address string) (*delegate.AllDelegations, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/delegates?address=%s", c.baseURL, address), nil)
	if err != nil {
		return nil, err
	}

	var result delegate.AllDelegations
	if _, err = c.sendRequest(ctx, req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
