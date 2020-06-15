package controller

import (
	"github.com/gin-gonic/gin"
	"go-blog/src/common"
	"go-blog/src/model"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func AddChannel(c *gin.Context) {
	channelInfo := model.ChannelInfo{
		ChannelName: "渠道名字",
		ChannelNo:   channelNo(),
		ParentId:    5,
	}
	err := channelInfo.ChannelAdd(channelInfo)
	if err != nil {
		panic(err)
	}
}

func ChannelList(c *gin.Context) {
	channelInfo, err := new(model.ChannelInfo).FindByParentId(0)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, common.ResponseData{
		Code: "0",
		Msg:  "OK",
		Data: channelInfo,
	})
}

func channelNo() string {
	A := 65
	arr := make([]int, 26, 32)
	var channelNo strings.Builder
	for i := 0; i < len(arr); i++ {
		arr[i] = A
		A++
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		channelNo.WriteString(string(rune(arr[rand.Intn(25)])))
	}
	return channelNo.String()
}
