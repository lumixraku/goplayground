package myinterface

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u *User) Hello() {
	fmt.Println("Hello")
}

//这里表示 Info 接受的参数类型是空接口  也就是任何类型都可以
func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("type ", t)

	v := reflect.ValueOf(o)
	fmt.Println("value is ", v)


	//获取各个字段的信息
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("info fName %v, fType %s, val %v\n", f.Name, f.Type, val)
	}

	//获取 method 信息
	for i := 0; i < t.NumMethod(); i++ {
		m:=t.Method(i)
		fmt.Printf("info fmethod %v %v\n", m.Name, m.Type)
	}
}

func TestInterface() {
	u := User{
		Id:   1,
		Name: "hehe",
		Age:  11,
	}
	//如果你是按照字段顺序赋值的  其实还可以这样
	u = User{1, "hehe", 11}

	u.Hello()
	Info(u)
}
