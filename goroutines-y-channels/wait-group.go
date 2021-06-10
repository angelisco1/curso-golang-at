package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		time.Sleep(time.Second * 3)
		wg.Done()
		fmt.Println("Fin1")
	}()

	go func() {
		time.Sleep(time.Second * 6)
		wg.Done()
		fmt.Println("Fin2")
	}()

	wg.Wait()

	fmt.Println("Fin3")

}
