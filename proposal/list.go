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
	TotalVp  float32
	Offset   int
	Limit    int
}
