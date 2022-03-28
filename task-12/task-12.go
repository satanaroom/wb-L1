package main

/*Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.*/

import "fmt"

// Функция удаления дубликатов из множества
func RemoveDuplicates(str []string) []string {
	// Объявление результируещего множества
	var unique []string
	// Создание мапы для определения уникальности значений в результирующем множестве
	attend := make(map[string]bool)
	for _, val := range str {
		// Если в мапе нет элемента, добавляем
		if !attend[val] {
			unique = append(unique, val)
			// Устанавливаем ключ уникальности
			attend[val] = true
		}
	}
	return unique
}

func main() {
	notSequence := []string{"cat", "cat", "dog", "cat", "tree"}
	sequence := RemoveDuplicates(notSequence)
	fmt.Println("sequence:", sequence)
}
