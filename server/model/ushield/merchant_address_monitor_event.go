// 自动生成模板MerchantAddressMonitorEvent
package ushield

import (
	"time"
)

// merchantAddressMonitorEvent表 结构体  MerchantAddressMonitorEvent
type MerchantAddressMonitorEvent struct {
	Id                int64     `json:"id" form:"id" gorm:"primarykey;column:id;size:20;"`                                    //id字段
	Amount            string    `json:"amount" form:"amount" gorm:"column:amount;size:20;"`                                   //amount字段
	Days              int64     `json:"days" form:"days" gorm:"column:days;"`                                                 //days字段
	Times             int64     `json:"times" form:"times" gorm:"column:times;size:10;"`                                      //times字段
	InsufficientTimes int64     `json:"insufficientTimes" form:"insufficientTimes" gorm:"column:insufficient_times;size:10;"` //insufficientTimes字段
	ChatId            int64     `json:"chatId" form:"chatId" gorm:"column:chat_id;size:19;"`                                  //chatId字段
	UserId            uint      `json:"userId" form:"userId" gorm:"column:user_id;size:19;"`                                  //userId字段
	Status            int64     `json:"status" form:"status" gorm:"column:status;"`                                           //status字段
	Network           string    `json:"network" form:"network" gorm:"column:network;size:10;"`                                //network字段
	Address           string    `json:"address" form:"address" gorm:"column:address;size:191;"`                               //address字段
	Callback          string    `json:"callback_url" form:"callback_url" gorm:"column:callback_url;size:191;"`                //address字段
	DeletedAt         time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;"`                                 //deletedAt字段
	CreatedAt         time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`                                 //createdAt字段
	UpdatedAt         time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`                                 //updatedAt字段
	CreatedDate       string    `json:"createdDate" form:"createdDate" gorm:"column:created_date;size:20;"`                   //createdDate字段
}

// TableName merchantAddressMonitorEvent表 MerchantAddressMonitorEvent自定义表名 merchant_address_monitor_event
func (MerchantAddressMonitorEvent) TableName() string {
	return "merchant_address_monitor_event"
}
