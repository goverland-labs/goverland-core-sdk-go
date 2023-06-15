package dao

type List struct {
	Items    []Dao
	TotalCnt int
	Offset   int
	Limit    int
}
