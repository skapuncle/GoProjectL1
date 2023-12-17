package main

import (
	"fmt"
	"time"
)

func main() {
	// Считываем значение N с консоли.
	N := 0
	fmt.Scanln(&N)

	// Создаем канал для обмена данными между горутинами.
	dataChannel := make(chan int)

	// Запускаем горутину writer.
	go writer(dataChannel)

	// Запускаем горутину reader.
	go reader(dataChannel)

	// Ожидаем N секунд перед выводом сообщения "Shutting down...".
	time.Sleep(time.Duration(N) * time.Second)

	fmt.Println("Shutting down...")
}

// Функция writer записывает числа от 0 до 99999 в канал dataChannel.
func writer(dc chan int) {
	for i := 0; i < 100000; i++ {
		dc <- i
	}
	// Закрываем канал после записи всех чисел.
	close(dc)
}

// Функция reader читает числа из канала и выводит их на экран.
func reader(dc chan int) {
	for v := range dc {
		fmt.Println("Value: ", v)
	}
}
