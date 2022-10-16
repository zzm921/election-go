// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Candidates is the golang structure of table candidates for DAO operations like Where/Data.
type Candidates struct {
	g.Meta       `orm:"table:candidates, do:true"`
	Id           interface{} //
	Name         interface{} // 候选人名称
	Introduction interface{} // 候选人简介
	Image        interface{} // 候选人简介
	Status       interface{} // 候选人状态 1 - 可用 0 - 不可用
	Createtime   *gtime.Time // 创建时间
	Updatetime   *gtime.Time // 更新时间
}