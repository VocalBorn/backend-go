package models

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	AccountID uuid.UUID `gorm:"type:uuid;primaryKey" json:"account_id"`
	Password  string    `gorm:"type:text;not null" json:"password"`
	Email     string    `gorm:"type:varchar(255);unique;not null" json:"email"`
	CreatedAt time.Time `gorm:"column:created_at;default:current_timestamp;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:current_timestamp;not null" json:"updated_at"`
	User      User      `gorm:"foreignKey:AccountID;references:AccountID" json:"user,omitempty"`
}

// TableName 設定資料表名稱
func (Account) TableName() string {
	return "accounts"
}

type AccountRegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
