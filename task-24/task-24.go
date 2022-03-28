package main

/*Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.*/

import (
	"fmt"
	"math"
)

// Структура с инкапсулированными полями - координатами точки
type Point struct {
	x float64
	y float64
}

// Инкапсулированный метод-конструктор инициализации точки на плоскости
func (p *Point) newPosition(x, y float64) {
	p.x = x
	p.y = y
}

func main() {
	var curPoint, newPoint Point
	// Инициализация текущей точки
	curPoint.newPosition(1.0, 1.0)
	// Инициализация конечной точки
	newPoint.newPosition(3.0, 1.0)

	// Расчет шага по оси X
	stepX := newPoint.x - curPoint.x
	// Расчет шага по оси Y
	stepY := newPoint.y - curPoint.y

	// Расчет расстояния по теореме Пифагора
	dist := math.Sqrt(stepX*stepX + stepY*stepY)
	fmt.Println("distance:", dist)
}
