package main

import (
	"fmt"
	"strings"
)

func IsIsogram(word string) bool {
	// Создание пустой строки для последующей конкатенации символов
	cache := ""
	// Создание регистронезависимости
	word = strings.ToLower(word)
	for _, val := range word {
		// Пропускаем тире, можно еще пробелы, для функции, работающей с предложениями
		if string(val) == "-" {
			continue
			// Проверка содержания в строке cache символа из слова
		} else if strings.Contains(cache, string(val)) {
			return false
			// Если проверка не удалась - добавляем в cache символ
		} else {
			cache += string(val)
		}
	}

	return true
}

func main() {
	testWords := []string{"abcd", "abCdefAaf", "aabcd", "raw-daw", "raw-pol"}
	for _, word := range testWords {
		if IsIsogram(word) {
			fmt.Println("word", word, "is isogram")
		} else {
			fmt.Println("word", word, "is not isogram")
		}
	}
}
