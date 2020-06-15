package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/src/common"
	"go-blog/src/model"
	"net/http"
)

func Index(c *gin.Context) {

	user := common.GetSession(common.ContextUserKey, c)
	userInfo, _ := user.(model.UserInfo)

	blogList, err := model.BlogLoad()
	if err != nil {
		fmt.Println(err)
	}

	c.HTML(http.StatusOK, "index/index.html", gin.H{
		"user":     userInfo.Username,
		"blogList": blogList,
	})
}

func ClassById(c *gin.Context) {
	class, err := model.SearchClassById(1)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": class,
	})
}

func StudentById(c *gin.Context) {
	student, err := model.SearchStudentById(1)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": student,
	})
}
