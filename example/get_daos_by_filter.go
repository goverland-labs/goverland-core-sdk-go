package main

import (
	"context"
	"fmt"

	goverlandcorewebsdk "github.com/goverland-labs/core-web-sdk"
)

func getDaoListByFilter() {
	cli := goverlandcorewebsdk.NewClient(defaultBaseURL)
	resp, err := cli.GetDaoList(context.TODO(), goverlandcorewebsdk.GetDaoListRequest{
		Limit: 2,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
