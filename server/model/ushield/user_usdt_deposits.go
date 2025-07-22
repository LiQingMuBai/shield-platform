// 自动生成模板UserUsdtDeposits
package ushield

import (
	"time"
)

// userUsdtDeposits表 结构体  UserUsdtDeposits
type UserUsdtDeposits struct {
	Id          int64     `json:"id" form:"id" gorm:"primarykey;column:id;size:20;"`                    //id字段
	CreatedAt   time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`                 //createdAt字段
	UpdatedAt   time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`                 //updatedAt字段
	DeletedAt   time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;"`                 //deletedAt字段
	UserId      int64     `json:"userId" form:"userId" gorm:"column:user_id;"`                          //userId字段
	Status      int64     `json:"status" form:"status" gorm:"column:status;"`                           //status字段
	Placeholder string    `json:"placeholder" form:"placeholder" gorm:"column:placeholder;size:3;"`     //placeholder字段
	Address     string    `json:"address" form:"address" gorm:"comment:地址;column:address;size:191;"`    //地址
	TxHash      string    `json:"txHash" form:"txHash" gorm:"comment:tx_hash;column:tx_hash;size:191;"` //tx_hash
	Amount      string    `json:"amount" form:"amount" gorm:"comment:金额;column:amount;size:191;"`       //金额
	Block       string    `json:"block" form:"block" gorm:"comment:区块;column:block;size:191;"`          //区块
	OrderNo     string    `json:"orderNo" form:"orderNo" gorm:"column:order_no;size:100;"`              //orderNo字段
}

// TableName userUsdtDeposits表 UserUsdtDeposits自定义表名 user_usdt_deposits
func (UserUsdtDeposits) TableName() string {
	return "user_usdt_deposits"
}
