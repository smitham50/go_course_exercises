package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	count := 0
	gr := 100
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			v := count
			runtime.Gosched()
			v++
			count = v
			fmt.Println(count)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("End value:", count)
}
