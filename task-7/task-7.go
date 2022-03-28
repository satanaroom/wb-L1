package main

/*Реализовать конкурентную запись данных в map.*/

import (
	"fmt"
	"sync"
)

// Объявление структуры с мьютексом и мапой, в которую будет производится
// запись конкурентно
type Counter struct {
	sync.Mutex
	m map[string]int
}

func main() {
	c := Counter{
		m: map[string]int{
			"debt": 1000,
			"acc":  5000,
		},
	}

	// Запускаем 7 горутин, которые должны уменьшить долг и сбережения
	for i := 0; i < 7; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				// Для того, чтобы исключить состояние гонки, необходимо залочить данные на время их изменения
				c.Lock()
				c.m["debt"]--
				c.m["acc"]--
				// И разлочить после
				c.Unlock()
			}
		}()
	}

	fmt.Scanln()
	money := c.m["acc"]
	fmt.Println("how much money do i have", money)
	fmt.Println("they gave me 500 cents")

	// Запускаем 5 горутин, которые должны увеличить сбережения
	for i := 0; i < 5; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				// Для того, чтобы исключить состояние гонки, необходимо залочить данные на время их изменения
				c.Lock()
				c.m["acc"]++
				// И разлочить после
				c.Unlock()
			}
		}()
	}

	fmt.Scanln()
	money = c.m["acc"]
	fmt.Println("how much money do i have now", money)
}
