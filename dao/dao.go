package dao

import (
	"time"

	"github.com/google/uuid"
)

type Categories []string

type Strategy struct {
	Name    string                 `json:"name"`
	Network string                 `json:"network"`
	Params  map[string]interface{} `json:"params"`
}

type Strategies []Strategy

type Treasury struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Network string `json:"network"`
}

type Treasuries []Treasury

type Voting struct {
	Delay       uint64  `json:"delay"`
	Period      uint64  `json:"period"`
	Type        string  `json:"type"`
	Quorum      float32 `json:"quorum"`
	Blind       bool    `json:"blind"`
	HideAbstain bool    `json:"hide_abstain"`
	Privacy     string  `json:"privacy"`
	Aliased     bool    `json:"aliased"`
}

type Dao struct {
	ID              uuid.UUID  `json:"id"`
	Alias           string     `json:"alias"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	Name            string     `json:"name"`
	Private         bool       `json:"private"`
	About           string     `json:"about"`
	Avatar          string     `json:"avatar"`
	Terms           string     `json:"terms"`
	Location        string     `json:"location"`
	Website         string     `json:"website"`
	Twitter         string     `json:"twitter"`
	Github          string     `json:"github"`
	Coingecko       string     `json:"coingecko"`
	Email           string     `json:"email"`
	Network         string     `json:"network"`
	Symbol          string     `json:"symbol"`
	Skin            string     `json:"skin"`
	Domain          string     `json:"domain"`
	Strategies      Strategies `json:"strategies"`
	Voting          Voting     `json:"voting"`
	Categories      Categories `json:"categories"`
	Treasures       Treasuries `json:"treasures"`
	FollowersCount  uint64     `json:"followers_count"`
	ProposalsCount  uint64     `json:"proposals_count"`
	Guidelines      string     `json:"guidelines"`
	Template        string     `json:"template"`
	ActivitySince   uint64     `json:"activity_since"`
	VotersCount     uint64     `json:"voters_count"`
	ActiveVotes     uint64     `json:"active_votes"`
	Verified        bool       `json:"verified"`
	PopularityIndex float64    `json:"popularity_index"`

	// TODO: parentID
}

type Recommendation struct {
	OriginalId string `json:"original_id"`
	InternalId string `json:"internal_id"`
	Name       string `json:"name"`
	Symbol     string `json:"symbol"`
	NetworkId  string `json:"network_id"`
	Address    string `json:"address"`
}

type Recommendations []Recommendation

type Delegate struct {
	Address               string  `json:"address"`
	ENSName               string  `json:"ens_name"`
	DelegatorCount        int32   `json:"delegator_count"`
	PercentOfDelegators   float64 `json:"percent_of_delegators"`
	VotingPower           float64 `json:"voting_power"`
	PercentOfVotingPower  float64 `json:"percent_of_voting_power"`
	About                 string  `json:"about"`
	Statement             string  `json:"statement"`
	VotesCount            int32   `json:"votes_count"`
	CreatedProposalsCount int32   `json:"created_proposals_count"`
}

type Delegates []Delegate

type DelegateProfile struct {
	Address              string                `json:"address"`
	VotingPower          float64               `json:"voting_power"`
	IncomingPower        float64               `json:"incoming_power"`
	OutgoingPower        float64               `json:"outgoing_power"`
	PercentOfVotingPower float64               `json:"percent_of_voting_power"`
	PercentOfDelegators  float64               `json:"percent_of_delegators"`
	Delegates            []ProfileDelegateItem `json:"delegates"`
}

type ProfileDelegateItem struct {
	Address        string  `json:"address"`
	ENSName        string  `json:"ens_name"`
	Weight         float64 `json:"weight"`
	DelegatedPower float64 `json:"delegated_power"`
}
