package main

import "fmt"

// Функция reverseString принимает строку в качестве входных данных и возвращает ее обратную копию.
func reverseString(inputString string) string {
	// Преобразуем строку в слайс рун (символов Unicode).
	runes := []rune(inputString)
	// Проходим по половине слайса рун, меняя местами соответствующие символы в начале и конце слайса.
	for i := 0; i < len(runes)/2; i++ {
		// Меняем местами i-й символ и символ с индексом, равным длине слайса минус i.
		runes[i], runes[len(runes)-1-i] = runes[len(runes)-1-i], runes[i]
	}
	// Преобразуем обратно в строку и возвращаем.
	return string(runes)
}

func main() {
	// Объявляем переменную для хранения входной строки.
	var inputString string
	// Читаем входную строку из стандартного ввода.
	fmt.Scanln(&inputString)
	// Вызываем функцию reverseString, передавая в нее входную строку.
	res := reverseString(inputString)
	// Выводим полученную строку на экран.
	fmt.Println(res)
}
