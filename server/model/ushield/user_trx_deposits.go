// 自动生成模板UserTrxDeposits
package ushield

import (
	"time"
)

// userTrxDeposits表 结构体  UserTrxDeposits
type UserTrxDeposits struct {
	Id          int64     `json:"id" form:"id" gorm:"primarykey;column:id;size:20;"`                //id字段
	CreatedAt   time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`             //createdAt字段
	UpdatedAt   time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`             //updatedAt字段
	DeletedAt   time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;"`             //deletedAt字段
	UserId      int64     `json:"userId" form:"userId" gorm:"column:user_id;size:19;"`              //userId字段
	Status      int64     `json:"status" form:"status" gorm:"column:status;size:19;"`               //status字段
	Placeholder string    `json:"placeholder" form:"placeholder" gorm:"column:placeholder;size:3;"` //placeholder字段
	Address     string    `json:"address" form:"address" gorm:"column:address;size:191;"`           //address字段
	TxHash      string    `json:"txHash" form:"txHash" gorm:"column:tx_hash;size:191;"`             //txHash字段
	Amount      string    `json:"amount" form:"amount" gorm:"column:amount;size:191;"`              //amount字段
	Block       string    `json:"block" form:"block" gorm:"column:block;size:191;"`                 //block字段
	OrderNo     string    `json:"orderNo" form:"orderNo" gorm:"column:order_no;size:100;"`          //orderNo字段
}

// TableName userTrxDeposits表 UserTrxDeposits自定义表名 user_trx_deposits
func (UserTrxDeposits) TableName() string {
	return "user_trx_deposits"
}
