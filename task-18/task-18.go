package main

/*Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.*/

import (
	"fmt"
	"sync/atomic"
	"time"
)

// Инициализация структуры-счетчика
type Counter struct {
	Iterator int64
}

func main() {
	var counter Counter
	for i := 0; i < 100; i++ {
		go func() {
			// Для обеспечения конкурентной атомарной инкрементации
			// необходимо использовать пакет atomic.
			// Методы пакета обеспечивают гарантию записи всех значений
			// в конкурентной среде
			atomic.AddInt64(&counter.Iterator, 1)
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println("total iterations:", counter.Iterator)
}
