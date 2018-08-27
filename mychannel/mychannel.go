package mychannel

import (
	"fmt"
	"math"
	//"reflect"
	"runtime"
	"sync"
	"time"
)

func Go() {

	fmt.Println("GOGO")

}

//chan bool 整体一起作为类型
func TestGo(c chan bool, timeCount *int, index int) {
	a := 0
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println("c3 value", a, timeCount, index)
	*timeCount = *timeCount + 1
	if *timeCount == 9 {
		c <- true //向 channel 中写入 true
	}

}

//c1 是从 app 中传入的  所以是引用传递
func CalcPI(c1 chan float64) {
	pi := 0.0
	looptime := 10000

	//stop:= make(chan bool)

	for i := 0.0; i < float64(looptime); i++ {
		go func(i float64) {
			//iv := 4 * math.Pow(-1, i) / (2*i + 1)
			iv := math.Pow(-1, i) * (4 / float64(2*i+1))
			//fmt.Println("iv", iv)
			c1 <- iv

		}(float64(i))

	}

	for i := 0; i < looptime; i++ {
		pi += <-c1
	}

	fmt.Println("pi", pi)

}

//返回一个单向的 channel  它只能取
func boring(c1 chan float64, i float64, looptime int) <-chan float64 {
	//iv := 4 * math.Pow(-1, i) / (2*i + 1)
	iv := math.Pow(-1, i) * (4 / float64(2*i+1))
	//fmt.Println("iv", iv)
	c1 <- iv

	// 这样似乎不对 因为 goroutine 是争抢执行的  靠 index 来判断是不行的
	if i >= float64(looptime-1) {
		close(c1)
	}
	return c1

}

func CalcPI2() {
	c1 := make(chan float64)
	pi := 0.0
	looptime := 10000

	//stop:= make(chan bool)

	for i := 0.0; i < float64(looptime); i++ {
		go boring(c1, float64(i), looptime)
	}

	//for i:=0; i< looptime; i++{
	//	pi += <- c1
	//}

	//循环接收消息可以用 range  相比上面的方式更简单一些
	//使用 range 需要显式的关闭 channel
	for v := range c1 {
		pi = pi + float64(v)
	}
	fmt.Println("pi2 ", pi)
}

func CalcPI3() {
	c1 := make(chan float64)
	pi := 0.0
	looptime := 10

	//stop:= make(chan bool)
	count := 0

	go func() {
		defer close(c1)
		for i := 0.0; i < float64(looptime); i++ {
			//iv := 4 * math.Pow(-1, i) / (2*i + 1)
			iv := math.Pow(-1, i) * (4 / float64(2*i+1))

			c1 <- iv
			count++

		}
	}()

	// select 往往和 无限循环一起用 因为 select 也不知道什么时候 channel 有了值
	// 使用 select 的话  因为时放在无限循环中 所以要知道什么时候停止
	for {
		select {
		case v, ok := <-c1:
			fmt.Printf("in calc3")
			if !ok {
				fmt.Println("NOT OK")
				return
			}
			pi = pi + float64(v)
		}
	}
	fmt.Println("pi3 ", pi)
}

func TestWaitGroup() {
	//创建10个同时运行的 goroutine 在所有任务执行完成之后输出
	//不创建 channel
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(10)
	rs := 0
	for i := 0; i < 10; i++ {
		go func(i int) {
			for j := 0; j < 1000000; j++ {
				rs += j
			}
			fmt.Println("wg ", i, rs)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func TestCacheChannel() {
	c1 := make(chan bool, 2)

	//类似于队列的概念
	c1 <- true //存入一次最多也就读取一次  只有全部 channel 中的数据都被读取走了 <-c 才有阻塞的作用
	<-c1
	//<- c1
	fmt.Println("cache channel")
}

func PopAndPush() {
	c1 := make(chan int)
	c2 := make(chan int)
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1e9) // one second
		timeout <- true
	}()

	pump1(c1)
	pump2(c2)
	fmt.Println(".........")
	suck(c1, c2, timeout)
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suck(ch1 chan int, ch2 chan int, timeout chan bool) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Received on channel 1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("Received on channel 2: %d\n", v)
		case v := <-timeout:
			fmt.Printf("Received timeout ", v)
			break
		}
	}
}
