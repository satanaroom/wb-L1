package main

import (
	"fmt"
	"time"
)

// Функция, которая ожидает время, переданное в аргументе, после чего отправляет
// текущее время в канал и тут же считывает его
func MySleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	unitsOfTime := 3
	fmt.Println("timer started with", unitsOfTime, "seconds")
	MySleep(time.Duration(unitsOfTime) * time.Second)
	fmt.Print("timer expired")
}
