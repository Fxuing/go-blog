package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func request(url string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer response.Body.Close()
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println(string(res))

}
func TestClient(t *testing.T) {
	var (
		a float64 = 0.33333333
		b float64 = 0.66666661
		c         = a + b
	)
	fmt.Println(c)
	fmt.Printf("%f", c)
}
func bb(data []int) {
	fmt.Printf("%p\n", data)
	data[0] = 1000
	for i := 0; i < 10; i++ {
		data = append(data, i*10)
	}
	fmt.Println(data)
	fmt.Printf("len:%d, cap:%d \n", len(data), cap(data))
}

func TestOain(t *testing.T) {

	f1 := func() string {
		return func() string {
			return "hello "
		}()
	}

	a := 1
	switch a {
	case 1, 2, 3, 4:
	default:

	}

	t.Log(f1())
}
