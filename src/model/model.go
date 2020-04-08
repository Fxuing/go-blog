package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type BaseModel struct {
}

type UserInfo struct {
	gorm.Model
	Username  string `gorm:"size:16"` // 用户名
	Gender    int    // 性别 0 女 1 男
	PhoneCode string `gorm:"size:11"` // 手机号
	PassWord  string `gorm:"size:16"` // 密码
}

var DB *gorm.DB

/**
初始化数据库，并建立连接
*/
func InitDB() (db *gorm.DB, err error) {
	//db, err = gorm.Open("mysql", "root:fxuing#@/go_blog?charset=utf8&parseTime=True&loc=Local")
	db, err = gorm.Open("mysql", "root:Zopen2013@(120.78.66.185:13306)/go_blog?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	exist := db.HasTable(UserInfo{})
	if !exist {
		db.CreateTable(UserInfo{})
	}
	DB = db
	return db, err
}

func UserAdd(userinfo UserInfo) error {
	return DB.Create(&userinfo).Error
}
func FindByUserName(userinfo UserInfo) (users []*UserInfo, err error) {
	err = DB.Where("username = ? ", userinfo.Username).Find(&users).Error
	return
}
func FindByUserNameAndPassWord(userinfo UserInfo) (*UserInfo, error) {
	var user = UserInfo{}
	err := DB.Where("username = ? and pass_word = ?", userinfo.Username, userinfo.PassWord).First(&user).Error
	return &user, err
}
