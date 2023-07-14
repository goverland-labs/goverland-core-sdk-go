package main

import (
	"context"
	"fmt"

	goverlandcorewebsdk "github.com/goverland-labs/core-web-sdk"
)

func createSubscriber() {
	cli := goverlandcorewebsdk.NewClient(defaultBaseURL)
	sub, err := cli.CreateSubscriber(context.TODO(), "http://calback.url")
	if err != nil {
		panic(err)
	}

	// save sub id for next usage
	fmt.Println(sub.SubscriberID)
}
