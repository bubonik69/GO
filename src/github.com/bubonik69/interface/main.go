package main

import (
	"fmt"
	"math"
	"unicode/utf8"
)
//

//
// ***********function return function******************
// *****************************************************
// *****************interfaces
func sum(com string) func(a,b float32) float32 {
	if com == "+"{
		return func(a,b float32) float32{return a+b}
	} else if com == "-"{
		return func(a,b float32) float32{return a-b}
	}else if com == "*"{
		return func(a,b float32) float32{return a*b}
	}else{
		return func(a,b float32) float32{return a/b}
	}
}

type myStr string
type myStrInterface interface{
	isBig() bool
}

func (s myStr) isBig () bool{
	if utf8.RuneCountInString(string(s))>10{
		return true
	}
	return false
}
func main(){
	println(math.Pow(2,2))

	c:= myStr("gh")
	var i myStrInterface
	i=c
	fmt.Println(i.isBig())


	//d:=func(){return }
}