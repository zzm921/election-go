// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ElectionUserDetails is the golang structure of table election_user_details for DAO operations like Where/Data.
type ElectionUserDetails struct {
	g.Meta      `orm:"table:election_user_details, do:true"`
	Id          interface{} //
	ElectionId  interface{} //
	CandidateId interface{} //
	IdCard      interface{} //
	Email       interface{} //
	Createtime  *gtime.Time //
}