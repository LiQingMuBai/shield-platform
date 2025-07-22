// 自动生成模板TgUsers
package ushield

import (
	"time"
)

// tgUsers表 结构体  TgUsers
type TgUsers struct {
	Id           int64     `json:"id" form:"id" gorm:"primarykey;column:id;size:19;"`                       //id字段
	Username     string    `json:"username" form:"username" gorm:"column:username;size:255;"`               //username字段
	Associates   string    `json:"associates" form:"associates" gorm:"column:associates;size:255;"`         //associates字段
	BackupChatId string    `json:"backupChatId" form:"backupChatId" gorm:"column:backup_chat_id;size:100;"` //backupChatId字段
	Amount       string    `json:"amount" form:"amount" gorm:"column:amount;size:255;"`                     //amount字段
	TronAmount   string    `json:"tronAmount" form:"tronAmount" gorm:"column:tron_amount;size:255;"`        //tronAmount字段
	TronAddress  string    `json:"tronAddress" form:"tronAddress" gorm:"column:tron_address;size:50;"`      //tronAddress字段
	EthAddress   string    `json:"ethAddress" form:"ethAddress" gorm:"column:eth_address;size:50;"`         //ethAddress字段
	EthAmount    string    `json:"ethAmount" form:"ethAmount" gorm:"column:eth_amount;size:255;"`           //ethAmount字段
	CreatedAt    time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`                    //createdAt字段
	Deadline     time.Time `json:"deadline" form:"deadline" gorm:"column:deadline;"`                        //deadline字段
	UpdatedAt    time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`                    //updatedAt字段
	UserId       string    `json:"userId" form:"userId" gorm:"column:user_id;size:255;"`                    //userId字段
	Times        int64     `json:"times" form:"times" gorm:"column:times;size:10;"`                         //times字段
	Address      string    `json:"address" form:"address" gorm:"column:address;size:100;"`                  //address字段
	PrivateKey   string    `json:"privateKey" form:"privateKey" gorm:"column:private_key;size:200;"`        //privateKey字段
}

// TableName tgUsers表 TgUsers自定义表名 tg_users
func (TgUsers) TableName() string {
	return "tg_users"
}
