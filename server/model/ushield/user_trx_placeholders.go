
// 自动生成模板UserTrxPlaceholders
package ushield
import (
	"time"
)

// userTrxPlaceholders表 结构体  UserTrxPlaceholders
type UserTrxPlaceholders struct {
  Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;size:20;"`  //id字段
  CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`  //createdAt字段
  UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`  //updatedAt字段
  DeletedAt  *time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;"`  //deletedAt字段
  Status  *int `json:"status" form:"status" gorm:"column:status;size:19;"`  //status字段
  Placeholder  *string `json:"placeholder" form:"placeholder" gorm:"column:placeholder;size:30;"`  //placeholder字段
}


// TableName userTrxPlaceholders表 UserTrxPlaceholders自定义表名 user_trx_placeholders
func (UserTrxPlaceholders) TableName() string {
    return "user_trx_placeholders"
}





