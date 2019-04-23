package main

import (
	"fmt"
	"sync"
	"time"
)

func Make_goroutine(i int, wg *sync.WaitGroup) {
	fmt.Printf("Goroutine %d commence\n", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d finit\n", i)
	wg.Done()
}

func main() {

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go Make_goroutine(i, &wg)
	}
	wg.Wait()
	fmt.Println("Tout les Goroutines se sont exécutés")
}
