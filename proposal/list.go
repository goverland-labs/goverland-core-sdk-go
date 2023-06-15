package proposal

type List struct {
	Items    []Proposal
	TotalCnt int
	Offset   int
	Limit    int
}

type VoteList struct {
	Items    []Vote
	TotalCnt int
	Offset   int
	Limit    int
}
