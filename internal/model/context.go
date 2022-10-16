package model

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type Context struct {
	Session       *ghttp.Session        // Session in context.
	ServiceResult *ContextServiceResult // User in context.
}

type ContextUser struct {
	Id       uint   // User ID.
	Passport string // User passport.
	Nickname string // User nickname.
}

type ContextAccount struct {
	Id       int    // User ID.
	Username string // User nickname.
}

type ContextServiceResult struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}
