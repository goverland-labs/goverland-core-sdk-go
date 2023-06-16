package main

import (
	"context"
	"fmt"

	goverlandcorewebsdk "github.com/goverland-labs/core-web-sdk"
)

func unsubscribeFromDao() {
	cli := goverlandcorewebsdk.NewClient(defaultBaseURL)
	err := cli.UnsubscribeFromDao(context.TODO(), defaultSubscriberID, "95a3b95b-6938-4eee-af82-a3a7e42878a6")
	if err != nil {
		panic(err)
	}

	fmt.Println("done")
}
