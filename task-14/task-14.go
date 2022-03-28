package main

/*Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.*/

import (
	"fmt"
)

// Функция, которая в рантайме определяет тип переменной
func GetValueType(t interface{}) {
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t)
	case bool:
		fmt.Printf("%T\n", t)
	case int:
		fmt.Printf("%T\n", t)
	case string:
		fmt.Printf("%T\n", t)
	case chan int:
		fmt.Printf("%T\n", t)
	}
}

func main() {
	var t interface{}

	num := 10
	str := "hello"
	flag := true
	ch := make(chan int, 1)

	t = num
	GetValueType(t)
	t = str
	GetValueType(t)
	t = flag
	GetValueType(t)
	t = ch
	GetValueType(t)
}
