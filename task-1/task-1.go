package main

/*Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской
 структуры Human (аналог наследования).*/

import (
	"fmt"
	"math"
	"strings"
)

type Human struct {
	Name   string
	Age    int
	Height float64
	Weight float64
}

// Сеттер имени
func (h *Human) SetName(name string) {
	h.Name = name
}

// Устанавливает имя в верхний регистр
func (h *Human) ToUpperName() {
	h.Name = strings.ToUpper(h.Name)
}

// Считает body mass index
func (h Human) CalculateBMI() float64 {
	BMI := h.Weight / (math.Pow(h.Height, h.Height))
	return BMI
}

type Action struct {
	Human     // Встраивание структуры Human (теперь у Action есть доступ ко всем полям и методам Human)
	Surname   string
	Indicator string
}

// Устанавливает фамилию в верхний регистр
func (a *Action) ToUpperSurname() {
	a.Surname = strings.ToUpper(a.Surname)
}

// Выводит полное имя
func (a Action) SayFullName(name string) {
	// используем метод унаследованный от Human
	a.SetName(name)
	fmt.Printf("Full name is         : %s %s\n", a.Name, a.Surname)
	// используем метод унаследованный от Human
	a.ToUpperName()
	a.ToUpperSurname()
	fmt.Printf("To upper full name is: %s %s\n", a.Name, a.Surname)
}

// Выводит индекс массы тела
func (a Action) SayMBI() {
	fmt.Printf("MBI                is: %f %s", a.CalculateBMI(), a.Indicator)
}

func main() {
	a := Action{
		Surname:   "Petrov",
		Indicator: "kgf/m²",
		Human: Human{
			Name:   "",
			Age:    25,
			Height: 1.78,
			Weight: 75,
		},
	}

	a.SayFullName("Igor")
	a.SayMBI()
}
