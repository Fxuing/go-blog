package test

import (
	"fmt"
	"testing"
)

func TestRequest(t *testing.T) {
	/*param := map[string]interface{}{
		"username":"yi_t@9hjf.com",
		"password":"yt123456",
		"userType":"ODP",
	}
	str := new(Request).getInstance(param).Post("https://www.o-banks.cn/odpapi/login")
	t.Log(str)*/
	m := map[string]interface{}{
		"username":  "yi_t@9hjf.com",
		"password":  "yt123456",
		"userType":  "ODP",
		"auserType": "ODP",
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println(m)
}
