package main

import (
	"fmt"
	"time"
	"strconv"
	myf "goplayground/myfunc"
	mys "goplayground/mystruct"
	myi "goplayground/myinterface"
	"goplayground/mychannel"

	"runtime"
	"reflect"
)

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


func main(){
	fmt.Println("sss", "bbb")

	fmt.Println("sd")
	fmt.Println("你什么")
	fmt.Println(time.Now().String())


	//局部变量可以这样简写声明  但是全局变量不行
	a2 := 1
	fmt.Println(a2)


	//类型转换
	a3 := 456732323231332
	fmt.Println(strconv.Itoa(a3))
	a33, _ := strconv.Atoi(strconv.Itoa(a3))
	fmt.Println(a33)

	//itoa
	fmt.Println(con1, con2, con3, con4, con5)


	// 条件语句是可以不用括号的  但是你想用也 OK
	if (a3 > 20 && 20 > 10) {
		fmt.Println("if haha")
	}



	//定义数组
	var s [2]bool  //这里定义了一个长度为2的数组  或者认为 定一个的变量 s 其类型是一个 长度为2的 bool 数组
	var s2 [3]bool  //注意不能 s = s2 因为他们是虽然都是 bool 数组 但是他们其实类型不同  s2是一个长度为3 的 bool 数组

	fmt.Println(s)
	var s1 = [...]int{0,1,2,3,4,5,6,7,8,9}  //使用这样三个点声明数组得是时候 后面要有字面值
	fmt.Println(s1, s2)
	//数组在函数传递时是 值类型  也就是传参的时候实际上是复制了一下这个数组  slice 是引用传递

	var s3 = [2][3]int {
		{1,2,3},
		{4,5,6}}
	fmt.Println(s3)
	//需要注意 初始化数组的值  用的{} 而不是[]
	//另外 }不能单独占一行  要和最后一个元素在一行


	//slice 非正式声明
	var sl10 =[] int{1,2,3}
	mp := make(map[string]interface{})
	mp["isSlice"] = sl10
	ArrayOrSlice(mp)

	var sl1 = s1[5:9]
	fmt.Println("slice ：", sl10, sl1)
	for i, val := range sl1{
		val = 999
		fmt.Print(i,val)  //注意在 range 这里得到的 val 是一个拷贝  对 val 的修改不会影响到 sl1 的
		//要改sl1的话 用 sl1[i] = xxx
	}


	for i,v := range sl1{
		fmt.Println("range ", i, v)
	}

	//slice 正式声明用 make
	var sl2 = make([]int,3,10)  //参数2 是初始元素个数（声明时会初始化）  参数3 是容量 不写参数3 则容量值就是参数2的值
	fmt.Println("slice2:", sl2, len(sl2), cap(sl2))
	fmt.Printf("sl2地址为 %p\n", sl2)
	sl2 = changeSlice(sl2)//这样也是可以修改 sl2的
	fmt.Println("sl2", sl2)
	fmt.Printf("sl2地址为 %p\n", sl2)

	//append(sl1, sl2)
	//copy(sl1, sl2)




	//map
	//需要注意  map slice 在函数中传递是引用 而数组 int string 是值传递 传的是一个拷贝
	m := make(map[int]string)  //key 为 int 值类型为 string
	m[1] = "ss"
	m[2] = "bb"

	//如果不用 make 的方式声明 要给出字面量
	m1 := map[int]string{1: "heh", 2: "xixi"}
	fmt.Println(m1)


	_, ok := m1[3]  //key不存在 ok 为 false
	fmt.Println("map key not exits ", ok)


	m2 := make([]map[int]string,5) //声明了一个 slice  其元素类型是一个 map  map 类型为[int]string
	m2[0] = make(map[int]string)
	m2[0][0] = "haha"
	fmt.Println(m2)

	fa, fb, fc := myf.A("ss", 11,22,33)
	fmt.Println(fa, fb, fc)
	fmt.Println(myf.A0("xixi "))

	fmt.Println("closure:::")
	fclosure1 := myf.Closure1(1)
	fmt.Println(fclosure1(8))
	myf.Closure2()
	myf.Closure3()
	myf.Closure4()

	//panic
	myf.Pan()

	//struct
	mys.TestStruct()

	myi.TestInterface()

	//GoRoutine
	//有缓存异步  c := make(chan bool, 1)

	//无缓存  同步阻塞的
	c := make(chan bool)
	go func(){
		fmt.Println("GoGOGO")
		c<- true //存入消息
	}()
	<-c //取出  这里会等待  等待 channel 有东西的时候才继续执行（什么时候 c 里面有东东呢  就是上面 goroutine中向 c 写入之后）


	c2 := make(chan bool)
	go func(){
		fmt.Println("GoGOGO2")
		c2<- true
		close(c2)
	}()
	//对 channel 使用 range 的时候要保证 channel 显示的关闭
	for v := range c2{
		fmt.Println(v)
	}


	//这里是用一个计数器 确保10个 goroutine 都执行后才结束
	//更好的做法是使用 sync 中的 wait group
	c3 := make(chan bool)
	runtime.GOMAXPROCS(runtime.NumCPU())
	timeCount := 0
	for i:=0; i<= 9; i++ {
		go mychannel.TestGo(c3, &timeCount, i)
	}
	<-c3

	//下面是 waitGroup 的形式
	//如果没有给 package 别名的话  package 的名字就是go 文件中第一行所指定的 package 的名字
	mychannel.TestWaitGroup()


	//带有缓存的 channel
	mychannel.TestCacheChannel()


	//calcPI
	c4 := make(chan float64)
	mychannel.CalcPI(c4)

	mychannel.CalcPI2()
	//myc.CalcPI3()

	//myc.PopAndPush()




}


func changeSlice(sl []int) ([]int){
	sl = append(sl, 8)
	return sl
}


func ArrayOrSlice(m map[string]interface{}) {
	for k, v := range m {
		rt := reflect.TypeOf(v)
		switch rt.Kind() {
		case reflect.Slice:
			fmt.Println(k, "is a slice with element type", rt.Elem())
		case reflect.Array:
			fmt.Println(k, "is an array with element type", rt.Elem())
		default:
			fmt.Println(k, "is something else entirely")
		}
	}
}