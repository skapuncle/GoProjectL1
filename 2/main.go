package main

import (
	"fmt"
	"sync"
)

func main() {
	// Массив чисел, квадраты которых мы хотим вычислить.
	numbers := []int{2, 4, 6, 8, 10}

	// WaitGroup используется для ожидания завершения всех горутин.
	var wg sync.WaitGroup

	// Для каждого числа в массиве создаем горутину.
	for _, number := range numbers {
		wg.Add(1) // Увеличиваем счетчик WaitGroup.
		go func(num int) {
			defer wg.Done()        // Уменьшаем счетчик WaitGroup после завершения горутины.
			fmt.Println(num * num) // Вычисляем квадрат числа и выводим его.
		}(number)
	}

	// Ждем завершения всех горутин.
	wg.Wait()
}

// Второй вариант решения.

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	// Массив чисел, квадраты которых мы хотим вычислить.
//	numbers := []int{2, 4, 6, 8, 10}
//
//	// Канал для передачи результатов из горутин.
//	results := make(chan int)
//
//	// Для каждого числа в массиве создаем горутину.
//	for _, number := range numbers {
//		go func(num int) {
//			// Вычисляем квадрат числа и отправляем результат в канал.
//			results <- num * num
//		}(number)
//	}
//
//	// Читаем результаты из канала и выводим их.
//	for i := 0; i < len(numbers); i++ {
//		fmt.Println(<-results)
//	}
//}
