package model

type AccountLoginInput struct {
	Username string
	Password string
}

type AccountLoginOut struct {
	Id       int
	Token    string
	Username string
	Role     int
}

type AccountCreateInput struct {
	Username string
	Password string
	Role     string
}

type AccountCreateOut struct {
	Username string
	Role     string
}
