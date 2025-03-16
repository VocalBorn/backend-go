package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID    uuid.UUID `gorm:"type:uuid;primaryKey" json:"user_id"`
	AccountID uuid.UUID `gorm:"type:uuid;not null;unique;foreignKey:account_id;references:account_id" json:"account_id"` //外鍵 關聯Account
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Gender    string    `gorm:"type:varchar(10);" json:"gender"`
	Age       int       `gorm:"type:integer;" json:"age"`
	Phone     string    `gorm:"type:varchar(20);" json:"phone"`
	CreatedAt time.Time `gorm:"column:created_at;default:current_timestamp;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:current_timestamp;not null" json:"updated_at"`
}

// TableName 設定資料表名稱
func (User) TableName() string {
	return "users"
}

