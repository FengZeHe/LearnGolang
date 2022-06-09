package main

import "fmt"

func main() {
	mymap := make(map[string]string)
	//Map添加元素
	mymap["name"] = "Feng"

	//Map读取元素
	fmt.Println(mymap["name"])

	// 遍历Map
	for k, v := range mymap {
		fmt.Println(k, v)
	}

	// 删除Map元素
	delete(mymap, "name")
	fmt.Println("delete:", mymap["name"])

	// 判断Map中是否存在某个元素
	_, exist := mymap["name"]
	if exist == false {
		fmt.Println("元素不存在")
	} else {
		fmt.Println("元素存在")
	}

}
