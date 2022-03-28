package main

/*Реализовать быструю сортировку массива (quicksort) встроенными методами языка.*/

import (
	"fmt"
	"sort"
)

type Arr struct {
	Num int
}

type ByArr []Arr

func (a ByArr) Len() int {
	return len(a)
}
func (a ByArr) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByArr) Less(i, j int) bool {
	return a[i].Num < a[j].Num
}

func main() {
	nums := []Arr{{5}, {1}, {3}, {22}, {1}, {13}, {54}, {2}, {90}, {71}}
	fmt.Println("arr before sort:", nums)
	// Функция Sort пакета sort выполняет быструю сортировку
	sort.Sort(ByArr(nums))
	fmt.Println("arr after  sort:", nums)
}
