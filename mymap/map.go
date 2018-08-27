// +build OMIT
package main

import "fmt"

func main() {
	//map
	//需要注意  map slice 在函数中传递是引用 而数组 int string 是值传递 传的是一个拷贝
	m := make(map[int]string) //key 为 int 值类型为 string
	m[1] = "ss"
	m[2] = "bb"

	//如果不用 make 的方式声明 要给出字面量
	m1 := map[int]string{1: "heh", 2: "xixi"}
	fmt.Println(m1)

	_, ok := m1[3] //key不存在 ok 为 false
	fmt.Println("map key not exits ", ok)

	m2 := make([]map[int]string, 5) //声明了一个 slice  其元素类型是一个 map  map 类型为[int]string
	m2[0] = make(map[int]string)
	m2[0][0] = "haha"
	fmt.Println(m2)

}
