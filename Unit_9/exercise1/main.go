package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("Hi from func 1")
		wg.Done()

	}()

	go func() {
		fmt.Println("Hi from func 2")
		wg.Done()
	}()
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
}
