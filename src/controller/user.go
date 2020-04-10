package controller

import (
	"github.com/gin-gonic/gin"
	"go-blog/src/common"
	"go-blog/src/model"
	"go-blog/src/service"
	"net/http"
)

// ========================= 登录
func LoginGet(c *gin.Context) {
	user := common.GetSession(common.ContextUserKey, c)
	userInfo, _ := user.(model.UserInfo)
	if (model.UserInfo{}) != userInfo {
		c.HTML(http.StatusOK, "index/index.html", nil)
		return
	}
	c.HTML(http.StatusOK, "auth/login.html", nil)
}
func LoginPost(c *gin.Context) {
	userInfo := check(c)
	c.JSON(http.StatusOK, service.Login(userInfo, c))
}

func check(c *gin.Context) model.UserInfo {
	username := c.PostForm("username")
	password := c.PostForm("password")
	userInfo := model.UserInfo{
		Username: username,
		PassWord: password,
	}
	return userInfo
}

// ==========================  注册
func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "auth/register.html", nil)
}
func RegisterPost(c *gin.Context) {
	userInfo := check(c)
	c.JSON(http.StatusOK, service.Register(userInfo))
}

// =========================  登出
func Logout(c *gin.Context) {
	common.DelSession(common.ContextUserKey, c)
	c.HTML(http.StatusOK, "index/index.html", nil)

}
