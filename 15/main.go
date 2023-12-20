package main

import (
	"bytes"
	"fmt"
	"strings"
)

// Функция создает огромную строку заданного размера.
func createHugeString(size int) string {
	var builder strings.Builder
	builder.WriteString(strings.Repeat("a", size))
	return builder.String()
}

var justString string

// Функция создает огромную строку и присваивает ее первые 100 символов переменной justString.
func someFunc() {
	// Создаем огромную строку.
	v := createHugeString(1 << 10)
	// Создаем буфер из строки и присваиваем ему первые 100 символов переменной justString.
	buf := bytes.NewBufferString(v)
	justString = buf.String()[:100]
}

func main() {
	// Вызываем функцию someFunc.
	someFunc()
	// Выводим значение переменной justString.
	fmt.Println(justString)
}
