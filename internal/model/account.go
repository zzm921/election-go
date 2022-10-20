package model

type AccountLoginInput struct {
	Username string
	Password string
}

type AccountLoginOut struct {
	Id       int    `json:"id"          description:"账号"`
	Token    string `json:"token"          description:"token"`
	Username string `json:"username"          description:"账号名"`
	Role     int    `json:"role"          description:"角色"`
}
