package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMut(t *testing.T) {
	var mut sync.Mutex
	var counter int
	for i := 0; i < 5000; i++ {
		go func(i int) {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}(i)
	}
	time.Sleep(time.Microsecond * 500)
	t.Log(counter)
}

func TestG(t *testing.T) {
	var waitGroup sync.WaitGroup
	var mut sync.Mutex
	var counter int
	for i := 0; i < 5000; i++ {
		waitGroup.Add(1)
		go func(i int) {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			waitGroup.Done()
		}(i)
	}
	waitGroup.Wait()
	t.Log(counter)
}

func service() string {
	//time.Sleep(time.Millisecond * 50)
	return "执行业务操作中"
}

func AsyncService() chan string {
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		retCh <- ret
	}()
	return retCh
}
func ohterTask() {
	fmt.Println("执行其他任务开始。")
	//time.Sleep(time.Millisecond * 100)
	fmt.Println("执行其他任务结束。")
}
func TestAsynService(t *testing.T) {
	retCh := AsyncService()
	ohterTask()
	fmt.Println(<-retCh)
}

func timeout() string {
	time.Sleep(time.Second * 3)
	return "userInfo======="
}
func userInfo() chan string {
	userCh := make(chan string)
	go func() {
		ret := timeout()
		userCh <- ret
	}()
	return userCh
}
func addressInfo() chan string {
	addCh := make(chan string)
	go func() {
		ret := func() string {
			return "address======="
		}
		addCh <- ret()
	}()
	return addCh
}
func TestSel(t *testing.T) {
	select {
	case ret := <-userInfo():
		t.Log(ret)
	//case ret := <- addressInfo():
	//	t.Log(ret)
	case <-time.After(time.Second):
		t.Error("time out")
	}
}

func provider(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 200; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}
func receiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Print(data, " ")
			} else {
				break
			}
		}
		wg.Done()
	}()
}
func TestChClose(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	provider(ch, &wg)
	wg.Add(1)
	receiver(ch, &wg)
	wg.Add(1)
	receiver(ch, &wg)
	wg.Wait()
}

type Singleton struct{}

var singleton *Singleton
var once sync.Once

func getInstance() *Singleton {
	once.Do(func() {
		fmt.Println("create singleton")
		singleton = new(Singleton)
	})
	return singleton
}
func TestRunOne(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			singleton := getInstance()
			fmt.Printf("%x %d", singleton, i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func cancel(ch chan string) {
	close(ch)
}
func isCancel(ch chan string) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}
func TestCancel(t *testing.T) {
	ch := make(chan string, 0)
	for i := 0; i < 5; i++ {
		go func(i int, ch chan string) {
			for {
				if isCancel(ch) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Done")
		}(i, ch)
	}
	cancel(ch)
	time.Sleep(time.Second * 2)
}

func TestContext(t *testing.T) {
}
