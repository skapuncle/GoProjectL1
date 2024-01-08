package main

import (
	"fmt"
	"unicode"
)

// Функция IsUnique проверяет, являются ли все символы в строке уникальными.
func IsUnique(str string) bool {
	// Создаем map, который будет хранить уже встреченные символы.
	seen := make(map[rune]bool)
	// Проходим по всем символам в строке.
	for _, r := range str {
		// Преобразуем символ в нижний регистр, чтобы проверка была регистронезависимой.
		r = unicode.ToLower(r)
		// Если символ уже есть в map, возвращаем false, т.к. символы в строке не уникальны.
		if seen[r] {
			return false
		}
		// Если символа еще нет в map, добавляем его туда.
		seen[r] = true
	}
	// Если мы прошли по всем символам и не нашли повторяющихся, возвращаем true.
	return true
}

func main() {
	// Проверяем, являются ли все символы в строках уникальными.
	fmt.Println(IsUnique("abcd"))
	fmt.Println(IsUnique("abCdefAaf"))
	fmt.Println(IsUnique("aabcd"))
}
