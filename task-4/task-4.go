package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

// Воркер, который считывает из канала число и выводит в терминал, добавляя свой номер
func Worker(in chan int, workerNum int) {
	for {
		fmt.Printf("random number %d from worker: %d\n", <-in, workerNum)
		time.Sleep(time.Second)
	}
}

func main() {
	// Создание канала, принимающего сигналы
	sigChan := make(chan os.Signal, 1)
	// Ретранслирование только сигнала interrupt
	signal.Notify(sigChan, syscall.SIGINT)

	// Получение аргументом количество воркеров
	if len(os.Args) == 1 {
		fmt.Println(`please, enter the number of workers: go run task-4 workersNum`)
		return
	}
	workerNum, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf(`"%s" is not a number`, os.Args[1])
		return
	}

	// Создание буферизированного канала для отправки в воркеров
	worketInput := make(chan int, 1)
	for i := 0; i < workerNum; i++ {
		go Worker(worketInput, i)
	}

	// Запуск горутины, в которой непрерывно подаются числа
	go func() {
		for {
			worketInput <- rand.Intn(workerNum * 2)
		}
	}()

	// В случае отправки сигнала interrupt (ctrl + C) программа завершается.
	// Сигнал попадет в копию процесса от имени которого запустилась программа.
	// При получении сигнала interrupt останавливаются все горутины.
	// Select ожидает чтения из канала, и в это время работают горутины.
	select {
	case <-sigChan:
		fmt.Println("work finished")
	}
}
