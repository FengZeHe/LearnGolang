package main

import (
	"fmt"
)

func main(){
	// &是取指针，*是获取指针的值
	var str string = "hello "
	strpointer := &str
	getStr := *&str
	getStr2 := *strpointer
	getStr3 := str
	fmt.Println(str,strpointer,*strpointer,getStr,getStr2,getStr3)

	str = "world"
	fmt.Println(str,strpointer,*strpointer,getStr,getStr2,getStr3)

	person1 := Student{Name:"Feng"}
	changeParam(&person1,"Person1")
	fmt.Println(person1)

	cannotchangeParam(person1,"Person2")
	fmt.Println(person1)

}

type Student struct{
	Name string
}

func changeParam(para *Student,value string){
	para.Name = value
}

func cannotchangeParam(para Student,value string){
	para.Name = value
}