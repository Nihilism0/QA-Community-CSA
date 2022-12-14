package dao

import (
	"CSAwork/global"
	"CSAwork/model"
)

// 若没有这个用户返回 false，反之返回 true
func SelectUser(username string) bool {
	var u struct {
		Username string
	}
	global.GlobalDb1.Model(&model.User{}).Where("username = ?", username).Find(&u)
	if u.Username == "" {
		return false
	} else {
		return true
	}
}
