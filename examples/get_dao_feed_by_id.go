package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	goverlandcorewebsdk "github.com/goverland-labs/goverland-core-sdk-go"
)

func getDaoFeedByID() {
	cli := goverlandcorewebsdk.NewClient(defaultBaseURL)
	resp, err := cli.GetDaoFeed(context.TODO(), uuid.MustParse("e604b913-dd21-4f1d-844a-21130d3e818d"), goverlandcorewebsdk.GetDaoFeedRequest{
		Limit: 2,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
