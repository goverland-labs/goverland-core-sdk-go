package subscriber

import (
	"github.com/google/uuid"
)

type Subscriber struct {
	SubscriberID uuid.UUID `json:"subscriber_id"`
}
