package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	goverlandcorewebsdk "github.com/goverland-labs/goverland-core-sdk-go"
)

func unsubscribeFromDao() {
	cli := goverlandcorewebsdk.NewClient(defaultBaseURL)
	err := cli.UnsubscribeFromDao(context.TODO(), defaultSubscriberID, uuid.MustParse("95a3b95b-6938-4eee-af82-a3a7e42878a6"))
	if err != nil {
		panic(err)
	}

	fmt.Println("done")
}
