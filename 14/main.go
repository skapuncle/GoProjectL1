package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Создаем переменную типа interface{}
	var i interface{}

	// Присваиваем ей значение типа int
	i = 10
	fmt.Println(reflect.TypeOf(i)) // Выводит: int

	// Присваиваем ей значение типа string
	i = "hello"
	fmt.Println(reflect.TypeOf(i)) // Выводит: string

	// Присваиваем ей значение типа bool
	i = true
	fmt.Println(reflect.TypeOf(i)) // Выводит: bool

	// Создаем канал
	ch := make(chan int)
	// Присваиваем ей значение канала
	i = ch
	fmt.Println(reflect.TypeOf(i)) // Выводит: chan int
}
