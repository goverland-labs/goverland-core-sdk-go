package main

import (
	"github.com/google/uuid"
)

const defaultBaseURL = "http://localhost:88/v1/"

var defaultSubscriberID = uuid.MustParse("b3be82c6-7ccc-4db6-9506-4d0c9c36ab76")

func main() {
	updateSubscriber()
}
