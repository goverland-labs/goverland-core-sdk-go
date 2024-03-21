package goverlandcorewebsdk

import (
	"context"
	"fmt"
	"net/http"

	"github.com/goverland-labs/goverland-core-sdk-go/stats"
)

func (c *Client) GetStatsTotals(ctx context.Context) (*stats.Totals, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/stats/totals", c.baseURL), nil)
	if err != nil {
		return nil, err
	}

	var result stats.Totals
	_, err = c.sendRequest(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
