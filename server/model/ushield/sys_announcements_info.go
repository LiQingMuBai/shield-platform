
// 自动生成模板SysAnnouncementsInfo
package ushield
import (
	"time"
)

// sysAnnouncementsInfo表 结构体  SysAnnouncementsInfo
type SysAnnouncementsInfo struct {
  Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;size:20;"`  //id字段
  Title  *string `json:"title" form:"title" gorm:"comment:公告标题;column:title;size:191;"`  //公告标题
  Content  *string `json:"content" form:"content" gorm:"comment:公告内容;column:content;"`  //公告内容
  CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`  //createdAt字段
  UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`  //updatedAt字段
  DeletedAt  *time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;"`  //deletedAt字段
}


// TableName sysAnnouncementsInfo表 SysAnnouncementsInfo自定义表名 sys_announcements_info
func (SysAnnouncementsInfo) TableName() string {
    return "sys_announcements_info"
}





