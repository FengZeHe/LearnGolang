package main

import (
	"fmt"
	"sync"
)

/*
	https://www.cnblogs.com/qcrao-2018/p/12736031.html
*/

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			return &user{}
		},
	}

	//Get返回的是interface ，因此需要类型断言
	u := pool.Get().(*user)
	fmt.Println("init -->", u.Name, u.Email)
	defer pool.Put(u)

	u.Reset("Tom", "holyshit@qq.com")
	fmt.Print(u.Name, u.Email)

}

type user struct {
	Name  string
	Email string
}

// 重置里面的字段
func (u *user) Reset(name string, email string) {
	u.Email = email
	u.Name = name
}
