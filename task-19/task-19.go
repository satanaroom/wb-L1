package main

/*Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.*/

import (
	"fmt"
	"os"
	"strings"
)

// Функция, которая переворачивает буквы в поданном слайсе строк
func ReverseLetters(str []string) string {
	var afterRev []string
	var reversedStr string
	for i := range str {
		// Создание слайса рун, для работы с каждым символом по отдельности
		var res []rune
		// Раскладывание каждого слова на слайс рун
		r := []rune(str[i])
		// Идем в цикле с конца слова, где добавляем в слайс рун символы
		for j := len(r) - 1; j >= 0; j-- {
			res = append(res, r[j])
		}
		// Добавление в итоговый слайс перевернутых слов
		afterRev = append(afterRev, string(res))
	}
	reversedStr = strings.Join(afterRev, " ")
	return reversedStr
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("please, enter words to reverse like this: go run task-20.go word1 word2 wordN")
		return
	}
	str := os.Args[1:]
	reversed := ReverseLetters(str)
	fmt.Println(reversed)
}
