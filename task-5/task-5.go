package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
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
	// Получение аргументом количество воркеров
	if len(os.Args) < 3 {
		fmt.Println("please, enter the number of workers and number of seconds to stop the programm: go run task-4 workersNum secondsNum")
		return
	}
	workerNum, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf(`"%s" is not a number`, os.Args[1])
		return
	}

	secondsNum, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf(`"%s" is not a number`, os.Args[2])
		return
	}
	// Созданние таймера для прерывания программы, по его завершению
	timer := time.NewTimer(time.Duration(secondsNum) * time.Second)

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

	// Select ожидает чтения из канала, когда таймер завершится.
	select {
	case <-timer.C:
		fmt.Println("work finished")
	}
}
