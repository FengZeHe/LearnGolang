package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	baseCtx := context.Background()
	ctx := context.WithValue(baseCtx, "name", "feng")
	ctxcancel, cancel := context.WithCancel(baseCtx)
	fmt.Println(ctx.Value("name"))

	go func() {
		for {
			select {
			case <-ctxcancel.Done():
				fmt.Println("done")
				return
			default:
				fmt.Println("run")
			}
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Stop")
	cancel()
	time.Sleep(1 * time.Second)
}
