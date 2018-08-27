// +build OMIT
package main

import (
	"fmt"
	"strconv"
)

func main() {

	//类型转换
	a3 := 456732323231332
	fmt.Println(strconv.Itoa(a3))
	a33, _ := strconv.Atoi(strconv.Itoa(a3))
	fmt.Println(a33)
}
