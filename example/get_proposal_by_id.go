package main

import (
	"context"
	"fmt"

	goverlandcorewebsdk "github.com/goverland-labs/core-web-sdk"
)

func getProposalByID() {
	cli := goverlandcorewebsdk.NewClient(defaultBaseURL)
	pr, err := cli.GetProposal(context.TODO(), "ed47b0bc-ecd9-4cd1-89c7-8bdcecb7a9d7")
	if err != nil {
		panic(err)
	}

	fmt.Println(pr.ID, pr.Network, pr.DaoID)
}
