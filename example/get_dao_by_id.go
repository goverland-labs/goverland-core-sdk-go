package main

import (
	"context"
	"fmt"

	goverlandcorewebsdk "github.com/goverland-labs/core-web-sdk"
)

func getDaoByID() {
	cli := goverlandcorewebsdk.NewClient(defaultBaseURL, defaultSubscriberID, "")
	dao, err := cli.GetDao(context.TODO(), "266bf267-cbd8-423d-a145-3d908c47684b")
	if err != nil {
		panic(err)
	}

	fmt.Println(dao.ID, dao.Name, dao.Categories)
}
