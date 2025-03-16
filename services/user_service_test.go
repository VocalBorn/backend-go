package services

import (
	"context"
	"testing"
	"vocalborn/backend-go/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 設置測試環境
func setupTestUserService() (*UserService, sqlmock.Sqlmock, error) {
	// 創建 SQL mock
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	// 使用 mock 創建 GORM 連接
	dialector := postgres.New(postgres.Config{
		Conn:       mockDB,
		DriverName: "postgres",
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	// 創建自定義的 UserService 並注入模擬資料庫
	userService := &UserService{db: db}

	return userService, mock, nil
}

// 測試 Register 方法
func TestRegister(t *testing.T) {
	// 設置環境
	userService, mock, err := setupTestUserService()
	assert.NoError(t, err)

	ctx := context.Background()
	req := &models.AccountRegisterRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	// 設置預期行為 - 檢查郵箱是否已存在
	emailCheckRows := sqlmock.NewRows([]string{"user_id", "account_id", "name"})
	mock.ExpectQuery(`SELECT .* FROM "users" WHERE email = \$1 ORDER BY .* LIMIT \$2`).
		WithArgs("test@example.com", 1).
		WillReturnRows(emailCheckRows)

	// 設置預期行為 - 開始事務
	mock.ExpectBegin()

	// 設置預期行為 - 創建 Account
	accountID := uuid.New()
	mock.ExpectQuery(`INSERT INTO "accounts"`).
		WillReturnRows(sqlmock.NewRows([]string{"account_id"}).AddRow(accountID))

	// 設置預期行為 - 創建 User
	mock.ExpectQuery(`INSERT INTO "users"`).
		WillReturnRows(sqlmock.NewRows([]string{"user_id"}).AddRow(uuid.New()))

	// 設置預期行為 - 提交事務
	mock.ExpectCommit()

	// 執行測試
	account, err := userService.Register(ctx, req)

	// 驗證結果
	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, req.Email, account.Email)

	// 驗證所有預期的SQL已執行
	assert.NoError(t, mock.ExpectationsWereMet())
}

// 測試 FindByEmail 方法
func TestFindByEmail(t *testing.T) {
	userService, mock, err := setupTestUserService()
	assert.NoError(t, err)

	ctx := context.Background()
	email := "test@example.com"
	userID := uuid.New()
	accountID := uuid.New()

	// 設置預期行為
	rows := sqlmock.NewRows([]string{"user_id", "account_id", "name"}).
		AddRow(userID, accountID, "Test User")
	mock.ExpectQuery(`SELECT .* FROM "users" WHERE email = \$1 ORDER BY .* LIMIT \$2`).
		WithArgs(email, 1). // 添加第二個參數，通常GORM的LIMIT預設值是1
		WillReturnRows(rows)

	// 執行測試
	user, err := userService.FindByEmail(ctx, email)

	// 驗證結果
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, userID, user.UserID)
	assert.Equal(t, accountID, user.AccountID)

	// 驗證所有預期的SQL已執行
	assert.NoError(t, mock.ExpectationsWereMet())
}
