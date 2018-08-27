// +build OMIT
package main

import (
	"fmt"
	"reflect"
)

func main() {
	//定义数组
	var s [2]bool  //这里定义了一个长度为2的数组  或者认为 定一个的变量 s 其类型是一个 长度为2的 bool 数组
	var s2 [3]bool //注意不能 s = s2 因为他们是虽然都是 bool 数组 但是他们其实类型不同  s2是一个长度为3 的 bool 数组

	fmt.Println(s)
	var s1 = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} //使用这样三个点声明数组得是时候 后面要有字面值
	fmt.Println(s1, s2)
	//数组在函数传递时是 值类型  也就是传参的时候实际上是复制了一下这个数组  slice 是引用传递

	var s3 = [2][3]int{
		{1, 2, 3},
		{4, 5, 6}}
	fmt.Println(s3)
	//需要注意 初始化数组的值  用的{} 而不是[]
	//另外 }不能单独占一行  要和最后一个元素在一行

	//slice 非正式声明
	var sl10 = []int{1, 2, 3}
	mp := make(map[string]interface{})
	mp["isSlice"] = sl10
	ArrayOrSlice(mp)

	var sl1 = s1[5:9]
	fmt.Println("slice ：", sl10, sl1)
	for i, val := range sl1 {
		val = 999
		fmt.Print(i, val) //注意在 range 这里得到的 val 是一个拷贝  对 val 的修改不会影响到 sl1 的
		//要改sl1的话 用 sl1[i] = xxx
	}

	for i, v := range sl1 {
		fmt.Println("range ", i, v)
	}

	//slice 正式声明用 make
	var sl2 = make([]int, 3, 10) //参数2 是初始元素个数（声明时会初始化）  参数3 是容量 不写参数3 则容量值就是参数2的值
	fmt.Println("slice2:", sl2, len(sl2), cap(sl2))
	fmt.Printf("sl2地址为 %p\n", sl2)
	sl2 = changeSlice(sl2) //这样也是可以修改 sl2的
	fmt.Println("sl2", sl2)
	fmt.Printf("sl2地址为 %p\n", sl2)

	//append(sl1, sl2)
	//copy(sl1, sl2)
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

func changeSlice(sl []int) []int {
	sl = append(sl, 8)
	return sl
}
