package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value int64
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

func (c *Counter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() {
	counter := NewCounter()
	var wg sync.WaitGroup

	// Запускаем 50 горутин, каждая из которых увеличивает счетчик 1000 раз
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	// Ждем завершения всех горутин
	wg.Wait()

	// Выводим итоговое значение счетчика
	fmt.Println(counter.Value())
}
