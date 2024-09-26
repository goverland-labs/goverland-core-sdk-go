package main

import (
	"context"
	"fmt"

	goverlandcorewebsdk "github.com/goverland-labs/goverland-core-sdk-go"
)

func getAllDelegatesByAddress() {
	cli := goverlandcorewebsdk.NewClient(defaultBaseURL)
	resp, err := cli.GetAllDelegationsByAddress(context.TODO(), "0x7be08c5E75fa1Ca9416E29Cbf58D61D856b01a8C")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
