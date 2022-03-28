package main

/*Реализовать все возможные способы остановки выполнения горутины.*/

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func receiveRandomValue() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10)
}

func main() {
	// Первый вариант:
	// Остановка сразу всех горутин через контекст
	ctx, cancel := context.WithCancel(context.Background())
	ch1 := make(chan int, 1)

	for i := 0; i < 5; i++ {
		go func(ctx context.Context, i int, ch1 chan<- int) {
			for {
				waitTime := time.Duration(rand.Intn(100)+10) * time.Millisecond
				fmt.Println(i, "sleep", waitTime)
				select {
				case <-ctx.Done():
					return
				case <-time.After(waitTime):
					fmt.Println("goroutine", i, "is finished")
					ch1 <- i
				}
				time.Sleep(500 * time.Millisecond)
			}
		}(ctx, i, ch1)
	}

	firstResult := <-ch1
	fmt.Println("first result goroutine:", firstResult)
	cancel()

	time.Sleep(1 * time.Second)
	fmt.Printf("\t\tFirst option finished\n")

	// Второй вариант:
	// Остановка горутины в случае, если необходимо остановить только одну
	ch2 := make(chan int, 1)
	stop := make(chan struct{})
	go func() {
		for {
			select {
			case ch2 <- receiveRandomValue():

			case <-stop:
				close(ch2)
				return
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		time.Sleep(1 * time.Second)
		stop <- struct{}{}
	}()

	for i := range ch2 {
		fmt.Println("received value: ", i)
	}

	fmt.Printf("\t\tSecond option finished\n")

	// Третий вариант:
	// Отслеживание закрытие канала
	ch3 := make(chan int, 1)
	wg1 := &sync.WaitGroup{}

	wg1.Add(1)
	go func() {
		defer wg1.Done()
		for i := 0; i < 5; i++ {
			ch3 <- i
		}
		close(ch3)
	}()

	for i := 0; i < 5; i++ {
		wg1.Add(1)
		go func(i int) {
			defer wg1.Done()
			for val := range ch3 {
				fmt.Println("value from channel:", val)
			}
			fmt.Println("goroutine", i, "finished")
		}(i)
	}
	wg1.Wait()
	fmt.Printf("\t\tThird option finished\n")
	time.Sleep(1 * time.Second)

	// Четвертый вариант:
	// Отановка по условию
	// (также можно останавливать по ошибке)
	ch4 := make(chan int, 1)
	wg2 := &sync.WaitGroup{}

	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for {
			n := receiveRandomValue()
			if n == 7 {
				fmt.Println("random value:", n, "finally 7")
				close(ch4)
				return
			}
			ch4 <- n
		}
	}()

	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for val := range ch4 {
			fmt.Println("random value:", val, "but not a 7")
		}
		fmt.Println("goroutine finished")
	}()

	wg2.Wait()
	fmt.Printf("\t\tFourth option finished\n")
}
