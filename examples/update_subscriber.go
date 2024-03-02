package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	goverlandcorewebsdk "github.com/goverland-labs/goverland-core-sdk-go"
)

func updateSubscriber() {
	cli := goverlandcorewebsdk.NewClient(defaultBaseURL)
	err := cli.UpdateSubscriber(context.TODO(), uuid.MustParse("4fe6b578-b34b-402a-8b88-bb0f83041119"), "http://new.callback.url")
	if err != nil {
		panic(err)
	}

	fmt.Println("updated")
}
