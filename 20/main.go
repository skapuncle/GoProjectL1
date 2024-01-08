package main

import (
	"fmt"
	"strings"
)

// Функция reverseWordsInString принимает строку в качестве входных данных и возвращает строку,
// где порядок слов перевернут.
func reverseWordsInString(inputString string) string {
	// Функция strings.Fields разбивает строку на слова и возвращает слайс этих слов.
	words := strings.Fields(inputString)

	// Проходим по половине слайса слов, меняя местами соответствующие слова в начале и конце слайса.
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}

	// Функция strings.Join объединяет слова обратно в строку с пробелом между словами.
	return strings.Join(words, " ")
}

func main() {
	// Объявляем переменную для хранения входной строки.
	inputString := "snow dog sun"
	// Вызываем функцию reverseWordsInString, передавая в нее входную строку.
	res := reverseWordsInString(inputString)
	// Выводим полученную строку на экран.
	fmt.Println(res)
}
