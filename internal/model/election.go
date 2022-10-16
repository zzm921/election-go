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

// 更新选举信息
type ElectionUpdateInput struct {
	ElectionId   int
	Title        string
	Introduction string
	Candidates   []int
}

type ElectionUpdateOut struct {
	ElectionId   int
	Title        string
	Introduction string
	Candidates   []int
}

// 修改选举状态
type ElectionChangeStatuInput struct {
	ElectionId int
	Status     string
}

type ElectionChangeStatuOut struct {
	ElectionId int
	Status     string
}
