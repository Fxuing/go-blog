package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/src/common"
	"go-blog/src/model"
	"net/http"
	"strconv"
)

func NewBlog(c *gin.Context) {
	user := common.GetSession(common.ContextUserKey, c)
	userInfo, _ := user.(model.UserInfo)
	c.HTML(http.StatusOK, "blog/new.html", gin.H{
		"user": userInfo.Username,
	})
}

func SaveBlog(c *gin.Context) {
	blogInfo := model.BlogInfo{
		Title:    c.PostForm("title"),
		BlogBody: c.PostForm("blogBody"),
	}
	err := model.BlogInfoAdd(blogInfo)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, common.ResponseData{
		Code: "1",
		Msg:  "成功",
	})
}
func DetailBlog(c *gin.Context) {
	user := common.GetSession(common.ContextUserKey, c)
	userInfo, _ := user.(model.UserInfo)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err)
	}
	blogInfo, err := model.BlogDetail(id)
	if err != nil {
		fmt.Print(err)
	}
	c.HTML(http.StatusOK, "blog/detail.html", gin.H{
		"user":     userInfo.Username,
		"blogInfo": blogInfo,
	})
}
