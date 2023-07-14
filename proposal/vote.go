package proposal

import (
	"github.com/google/uuid"
)

type Vote struct {
	ID               string    `json:"id"`
	Ipfs             string    `json:"ipfs"`
	Voter            string    `json:"voter"`
	DaoID            uuid.UUID `json:"dao_id"`
	ProposalID       string    `json:"proposal_id"`
	Choice           int       `json:"choice"`
	Created          uint64    `json:"created"`
	Reason           string    `json:"reason"`
	VotingPower      float64   `json:"vp"`
	VotingPowerState string    `json:"vp_state"`
}
