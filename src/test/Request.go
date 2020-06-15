package test

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Request struct {
	Body   map[string]interface{} `json:"body"`
	Header map[string]interface{} `json:"header"`
}

func (r *Request) getInstance(param map[string]interface{}) *Request {
	request := Request{
		Body: param,
		Header: map[string]interface{}{
			"appId":       "pc2f6b4e4645304d577841303d",
			"format":      "json",
			"timeStamp":   strconv.Itoa(int(time.Now().UnixNano() / 1e6)),
			"partnerId":   "07",
			"signMethod":  "MD5",
			"deviceType":  "APP",
			"channelCode": "624e6e48304a4650435a773d",
			"diviceId":    "8a643b464769af800a075b29e1ad91f8",
		},
	}
	jsonStr, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("加密前:%s\n", string(jsonStr))

	md5 := md5.New()
	md5.Write([]byte("zhfY8Orq1ntPfmEYABrQDAxCUU9RPj"))
	md5.Write(jsonStr)
	sign := hex.EncodeToString(md5.Sum([]byte(nil)))
	request.Header["sign"] = strings.ToUpper(sign)
	jsonStr2, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("加密后:%s\n", string(jsonStr2))
	return &request
}

func (request *Request) Post(url string) string {
	jsonStr2, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
	}
	buf := &bytes.Buffer{}
	err = binary.Write(buf, binary.BigEndian, []byte(string(jsonStr2)))
	if err != nil {
		panic(err)
	}

	res, err := http.Post(url, "application/json", buf)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	resb, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return string(resb)
}
