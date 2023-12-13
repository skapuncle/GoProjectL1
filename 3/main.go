package main

import (
	"fmt"
)

func main() {
	// Массив чисел, квадраты которых мы хотим вычислить.
	numbers := []int{2, 4, 6, 8, 10}

	// Канал для передачи результатов из горутин.
	results := make(chan int)

	// Для каждого числа в массиве создаем горутину.
	for _, number := range numbers {
		go func(num int) {
			// Вычисляем квадрат числа и отправляем результат в канал.
			results <- num * num
		}(number)
	}

	// Суммируем результаты из канала.
	var sum int
	for i := 0; i < len(numbers); i++ {
		sum += <-results
	}

	// Выводим сумму.
	fmt.Println(sum)
}
