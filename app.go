package main

var global_varible = "xx" //可以在整个 package 中使用

//全局变量的声明可以使用像上面 fmt 的那样组的形式进行
//但是局部变量不行
var (
	g1 = "sss"
	g2 = 99999999999
)

//常量也可以是以组的方式进行
const (
	con1 = "A"
	con2 = iota
	con3 = iota
	con4 //后面省略没有赋值 就是表示和之前上一个一样
	con5
)

func main() {

	// fa, fb, fc := myf.A("ss", 11, 22, 33)
	// fmt.Println(fa, fb, fc)
	// fmt.Println(myf.A0("xixi "))

	// fmt.Println("closure:::")
	// fclosure1 := myf.Closure1(1)
	// fmt.Println(fclosure1(8))
	// myf.Closure2()
	// myf.Closure3()
	// myf.Closure4()

	// //panic
	// myf.Pan()

	// //struct
	// mys.TestStruct()

	// myi.TestInterface()

	// //GoRoutine
	// //有缓存异步  c := make(chan bool, 1)

	// //无缓存  同步阻塞的
	// c := make(chan bool)
	// go func() {
	// 	fmt.Println("GoGOGO")
	// 	c <- true //存入消息
	// }()
	// <-c //取出  这里会等待  等待 channel 有东西的时候才继续执行（什么时候 c 里面有东东呢  就是上面 goroutine中向 c 写入之后）

	// c2 := make(chan bool)
	// go func() {
	// 	fmt.Println("GoGOGO2")
	// 	c2 <- true
	// 	close(c2)
	// }()
	// //对 channel 使用 range 的时候要保证 channel 显示的关闭
	// for v := range c2 {
	// 	fmt.Println(v)
	// }

	// //这里是用一个计数器 确保10个 goroutine 都执行后才结束
	// //更好的做法是使用 sync 中的 wait group
	// c3 := make(chan bool)
	// runtime.GOMAXPROCS(runtime.NumCPU())
	// timeCount := 0
	// for i := 0; i <= 9; i++ {
	// 	go mychannel.TestGo(c3, &timeCount, i)
	// }
	// <-c3

	// //下面是 waitGroup 的形式
	// //如果没有给 package 别名的话  package 的名字就是go 文件中第一行所指定的 package 的名字
	// mychannel.TestWaitGroup()

	// //带有缓存的 channel
	// mychannel.TestCacheChannel()

	// //calcPI
	// // c4 := make(chan float64)
	// // mychannel.CalcPI(c4)

	// mychannel.CalcPI2()
	// // mychannel.CalcPI3()

	// //mychannel.PopAndPush()
	// mytimetask.TimeEntry()
	// fmt.Println("mypackage.G", mypackage.G)
	// mypackage.InitGG()
	// fmt.Println("mypackage.GG ", mypackage.GG)

	// myjson.TestParse()

}
