package myfunc

import "fmt"


//没有命名返回值
func A0(p1 string, pn ...int) (string){
	return p1 + " : hahah"
}

//命名返回值
func A(p1 string, pn ...int) (a,b,c int){
	//pn是 slice
	fmt.Println("func A ", pn)
	//注意这里不能写为 a, b, c := 1,2,3
	//因为命名返回值 相当于已经在函数中定义了变量 abc
	a, b, c = 1, 2, 3

	return
}

//闭包  返回值是一个函数 因此要说明函数的参数和函数的返回值
func Closure1(x int) (func(int) int){
	return func(y int) int{
		return x + y
	}
}

//defer 在函数出现重大 error 的时候也能执行  往往用于统计等
func Closure2(){
	fmt.Println("defer")

}

func Closure3(){
	fmt.Println("defer 1")
	for i :=0; i<3; i++  {
		 defer (func() {
			 fmt.Print(i)  //输出的都是3 因为这里 i 是一个引用
		 })()
	}
	fmt.Println("defer 2")
}

func Closure4(){

	for i :=0; i<3; i++  {
		defer (func(i int) {
			fmt.Println("closure4" , i)
		})(i) //值拷贝  所以这里可以顺利的输出 210
	}

}

//go中没有 try catch 使用 panic recover 的方式来处理错误
//panic 类似于抛出异常 会让程序停止执行  recover 只有在 defer 中才能调用
func Pan(){
	fmt.Println("\n")
	Pan1()
	Pan2()
	Pan3()
}
func Pan1(){
	fmt.Println("pan1")
}


func Pan2(){
	defer func(){
		//recover 这里相当于捕获了异常

		//下面这里可以简写 （if 语句可以包含一个初始化语句 他有固定的写法）
		//err := recover()
		//if err != nil {
		//	fmt.Println("pan2")
		//}



		if err := recover(); err != nil {
			fmt.Println("pan2")
		}
	}()

	panic("panic") //遇到 panic 后 panic 之后的语句将不再执行 因此 recover 要在 panic 之前

}


func Pan3(){
	fmt.Println("pan3")
}

