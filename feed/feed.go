package feed

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID           uuid.UUID       `json:"id"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	DaoID        uuid.UUID       `json:"dao_id"`
	ProposalID   string          `json:"proposal_id"`
	DiscussionID string          `json:"discussion_id"`
	Type         string          `json:"type"`
	Action       string          `json:"action"`
	Snapshot     json.RawMessage `json:"snapshot"`
	Timeline     json.RawMessage `json:"timeline"`
}

type Feed struct {
	Items    []Item
	TotalCnt int
	Offset   int
	Limit    int
}
