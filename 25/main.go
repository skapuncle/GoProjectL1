package main

import (
	"fmt"
	"time"
)

// Функция Sleep принимает продолжительность ожидания в виде time.Duration.
// Она ожидает указанное количество времени, используя цикл for для постоянного проверки текущего времени.
func Sleep(duration time.Duration) {
	// Получаем текущее время.
	currentTime := time.Now()
	// Вычисляем время окончания ожидания.
	endTime := currentTime.Add(duration)

	// Цикл продолжается, пока текущее время меньше времени окончания ожидания.
	for currentTime.Before(endTime) {
		// Обновляем текущее время.
		currentTime = time.Now()
	}
}

func main() {
	// Выводим сообщение "Start".
	fmt.Println("Start")
	// Вызываем функцию Sleep, передавая в нее продолжительность ожидания 2 секунды.
	Sleep(2 * time.Second)
	// Выводим сообщение "End".
	fmt.Println("End")
}
