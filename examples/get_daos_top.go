package main

import (
	"context"
	"fmt"

	goverlandcorewebsdk "github.com/goverland-labs/goverland-core-sdk-go"
)

func getDaoTop() {
	cli := goverlandcorewebsdk.NewClient(defaultBaseURL)
	resp, err := cli.GetDaoTop(context.TODO(), goverlandcorewebsdk.GetDaoTopRequest{
		Limit: 1,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
