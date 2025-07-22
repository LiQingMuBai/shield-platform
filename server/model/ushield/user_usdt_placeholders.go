
// 自动生成模板UserUsdtPlaceholders
package ushield
import (
	"time"
)

// userUsdtPlaceholders表 结构体  UserUsdtPlaceholders
type UserUsdtPlaceholders struct {
  Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;size:20;"`  //id字段
  CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`  //createdAt字段
  UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`  //updatedAt字段
  DeletedAt  *time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;"`  //deletedAt字段
  Status  *int `json:"status" form:"status" gorm:"column:status;"`  //status字段
  Placeholder  *string `json:"placeholder" form:"placeholder" gorm:"column:placeholder;size:30;"`  //placeholder字段
}


// TableName userUsdtPlaceholders表 UserUsdtPlaceholders自定义表名 user_usdt_placeholders
func (UserUsdtPlaceholders) TableName() string {
    return "user_usdt_placeholders"
}





