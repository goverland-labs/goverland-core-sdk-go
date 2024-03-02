package main

import (
	"context"
	"fmt"

	goverlandcorewebsdk "github.com/goverland-labs/goverland-core-sdk-go"
)

func getProposalVotes() {
	cli := goverlandcorewebsdk.NewClient(defaultBaseURL)
	resp, err := cli.GetProposalVotes(context.TODO(), "18b58d03-9ff0-4aa0-beb7-9bfec2aff406", goverlandcorewebsdk.GetProposalVotesRequest{
		Limit:  2,
		Offset: 3,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
