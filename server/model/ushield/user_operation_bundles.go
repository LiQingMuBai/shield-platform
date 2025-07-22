
// 自动生成模板UserOperationBundles
package ushield
import (
	"time"
)

// userOperationBundles表 结构体  UserOperationBundles
type UserOperationBundles struct {
  Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;size:20;"`  //id字段
  CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`  //createdAt字段
  UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`  //updatedAt字段
  DeletedAt  *time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;"`  //deletedAt字段
  Status  *int `json:"status" form:"status" gorm:"column:status;size:19;"`  //status字段
  Name  *string `json:"name" form:"name" gorm:"column:name;size:30;"`  //name字段
  Amount  *string `json:"amount" form:"amount" gorm:"column:amount;size:191;"`  //amount字段
}


// TableName userOperationBundles表 UserOperationBundles自定义表名 user_operation_bundles
func (UserOperationBundles) TableName() string {
    return "user_operation_bundles"
}





