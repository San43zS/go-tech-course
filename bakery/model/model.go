package model

const (
	K  = 10000
	N  = 8
	M  = 5
	T1 = 100
	T2 = 150
)

type Cake struct {
	ID       int
	BakedBy  int
	BakeTime int
	PackedBy int
	PackTime int
}
