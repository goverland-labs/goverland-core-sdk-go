package main

import (
	"context"
	"fmt"

	goverlandcorewebsdk "github.com/goverland-labs/core-web-sdk"
)

func updateSubscriber() {
	cli := goverlandcorewebsdk.NewClient(defaultBaseURL)
	err := cli.UpdateSubscriber(context.TODO(), "4fe6b578-b34b-402a-8b88-bb0f83041119", "http://new.callback.url")
	if err != nil {
		panic(err)
	}

	fmt.Println("updated")
}
