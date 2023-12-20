package main

import (
	"errors"
	"fmt"
	"sync"
)

// SafeNumbers - структура для безопасной работы с map.
type SafeNumbers struct {
	sync.RWMutex
	numbers map[int]int
}

// Add - метод для добавления элемента в map.
func (sn *SafeNumbers) Add(num int) {
	sn.Lock()
	defer sn.Unlock()
	sn.numbers[num] = num
}

// Get - метод для получения элемента из map.
func (sn *SafeNumbers) Get(num int) (int, error) {
	sn.RLock()
	defer sn.RUnlock()
	if number, ok := sn.numbers[num]; ok {
		return number, nil
	}
	return 0, errors.New("Number does not exists")
}

// worker - функция, которая будет выполняться в каждой горутине.
func worker(id int, jobs <-chan int, safeNumbers *SafeNumbers, wg *sync.WaitGroup) {
	for j := range jobs {
		safeNumbers.Add(j)
		wg.Done()
	}
}

func main() {
	// Создаем экземпляр структуры SafeNumbers.
	safeNumbers := &SafeNumbers{
		numbers: make(map[int]int),
	}

	// Задаем общее количество задач (N) и количество воркеров (W).
	N := 1000
	W := 5

	// Создаем канал jobs для передачи задач воркерам.
	jobs := make(chan int, N)
	var wg sync.WaitGroup

	// Устанавливаем счетчик WaitGroup равным общему количеству задач.
	wg.Add(N)
	// Запускаем W воркеров.
	for w := 1; w <= W; w++ {
		go worker(w, jobs, safeNumbers, &wg)
	}

	// Заполняем канал jobs числами от 0 до N-1.
	for i := 0; i < N; i++ {
		jobs <- i
	}
	// Закрываем канал jobs, что означает, что воркеры больше не будут получать новые задачи.
	close(jobs)

	// Ожидаем завершения работы всех воркеров.
	wg.Wait()

	// Читаем числа из SafeNumbers и выводим их на экран.
	for i := 0; i < N; i++ {
		number, err := safeNumbers.Get(i)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(number)
		}
	}
}
