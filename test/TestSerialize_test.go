package test

import (
	"testing"

	"code.byted.org/gopkg/pkg/log"
	"encoding/json"
	"fmt"
)
type IOuter interface {
	OutFn() string
}


type Inner struct {
	Sub string
}
type Outer struct {
	ID int
	Sub Inner
}
func (o Outer) OutFn() (string){
	return "str"
}


func Serial(data interface{}) (s []byte){
	var err error
	if s, err = json.Marshal(data); err != nil {
		log.Print(err)
	}
	return s
}
func DeSerial(s []byte) interface{}{
	outer := Outer{}
	if err := json.Unmarshal(s, &outer); err != nil {
		log.Print(err)
	}
	fmt.Printf("%+v \n", outer)
	log.Print(outer)
	return outer
}

func TestSerialize(t *testing.T) {
	outer := Outer{
		ID: 1,
		Sub: Inner{
			Sub:"jaja",
		},

	}
	rs := DeSerial(Serial(outer))
	if realV, ok := rs.(Outer); ok {
		fmt.Printf("in test %+v \n", realV)
	}else{
		fmt.Println("not ok")
	}

}