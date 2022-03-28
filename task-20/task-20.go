package main

/*Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».*/

import (
	"fmt"
	"os"
	"strings"
)

// Функция, которая переворачивает слова в поданном слайсе строк
func ReverseWords(str []string) string {
	var afterRev []string
	var reversed string

	// Идем в цикле с последнего слова, который добавляем в новый слайс
	for j := len(str) - 1; j >= 0; j-- {
		afterRev = append(afterRev, str[j])
	}
	reversed = strings.Join(afterRev, " ")
	return reversed
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("please, enter words to reverse like this: go run task-20.go word1 word2 wordN")
		return
	}
	str := os.Args[1:]
	reversed := ReverseWords(str)
	fmt.Println(reversed)
}
