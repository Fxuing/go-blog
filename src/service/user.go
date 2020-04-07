package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/src/common"
	"go-blog/src/model"
)

/**
注册
*/
func Register(userinfo model.UserInfo) common.ResponseData {
	users, err := model.FindByUserName(userinfo)
	if err != nil {
		fmt.Println(err)
	}
	if len(users) > 0 {
		return common.ResponseData{
			Code: "500",
			Msg:  "用户名已被注册了",
			Data: nil,
		}
	}
	err = model.UserAdd(userinfo)
	if err != nil {
		fmt.Print(err)
		return common.ResponseData{
			Code: "500",
			Msg:  err.Error(),
			Data: nil,
		}
	}
	return common.ResponseData{
		Code: "1",
		Msg:  "成功",
		Data: nil,
	}
}

/**
登录
*/
func Login(userinfo model.UserInfo, c *gin.Context) common.ResponseData {
	users, err := model.FindByUserName(userinfo)
	if err != nil {
		fmt.Print(err)
	}
	if len(users) == 0 {
		return common.ResponseData{
			Code: "500",
			Msg:  "用户名未注册",
		}
	}
	user, err := model.FindByUserNameAndPassWord(userinfo)
	if err != nil {
		fmt.Print(err)
		var msg string
		if user.ID == 0 {
			msg = "您输入的密码有误，请重新输入"
		} else {
			msg = err.Error()
		}
		return common.ResponseData{
			Code: "500",
			Msg:  msg,
		}
	}
	if user.ID == 0 {
		return common.ResponseData{
			Code: "500",
			Msg:  "您输入的密码有误，请重新输入",
		}
	}
	common.SetSession(common.ContextUserKey, user, c)
	//c.Set(common.ContextUserKey, user)
	return common.ResponseData{
		Code: "1",
		Msg:  "成功",
	}
}
