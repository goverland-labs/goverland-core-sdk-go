package main

import (
	"context"
	"fmt"

	goverlandcorewebsdk "github.com/goverland-labs/goverland-core-sdk-go"
)

func getProposalListByFilter() {
	cli := goverlandcorewebsdk.NewClient(defaultBaseURL)
	resp, err := cli.GetProposalList(context.TODO(), goverlandcorewebsdk.GetProposalListRequest{
		Limit:  2,
		Offset: 3,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
