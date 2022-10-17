package model

//创建选举
type ElectionCreateInput struct {
	Title        string
	Introduction string
	Candidates   []int
}

type ElectionCreateOut struct {
	Title        string
	Introduction string
	Candidates   []int
}

// 修改选举状态
type ElectionChangeStatuInput struct {
	ElectionId int
	Status     int
}

type ElectionChangeStatuOut struct {
}

// 修改选举状态
type ElectionGetInput struct {
	Page int
	Size int
}

type ElectionGetOutListObject struct {
	Id           int
	Title        string
	Introduction string
	Status       int
	Candidates   []int
}

type ElectionGetOut struct {
	Count int
	List  []*ElectionGetOutListObject
}
