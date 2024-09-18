package main

import (
	"fmt"
	"runtime"
	"sync"
)

var ws sync.WaitGroup

func main() {
	ws.Add(2)

	go foo()
	go bar()

	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	ws.Wait()
}

func foo() {
	for i := 0; i <= 10; i++ {
		fmt.Println("Ejecutandose foo #", i)
	}
	ws.Done()
}

func bar() {
	for i := 0; i <= 10; i++ {
		fmt.Println("Ejecutandose bar #", i)
	}
	ws.Done()
}
