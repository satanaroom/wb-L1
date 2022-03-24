package main

import (
	"fmt"
)

// Функция, которая создает мапу с ключем-диапозоном и слайсом значений, находящихся в пределах ключа + переданный шаг
func GetSubsets(temp []float64, step int) map[int][]float64 {
	// Инициализация структуры, для хранения подмножеств
	subset := make(map[int][]float64)
	// Обход переданного слайса значений с вычислением диапозона и созданием подмножеств
	for i := range temp {
		stepKey := int(temp[i]/float64(step)) * step
		fmt.Println(stepKey, temp[i])
		subset[stepKey] = append(subset[stepKey], temp[i])
	}
	return subset
}

func main() {
	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	step := 10
	subsets := GetSubsets(temp, step)
	fmt.Println(subsets)
}
