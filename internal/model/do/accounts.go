// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Accounts is the golang structure of table accounts for DAO operations like Where/Data.
type Accounts struct {
	g.Meta     `orm:"table:accounts, do:true"`
	Id         interface{} //
	Username   interface{} // 用户名
	Password   interface{} // 密码
	Role       interface{} // 账号角色 0 - 超级管理员 1 - 管理员
	Status     interface{} // 账号状态 ：  0 - 禁用       1 - 可用
	Createtime *gtime.Time // 创建时间
	Updatetime *gtime.Time // 更新时间
}