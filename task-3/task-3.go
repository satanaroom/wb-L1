package main

import (
	"fmt"
	"math"
	"sync"
)

// функция подсчета степени числа и накапливания суммы
func DoSumOfPow(in int, res *int, wg *sync.WaitGroup) {
	// Отложенный вызов метода, который уменьшает счетчик горутин
	defer wg.Done()
	*res += int(math.Pow(float64(in), float64(in)))
}

func main() {
	var nums [5]int
	var result int

	// Заполнение массива числами 2, 4, 6, 8, 10
	for i := range nums {
		nums[i] = (i + 1) * 2
	}
	// Добавление вэйт-группы для ожидания выполнения работы функции
	// Ссылка позволяет не копировать структуру WaitGroup после создания
	wg := &sync.WaitGroup{}
	for i := range nums {
		// Добавление горутины в WaitGroup при её запуске,
		// чтобы избежать случая, когда горутина может запуститься
		// до того, как завершится цикл по nums
		wg.Add(1)
		go DoSumOfPow(nums[i], &result, wg)
	}
	// Метод, ожидающий пока счетчик горутин не станет нулевым
	wg.Wait()
	// После того, как подождали, пока все горутины просуммируются, можно выводить результат
	fmt.Println(result)
}
