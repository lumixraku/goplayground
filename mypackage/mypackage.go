package mypackage

import (
	"fmt"
	"goplayground/mystruct"
)

var (
	G  mystruct.SA
	GG *mystruct.SA
)

func init() {
	G = mystruct.SA{Prop1: "G"}
	fmt.Println("mypackge init")
}

func InitGG() {
	GG = &mystruct.SA{Prop1: "GG"}
}

func Send() {
	fmt.Print("Send")
}
