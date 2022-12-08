package main

import (
	"fmt"
	"sync"
)

/*
	文档： https://cyent.github.io/golang/goroutine/sync_rwmutex/
	互斥锁 Mutex
	读写锁 RWMutex
	1. 和互斥锁不同的是，它可以分别针对读操作和写操作进行锁定和解锁操作
	2. 在读写锁管理的范围内，允许任意多个读操作同时进行，但相同时刻只允许一个写操作，并且写操作进行时同时不允许进行读操作。

*/

var mutex sync.Mutex
var rwmutex sync.RWMutex

func Mutex() {
	mutex.Lock()
	//如果加两次锁那么程序就会崩溃，不加锁直接解锁也会崩溃
	//mutex.Lock()
	defer mutex.Unlock()
	fmt.Print("wuwuwuwu")
}

/*
	func (rw *RWMutex) Lock       //写锁定
	func (rw *RWMutex) Unlock     //写解锁
	func (rw *RWMutex) RLock      //读锁定
	func (rw *RWMutex) RUnlock    //读解锁
*/

func RWMutex() {

	//加读锁
	//rwmutex.RLock()
	//defer rwmutex.RUnlock()

	//加写锁
	rwmutex.Lock()
	defer rwmutex.Unlock()
	fmt.Println("biubiubiu")

}

func main() {
	//Mutex()
	RWMutex()
}
