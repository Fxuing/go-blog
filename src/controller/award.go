package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/src/model"
	"math/rand"
	"net/http"
	"time"
)

var str chan string

func Lottery(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": lottery(),
	})
	go func() {
		ret := func() string {
			return "test"
		}
		str <- ret()
	}()

}

func sendMsg() {
	s := <-str
	time.Sleep(time.Second * 10)
	fmt.Println(s)
}

func lottery() string {
	var (
		awards         []*model.ActivityAward
		err            error
		probability    = 0
		probabilityArr [2]int
		msg            string
	)
	awards, err = model.SearchAward()
	if err != nil {
		fmt.Println(err)
	}
	// 概率数组
	probabilitys := make([]int, len(awards), cap(awards))
	for i, award := range awards {
		probabilitys[i] = award.Probability
	}
	// 奖品概率区间
	awardProbability := make(map[int][2]int)
	for i := 0; i < len(probabilitys); i++ {
		probabilityArr[0] = probability
		probabilityArr[1] = probability + probabilitys[i]
		probability = probabilityArr[1]
		awardProbability[i] = probabilityArr
	}
	// 随机数对应抽奖区间
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(100)
	for k, v := range awardProbability {
		if random >= v[0] && random <= v[1] {
			msg = awards[k].AwardDesc
			err = model.InventoryReduction(awards[k].Id)
			if err != nil {
				fmt.Println(err)
				msg = "未中奖，再接再厉"
			}
			break
		} else {
			msg = "未中奖，再接再厉"
		}
	}
	return msg
}
