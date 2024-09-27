package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	contador := 0
	gs := 100

	wg.Add(gs)

	for i := 0; i <= gs; i++ {
		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			fmt.Println(contador)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("El valor final es: ", contador)
}
