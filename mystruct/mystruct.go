package mystruct

import (
	//"encoding/json"
	//"strconv"
	"fmt"
	"time"
)

type person struct {
	Name string
	Age  int
}

type A1 struct {
	Form_id    string
	Data       string
	CreateTime time.Time
}

type struct1 struct {
	i1  int
	f1  float32
	str string
}

type SA struct {
	Prop1 string
}
type SB struct {
	SA
	Prop2 string
}

func TestStruct() {
	//字面量形式的初始化  //在 GO 中字面量初始化很常见
	p1 := person{
		Name: "persion1",
		Age:  19,
	}
	fmt.Println(p1)
	changePerson1(p1)
	fmt.Println(p1) //还是19 可见 struct 直接这样传递还是值传递
	changePerson2(&p1)
	fmt.Println(p1)

	//怎么样让p1本身就是一个指针呢  免得每次都取一下地址//这是推荐做法哟=====================
	p2 := &person{
		Name: "persion1",
		Age:  19,
	}
	fmt.Println(p2)

	//组合形式的结构
	sb := SB{
		SA:    SA{Prop1: "111"},
		Prop2: "22",
	}
	fmt.Println(sb)

	//方法
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Chris"

	fmt.Printf("The int is: %d\n", ms.i1)
	fmt.Printf("The float is: %f\n", ms.f1)
	fmt.Printf("The string is: %s\n", ms.str)
	fmt.Println(ms)
}

func changePerson1(person person) { //形参实际上是实参的一个拷贝
	person.Age = 1
	fmt.Println(person.Age)
}

func changePerson2(person *person) {
	person.Age = 12
}

//遵循函数的传值方式 传的 struct 是值 因此f2(a1 A1）得到的是拷贝  因此还是应该用传引用的方式 加上*
func f1(a1 *A1) {
	a1.Form_id = "1"
	t := time.Now()
	fmt.Println(t.Format("20060102150405"))
	a1.CreateTime = t
	fmt.Print(a1.Form_id)

	f2(a1)
}

func f2(a1 *A1) {
	a2 := a1
	a2.Form_id = "2"
	fmt.Print(a1.Form_id)
	fmt.Print(a2.Form_id)
}
