package main

/*Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.*/

import "fmt"

func main() {
	var num, res, bit, swap, flag int64

	fmt.Println("please, enter random number")
	fmt.Scanln(&num)
	fmt.Println("what number of bit you want to choose in", num)
	fmt.Scanln(&bit)
	fmt.Println("you want set 1 or 0?")
	fmt.Scanln(&flag)

	fmt.Printf("number before change: %b (%d in decimal number system)\n", num, num)
	// Создание бита для свапа, например если необходимо поменять 5-й бит,
	// то в двоичной системе это будет выглядеть как 10000
	swap = 1 << (bit - 1)
	if flag == 1 {
		// Если необходимо поставить на 5-й бит единицу, тогда она встанет
		// при помощи логического ИЛИ, если не стояла до этого
		res = swap | num
	} else if flag == 0 {
		// Иначе если выбрали поставить 0, он встанет на 5-й бит при помощи
		// сброса бита.
		// В результирующем числе будет работать правило:
		// каждый бит RES равен 0, если соответствующий бит SWAP равен 1.
		// если бит в SWAP равен 0, то берется значение соответствующего бита из NUM
		res = num &^ swap
	} else {
		fmt.Printf("%d is not a 1 or 0 number\n", flag)
		return
	}
	fmt.Printf("number after  change: %b (%d in decimal number system)\n", res, res)
}
