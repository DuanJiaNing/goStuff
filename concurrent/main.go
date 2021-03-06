package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

var (
	grc  chan int
	grc1 chan map[string]interface{}
)

func init() {
	grc = make(chan int)
	grc1 = make(chan map[string]interface{})
}

func main6() {
	fmt.Println(1 << 10)
}

func main3() {

	//pc1()
	//pc()
	//pc2()
	//pc3()
	//pc4()

	//c := make(chan int)
	//c <- 3
	//<-c // should receive in different goroutine

	//pc5()
	pc6()

	c1 := make(chan int)
	close(c1)
	//c1 <- 1 // will panic
	// <- c1 // not block, but get zero value

	time.Sleep(time.Second * 4)
}

// 只读 chan: time Ticker Timer
func pc6() {
	ticker := time.NewTicker(time.Second)
	timer := time.NewTimer(time.Second * 2)

	count := 0
	for {
		select {
		case <-timer.C:
			p("time to quit")
			goto stop
		case <-ticker.C:
			count++
			p("ticker fired ", count, "th")
		}
	}

stop:
	p("select finished,count is ", count)
}

// 主 goroutine 等待子 goroutine 结束
func pc5() {
	exit := make(chan int)

	go func() {
		time.Sleep(time.Second * 2)

		exit <- 0
	}()

	<-exit
	p("finish")

}

// 不阻塞 chan，加上缓冲参数
func pc4() {
	c := make(chan int, 10) // 缓冲通道

	for i := 0; i < 10; i++ {
		c <- i
	}

	go func() {
		for v := range c {
			p(v)
		}
	}()

}

func pc3() {
	timer := time.NewTimer(time.Second)
	//for {
	p(<-timer.C)
	//}
}

// 单向 chan
func pc2() {

	c := make(chan<- int) // 只写：只能发送
	//<- c  illegal
	cc := make(<-chan int) // 只读：只能接收
	//cc <- 4 illegal
	go func() {
		<-cc
		c <- 3
	}()

}

// 一个 goroutine 中读，匹配另一个 goroutine 中的写
func pc1() {

	p("main goroutine")
	go func() {
		a := <-grc
		p("new goroutine", a, <-grc)

		b, ok := <-grc
		p("final data", b, ok)

		for d := range grc {
			p("iterate data", d)
		}

		p("finish")

	}()

	grc <- 2
	grc <- 3

	grc <- 4

	grc <- 1
	grc <- 2
	grc <- 3
	grc <- 4

}

func p(o ...interface{}) {
	fmt.Println(o...)
}

// 消费 chan 中数据，并作出回应
func printer(c chan int) {
	// 开始无限循环等待数据
	for {
		// 从channel中获取一个数据
		data := <-c
		// 将0视为数据结束
		if data == 0 {
			break
		}
		// 打印数据
		fmt.Println(data)
	}
	// 通知main已经结束循环(我搞定了!)
	c <- 0
}

// 发送数据到 chan
func pc() {
	// 创建一个channel
	c := make(chan int)
	// 并发执行printer, 传入channel
	go printer(c)
	for i := 1; i <= 10; i++ {
		// 将数据通过channel投送给printer
		c <- i
	}
	// 通知并发的printer结束循环(没数据啦!)
	c <- 0
	// 等待printer结束(搞定喊我!)
	<-c
}

// 模拟 rpc 调用方
func RPCClient(ch chan string, req string) (string, error) {
	ch <- req
	select {
	case ack := <-ch:
		return ack, nil
	case <-time.After(time.Second):
		return "", errors.New("Time out")
	}
}

// 模拟 rpc 调用服务提供方
func RPCServer(ch chan string) {
	for {
		data := <-ch
		fmt.Println("server received:", data)
		ch <- "received"
	}
}

func main1() {
	ch := make(chan string)

	go RPCServer(ch)
	recv, err := RPCClient(ch, "hi")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("client received", recv)
	}
}

type Obj struct {
	tm time.Time
}

func main() {
	s := `W/"4619f22a8bfeea0f105323899ff4e22d"`
	fmt.Println(s[3 : len(s)-1])
}

func main11() {
	sc := make(chan interface{}, 5)
	errSignal := make(chan error)
	for i := 0; i < 3; i++ {
		go func(index int) {
			if index == 1 {
				errSignal <- errors.New("err")
			}

			fmt.Println(time.Now(), " deliver ", index)
			sc <- strconv.Itoa(index)
			time.Sleep(time.Second * 2)
		}(i)
	}

	reciveCount := 0
fn:
	for {
		select {
		case err := <-errSignal:
			fmt.Println(time.Now(), err)
			break fn
		case l := <-sc:
			reciveCount++
			fmt.Println(time.Now(), " append ", l)
			if reciveCount == 3 {
				break fn
			}
		}
	}

	fmt.Println(time.Now(), " close")

}
