package models

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/astaxie/beego"
)

type User struct {
	Model
	Account  string `gorm:"not null;unique" json:"account"`
	Password string `gorm:"not null" json:"password,omitempty"`
	UserName string `json:"username,omitempty"`
	Tag      string `json:"tag,omitempty"`
}

func PasswordMd5(password string) string {
	var hash = md5.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

func (user *User) CheckPassword(password string, isMd5 bool) bool {
	if user.Password == "" {
		return false
	}
	if isMd5 {
		return user.Password == password
	} else {
		return user.Password == PasswordMd5(password)
	}
}

type UserQuery struct{}

func (u *UserQuery) GetUserByAccount(account string) (*User, error) {
	user := &User{}
	user.Account = account
	if db, err := DB(); err != nil {
		beego.Error(err)
		return nil, err
	} else {
		defer db.Close()
		db.Select("account,password").Where("account = ?", account).First(user)
		return user, nil
	}
}
