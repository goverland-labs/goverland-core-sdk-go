package dao

import "github.com/google/uuid"

type List struct {
	Items    []Dao
	TotalCnt int
	Offset   int
	Limit    int
}

type DaoIds struct {
	Ids      []uuid.UUID
	TotalCnt int
}
