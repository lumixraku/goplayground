package main

import (
	"fmt"

	"code.byted.org/gopkg/pkg/log"
	"github.com/go-errors/errors"
)

// func main() {
// 	defer func() { //必须要先声明defer，否则不能捕获到panic异常
// 		if err := recover(); err != nil {
// 			fmt.Println("catch", err) //这里的err其实就是panic传入的内容，
// 		}
// 	}()
// 	a := []int{1, 2}
// 	fmt.Println(a[3]) // 越界访问，肯定出现异常
// }
func main() {
	defer func() { //必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("xiaorui.cc start")
		if err := recover(); err != nil {
			fmt.Println("....")
			log.Error(err.(*errors.Error).ErrorStack())
			// fmt.Println(err.(*errors.Error).ErrorStack()) //之后的内容不会执行
			fmt.Println("aaa")

			// fmt.Println(err) //这里的err其实就是panic传入的内容，"bug"
		}
		fmt.Println("xiaorui.cc end................................................")
	}()
	fmt.Println("Hello, playground")

	a := []int{1, 2}
	fmt.Println(a[3]) // 越界访问，肯定出现异常
}
