package main

import (
	"context"
	"fmt"

	goverlandcorewebsdk "github.com/goverland-labs/goverland-core-sdk-go"
)

func getFeedByFilter() {
	cli := goverlandcorewebsdk.NewClient(defaultBaseURL)
	resp, err := cli.GetFeedByFilters(context.TODO(), goverlandcorewebsdk.FeedByFiltersRequest{
		DaoList: []string{"621cf7b5-89c3-4da9-a230-fc3e6ee01ffe", "8599e777-c4b1-499c-889b-d4154be6a6f5"},
		Types:   []string{"proposal"},
		Limit:   2,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
