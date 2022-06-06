package main

import "fmt"

func main() {
	mymap := make(map[string]string)
	//Map添加元素
	mymap["name"] = "Feng"

	//Map读取元素
	fmt.Println(mymap)
}
