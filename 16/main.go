package main

import "fmt"

func main() {
	// Создаем слайс чисел, который нужно отсортировать.
	numbers := []int{2, 27, 14, 52, 31, 96, 73, 47, 22, 6, -9, 100, 50, -27, 0}
	fmt.Println("Unsorted slice: ", numbers)
	// Вызываем функцию quickSort для сортировки слайса.
	quickSort(numbers, 0, len(numbers)-1)
	fmt.Println("Sorted slice: ", numbers)
}

// Функция quickSort реализует алгоритм быстрой сортировки.
// Она принимает слайс чисел и индексы начала и конца слайса.
func quickSort(numbers []int, low int, high int) {
	// Если индекс начала слайса меньше индекса конца, продолжаем сортировку.
	if low < high {
		// Вызываем функцию partition для разделения слайса на две части.
		pivot := partition(numbers, low, high)
		// Рекурсивно сортируем две части слайса.
		quickSort(numbers, low, pivot-1)
		quickSort(numbers, pivot+1, high)
	}
}

// Функция partition разделяет слайс на две части относительно опорного элемента.
// Она принимает слайс чисел и индексы начала и конца слайса.
// Возвращает индекс опорного элемента после разделения.
func partition(numbers []int, low int, high int) int {
	// Выбираем опорный элемент.
	pivot := numbers[high]
	i := low - 1
	// Перебираем элементы слайса от начала до конца.
	for j := low; j < high; j++ {
		// Если текущий элемент меньше опорного, меняем их местами и увеличиваем счетчик i.
		if numbers[j] < pivot {
			i++
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}
	}
	// Меняем местами опорный элемент и элемент, следующий за последней парой (i+1).
	numbers[i+1], numbers[high] = numbers[high], numbers[i+1]
	// Возвращаем индекс опорного элемента после разделения.
	return i + 1
}
