package models

import (
	"gin-blog/src/gin-blog/pkg/logging"
	"gin-blog/src/gin-blog/pkg/util"
	"github.com/jinzhu/gorm"
	"time"
)

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthAccount struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		logging.Info("验证成功")
		return true
	}
	return false
}

func CreateAuthAccount(username string, password string) bool {
	err := db.Create(&Auth{
		Username: username,
		Password: util.Encode(password),
	})
	if err != nil {
		logging.Error("创建失败")
		return false
	} else {
		return true
	}
}

func (auth *Auth) BeforeCreated(scope *gorm.Scope) error {
	scope.SetColumn("Created", time.Now().Unix())

	return nil
}

func (auth *Auth) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("Updated", time.Now().Unix())

	return nil
}

func CheckAuthName(username string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username}).First(&auth)
	if auth.ID > 0 {
		logging.Info("验证成功")
		return false
	}

	return true
}
