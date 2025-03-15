package services

import (
	"context"
	"errors"
	"vocalborn/backend-go/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserService 處理與用戶相關的業務邏輯
type UserService struct {
	db *gorm.DB
}

// 用戶服務的構造函數
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// FindByEmail 根據電子郵件查詢用戶
func (userSrv *UserService) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	result := userSrv.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// Create 創建新用戶
func (userSrv *UserService) Create(ctx context.Context, user *models.Account) error {
	return userSrv.db.Create(user).Error
}

func (userSrv *UserService) Register(ctx context.Context, req *models.AccountRegisterRequest) (*models.Account, error) {
	// 檢查用戶是否已存在
	var existingUser models.User
	result := userSrv.db.Where("email = ?", req.Email).First(&existingUser)

	// 只有當錯誤不是"記錄未找到"時才返回錯誤
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	// 如果找到用戶，表示郵箱已註冊
	if result.RowsAffected > 0 {
		return nil, errors.New("email already registered")
	}

	// 密碼加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// 在一個事務中創建 Account 和 User
	var account models.Account
	err = userSrv.db.Transaction(func(tx *gorm.DB) error {
		// 創建 Account
		account = models.Account{
			AccountID: uuid.New(),
			Email:     req.Email,
			Password:  string(hashedPassword),
		}
		if err := tx.Create(&account).Error; err != nil {
			return err
		}

		// 創建 User
		user := models.User{
			UserID:    uuid.New(),
			AccountID: account.AccountID,
		}
		return tx.Create(&user).Error
	})
	if err != nil {
		return nil, err
	}

	// 隱藏密碼後返回用戶資訊
	account.Password = ""
	return &account, nil
}
