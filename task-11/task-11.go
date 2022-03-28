package main

/*Реализовать пересечение двух неупорядоченных множеств.*/

import "fmt"

// Функция, которая создает пересечение неупорядоченных множеств
func Intersection(str1, str2 []string) []string {
	// Объявление результируещего множества-пересечения
	var inter []string
	// Создание мапы для определения уникальности значений в первом множестве
	hash := make(map[string]bool)
	// Пробегаем по каждому элементу первого множества и устанавливаем ключ уникальности в мапе
	for _, val := range str1 {
		hash[val] = true
	}
	// Пробегаем по каждому элементу второго множества, проверяя, есть ли в мапе
	// элемент из первого множества. Если есть, добавляем в результирующее множество
	for _, val := range str2 {
		if hash[val] {
			inter = append(inter, val)
		}
	}
	// Удаление дубликатов
	inter = RemoveDuplicates(inter)
	return inter
}

// Функция удаления дубликатов из пересеченного множества
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
	lots1 := []string{"one", "two", "three", "four"}
	lots2 := []string{"one", "two", "five", "six", "one"}
	inter := Intersection(lots1, lots2)
	fmt.Println(inter)
}
