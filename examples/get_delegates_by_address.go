package main

import (
	"context"
	"fmt"

	goverlandcorewebsdk "github.com/goverland-labs/goverland-core-sdk-go"
)

func getTopDelegatesByAddress() {
	cli := goverlandcorewebsdk.NewClient(defaultBaseURL)
	resp, err := cli.GetTopDelegatesByAddress(context.TODO(), "0x7be08c5E75fa1Ca9416E29Cbf58D61D856b01a8C")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}

func getDelegatesList() {
	cli := goverlandcorewebsdk.NewClient(defaultBaseURL)
	resp, err := cli.GetDelegatesList(context.TODO(), goverlandcorewebsdk.GetDelegatesListRequest{
		DaoID:   "0f08144e-fe19-4098-baf7-88adf1f9a428",
		Address: "0x7be08c5E75fa1Ca9416E29Cbf58D61D856b01a8C",
		Limit:   10,
		Offset:  0,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}

func getTopDelegatorsByAddress() {
	cli := goverlandcorewebsdk.NewClient(defaultBaseURL)
	resp, err := cli.GetTopDelegatorsByAddress(context.TODO(), "0x6cc5b30cd0a93c1f85c7868f5f2620ab8c458190")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}

func getDelegatorsList() {
	cli := goverlandcorewebsdk.NewClient(defaultBaseURL)
	resp, err := cli.GetDelegatorsList(context.TODO(), goverlandcorewebsdk.GetDelegatorsListRequest{
		DaoID:   "755f2283-8384-4987-9f76-49089fde1b24",
		Address: "0x6cc5b30cd0a93c1f85c7868f5f2620ab8c458190",
		Limit:   10,
		Offset:  1,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}

func getTotalDelegationsByAddress() {
	cli := goverlandcorewebsdk.NewClient(defaultBaseURL)
	resp, err := cli.GetTotalDelegationsByAddress(context.TODO(), "0x7be08c5E75fa1Ca9416E29Cbf58D61D856b01a8C")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
