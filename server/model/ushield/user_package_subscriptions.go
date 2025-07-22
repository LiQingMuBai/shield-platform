
// 自动生成模板UserPackageSubscriptions
package ushield
import (
	"time"
)

// userPackageSubscriptions表 结构体  UserPackageSubscriptions
type UserPackageSubscriptions struct {
  Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;size:20;"`  //id字段
  CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`  //createdAt字段
  UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`  //updatedAt字段
  DeletedAt  *time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;"`  //deletedAt字段
  Status  *int `json:"status" form:"status" gorm:"column:status;size:19;"`  //status字段
  BundleId  *int `json:"bundleId" form:"bundleId" gorm:"column:bundle_id;size:19;"`  //bundleId字段
  UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;size:19;"`  //userId字段
}


// TableName userPackageSubscriptions表 UserPackageSubscriptions自定义表名 user_package_subscriptions
func (UserPackageSubscriptions) TableName() string {
    return "user_package_subscriptions"
}





