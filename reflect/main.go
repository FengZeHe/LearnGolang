package main

import (
	"reflect"
	"fmt"
)

func main(){
	str := "xxx"
	fmt.Println(reflect.TypeOf(str),reflect.ValueOf(str))

	myMap := make(map[string]string)
	myMap["name"]= "Feng"
	fmt.Println(reflect.TypeOf(myMap),reflect.TypeOf(myMap["name"]))

	person := Student{name:"xu",id:1}
	fmt.Println(person,reflect.TypeOf(person),reflect.ValueOf(person))
}

type Student struct{
	name string
	id int
}