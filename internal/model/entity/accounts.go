// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Accounts is the golang structure for table accounts.
type Accounts struct {
	Id         int         `json:"id"         description:""`
	Username   string      `json:"username"   description:"用户名"`
	Password   string      `json:"password"   description:"密码"`
	Role       int         `json:"role"       description:"账号角色 0 - 超级管理员 1 - 管理员"`
	Status     int         `json:"status"     description:"账号状态 ：  0 - 禁用       1 - 可用"`
	Createtime *gtime.Time `json:"createtime" description:"创建时间"`
	Updatetime *gtime.Time `json:"updatetime" description:"更新时间"`
}
