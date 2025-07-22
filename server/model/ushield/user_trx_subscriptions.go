
// 自动生成模板UserTrxSubscriptions
package ushield
import (
	"time"
)

// userTrxSubscriptions表 结构体  UserTrxSubscriptions
type UserTrxSubscriptions struct {
  Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;size:20;" binding:"required"`  //id字段
  CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`  //createdAt字段
  UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`  //updatedAt字段
  DeletedAt  *time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;"`  //deletedAt字段
  Status  *int `json:"status" form:"status" gorm:"column:status;"`  //status字段
  Name  *string `json:"name" form:"name" gorm:"column:name;size:30;"`  //name字段
  Amount  *string `json:"amount" form:"amount" gorm:"column:amount;size:191;"`  //amount字段
}


// TableName userTrxSubscriptions表 UserTrxSubscriptions自定义表名 user_trx_subscriptions
func (UserTrxSubscriptions) TableName() string {
    return "user_trx_subscriptions"
}





