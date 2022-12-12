package main

import (
	"fmt"
	"sync/atomic"
)

/*
Go中的原子性：一个或多个操作在CPU的执行过程中不被中断的特性被称为原子性。这些操作对外表现成一个不可分割的整体，
他们要么都执行，要么都不执行，外界不会看到它们执行到一半的状态。

原子操作：进行过程中不能被中断的操作，原子操作由底层硬件支持，而锁则是由操作系统提供API实现，若实现相同的功能前者效率会更高
*/
var value int32 = 0

func main() {
	atomic.AddInt32(&value, 10)
	nv := atomic.LoadInt32(&value)
	fmt.Println(nv, value)
	//	输出是10

	swapped := atomic.CompareAndSwapInt32(&value, 10, 20)
	fmt.Println(swapped, value)

	swapped = atomic.CompareAndSwapInt32(&value, 20, 50)
	fmt.Println(swapped, value)

	swapped = atomic.CompareAndSwapInt32(&value, 19, 60)
	fmt.Println(swapped, value)
	//	因为旧值填的是19和之前的值不一样，所以交换值失败

	old := atomic.SwapInt32(&value, 40)
	fmt.Println(old, value)
}
