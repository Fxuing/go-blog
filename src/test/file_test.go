package test

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestFileRead(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	average(100, 10)
}

func average(amount float64, nums int) {
	remainNum := float64(nums)
	var value float64
	for i := 1; i <= nums; i++ {
		if remainNum == 1 {
			value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", amount), 64)
		} else {
			value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", rand.Float64()*(amount/remainNum*2-0.01)), 64)
			//value = rand.Float64() * (amount/remainNum*2 - 0.01)
		}
		amount = amount - value
		remainNum--
		fmt.Printf("第%d个人抢到%.2f元红包，剩余红包%.2f元\n", i, value, amount)
	}
}

func TestBase64(t *testing.T) {
	var encode base64.Encoding
	res := encode.EncodeToString([]byte("teststring"))
	t.Log(res)

}
