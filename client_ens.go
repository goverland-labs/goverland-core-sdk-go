package goverlandcorewebsdk

import (
	"context"
	"fmt"
	"github.com/goverland-labs/goverland-core-sdk-go/ens"
	"net/http"
	"strings"
)

type GetEnsNamesRequest struct {
	Addresses []string
}

func (c *Client) GetEnsNames(ctx context.Context, params GetEnsNamesRequest) (*ens.EnsNameList, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/ens-name", c.baseURL), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if len(params.Addresses) != 0 {
		q.Add("addresses", strings.Join(params.Addresses, ","))
	}
	req.URL.RawQuery = q.Encode()

	var result []ens.EnsName
	_, err = c.sendRequest(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return &ens.EnsNameList{
		EnsNames: result,
	}, nil
}

type GetAddressesByEnsNamesRequest struct {
	Names []string
}

func (c *Client) GetAddressesByEnsNames(ctx context.Context, params GetAddressesByEnsNamesRequest) (*ens.EnsNameList, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/ens-address", c.baseURL), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if len(params.Names) != 0 {
		q.Add("names", strings.Join(params.Names, ","))
	}
	req.URL.RawQuery = q.Encode()

	var result []ens.EnsName
	_, err = c.sendRequest(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return &ens.EnsNameList{
		EnsNames: result,
	}, nil
}
