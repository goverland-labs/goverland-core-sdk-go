package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	goverlandcorewebsdk "github.com/goverland-labs/core-web-sdk"
)

func getDaoByID() {
	cli := goverlandcorewebsdk.NewClient(defaultBaseURL)
	dao, err := cli.GetDao(context.TODO(), uuid.MustParse("266bf267-cbd8-423d-a145-3d908c47684b"))
	if err != nil {
		panic(err)
	}

	fmt.Println(dao.ID, dao.Name, dao.Categories)
}
