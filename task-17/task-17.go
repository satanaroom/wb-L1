package main

import (
	"fmt"
	"sort"
)

func main() {
	x := 3
	nums := [10]int{1, 2, 3, 4, 6, 10, 20, 30, 40, 100}
	// Функция Search пакета sort осуществляет бинарный поиск в отсортированной структуре данных
	i := sort.Search(len(nums), func(i int) bool { return nums[i] >= x })
	if i < len(nums) && nums[i] == x {
		fmt.Printf("%d is present at nums[%d]", x, i)
	} else {
		fmt.Printf("%d is not present in nums, but %d - index where it would be inserted", x, i)
	}
}
