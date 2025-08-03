// 自动生成模板UserAddressMonitorEvent
package ushield

import (
	"time"
)

// userAddressMonitorEvent表 结构体  UserAddressMonitorEvent
type UserAddressMonitorEvent struct {
	Id                int64     `json:"id" form:"id" gorm:"primarykey;column:id;size:20;"`                                      //id字段
	CreatedAt         time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`                                   //createdAt字段
	UpdatedAt         time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`                                   //updatedAt字段
	DeletedAt         time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;"`                                   //deletedAt字段
	Days              int64     `json:"days" form:"days" gorm:"column:days;size:19;"`                                           //days字段
	ChatId            int64     `json:"chatId" form:"chatId" gorm:"column:chat_id;size:19;"`                                    //chatId字段
	Status            int64     `json:"status" form:"status" gorm:"column:status;size:19;"`                                     //status字段
	Times             int64     `json:"times" form:"times" gorm:"column:times;size:19;"`                                        //status字段
	InsufficientTimes int64     `json:"insufficient_times" form:"insufficient_times" gorm:"column:insufficient_times;size:19;"` //status字段
	Network           string    `json:"network" form:"network" gorm:"column:network;size:10;"`                                  //network字段
	Address           string    `json:"address" form:"address" gorm:"column:address;size:191;"`                                 //address字段
}

// TableName userAddressMonitorEvent表 UserAddressMonitorEvent自定义表名 user_address_monitor_event
func (UserAddressMonitorEvent) TableName() string {
	return "user_address_monitor_event"
}
