package main

import (
	"fmt"
	"reflect"
)

func main() {
	myarr := [5]int{1, 2, 3, 4, 5}
	myslice := []int{6, 7, 8, 9}
	fmt.Println(myarr, myslice, reflect.TypeOf(myarr), reflect.TypeOf(myslice))
	myslice1 := myslice[1:]
	myslice2 := myslice[:2]
	myslice3 := myslice[:]
	myslice4 := append(myslice1, myslice2...)
	fmt.Println(myslice1, myslice2, myslice3, myslice4)
	myslice4[0] = 99
	fmt.Println(myslice4)
}
