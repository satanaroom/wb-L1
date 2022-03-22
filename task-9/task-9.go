package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	nums := [8]int{2, 4, 6, 8, 10, 12, 14, 16}
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	// Запись в первый канал
	go func() {
		defer wg.Done()
		for i := range nums {
			fmt.Println("add to ch1:", nums[i])
			ch1 <- nums[i]
			runtime.Gosched()
		}
		close(ch1)
	}()

	wg.Add(1)
	// Чтение значения из первого канала и запись удвоенного во второй
	go func() {
		defer wg.Done()
		for val := range ch1 {
			res := val * 2
			fmt.Println("\tadd to ch2 from ch1:", res)
			ch2 <- res
		}
		close(ch2)
	}()

	// Чтение их второго канала и вывод в терминал
	go func() {
		for val := range ch2 {
			fmt.Println("\t\tvalue from conveyor:", val)
		}
	}()

	wg.Wait()
}
