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
