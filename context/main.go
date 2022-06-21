package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	baseCtx := context.Background()
	ctx := context.WithValue(baseCtx, "name", "feng")
	fmt.Println(ctx.Value("name"))

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
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
