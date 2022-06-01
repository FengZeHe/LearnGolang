package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// var a, b, c, d int = 1, 2, 3, 4
	var a, b, c, d = "six", 6, true, 6.6
	fmt.Println(a, b, c, d)

	// 类型转换 int-> float64 -> uint
	var i int = 12
	var f64 float64 = float64(i)
	var u uint = uint(f64)
	fmt.Println(u)

	// string转int/int64
	numint, _ := strconv.Atoi("99")
	numint64, _ := strconv.ParseInt("99", 10, 64)
	fmt.Println(numint, numint64, reflect.TypeOf(numint), reflect.TypeOf(numint64))

	//int/int64转string
	var num1 int = 99
	var num2 int64 = 99
	intstr := strconv.Itoa(num1)
	int64str := strconv.FormatInt(num2, 10)
	fmt.Println(intstr, int64str, reflect.TypeOf(intstr), reflect.TypeOf(int64str))
}
