package proposal

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Vote struct {
	ID                    string          `json:"id"`
	Ipfs                  string          `json:"ipfs"`
	Voter                 string          `json:"voter"`
	EnsName               string          `json:"ens_name"`
	DaoID                 uuid.UUID       `json:"dao_id"`
	ProposalID            string          `json:"proposal_id"`
	Choice                json.RawMessage `json:"choice"`
	Created               uint64          `json:"created"`
	Reason                string          `json:"reason"`
	App                   string          `json:"app"`
	VotingPower           float64         `json:"vp"`
	VotingPowerState      string          `json:"vp_state"`
	VotingPowerByStrategy []float32       `json:"vp_by_strategy"`
}

type VoteValidation struct {
	OK                  bool                 `json:"ok"`
	VotingPower         float64              `json:"voting_power"`
	VoteValidationError *VoteValidationError `json:"error,omitempty"`
}

type VoteValidationError struct {
	Message string `json:"message"`
	Code    uint32 `json:"code"`
}

type VotePreparation struct {
	ID        string `json:"id"`
	TypedData string `json:"typed_data"`
}

type SuccessfulVote struct {
	ID      string  `json:"id"`
	IPFS    string  `json:"ipfs"`
	Relayer Relayer `json:"relayer"`
}

type Relayer struct {
	Address string `json:"address"`
	Receipt string `json:"receipt"`
}
