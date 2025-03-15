package test

import (
	"testing"
	"vocalborn/backend-go/models"
	"vocalborn/backend-go/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	_ "github.com/joho/godotenv/autoload"
)

func TestCreateAccount(t *testing.T) {
	db := utils.ConnectDB()

	// 測試前清理數據
	// db.Exec("DELETE FROM users")
	// db.Exec("DELETE FROM accounts")

	// 建立帳號
	accountID := uuid.New()
	account := models.Account{
		AccountID: accountID,
		Email:     "test@example.com",
		Password:  "securepassword",
	}

	result := db.Create(&account)
	assert.NoError(t, result.Error)
	assert.Equal(t, int64(1), result.RowsAffected)

	// 驗證帳號已建立
	var fetchedAccount models.Account
	result = db.First(&fetchedAccount, "email = ?", "test@example.com")
	assert.NoError(t, result.Error)
	assert.Equal(t, account.Email, fetchedAccount.Email)
}

func TestCreateUser(t *testing.T) {
	db := utils.ConnectDB()

	// 測試前清理數據
	// db.Exec("DELETE FROM users")
	// db.Exec("DELETE FROM accounts")

	// 先建立帳號
	accountID := uuid.New()
	db.Create(&models.Account{
		AccountID: accountID,
		Email:     "test@example.com",
		Password:  "securepassword",
	})

	// 建立使用者
	user := models.User{
		UserID:    uuid.New(),
		AccountID: accountID,
		Name:      "測試用戶",
		Gender:    "男",
		Age:       25,
		Phone:     "0912345678",
	}

	result := db.Create(&user)
	assert.NoError(t, result.Error)

	// 驗證使用者已建立
	// var fetchedUser models.User
	// result = db.Preload("Account").First(&fetchedUser, "name = ?", "測試用戶")
	// assert.NoError(t, result.Error)
	// assert.Equal(t, user.Name, fetchedUser.Name)
	// assert.Equal(t, accountID, fetchedUser.AccountID)
}
