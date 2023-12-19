package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// Создаем worker с двумя каналами остановки.
	worker := &Worker{stopCh: make(chan struct{})}
	fmt.Println("Goroutines run")
	worker.wg.Add(2)
	// Запускаем горутину stopByChannel.
	go worker.stopByChannel()

	// Создаем контекст с функцией отмены.
	worker.ctx, worker.cancel = context.WithCancel(context.Background())
	// Запускаем горутину stopByContext.
	go worker.stopByContext()

	// Ожидаем 3 секунды.
	time.Sleep(3 * time.Second)

	// Останавливаем горутины.
	fmt.Println("Goroutines stop")
	close(worker.stopCh)
	worker.cancel()

	// Ждем завершения всех горутин.
	worker.wg.Wait()

	fmt.Println("Shutting down...")
}

// Структура Worker содержит все необходимые каналы и контекст для остановки горутин.
type Worker struct {
	stopCh chan struct{} // Канал для передачи сигнала остановки.
	wg     sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
}

// Метод stopByChannel использует канал для передачи сигнала остановки горутине.
func (w *Worker) stopByChannel() {
	defer w.wg.Done()
	for {
		select {
		case <-w.stopCh:
			fmt.Println("Worker (stopByChannel) is stopped")
			return
		default:
			fmt.Println("Worker (stopByChannel) is running")
			time.Sleep(time.Second)
		}
	}
}

// Метод stopByContext использует контекст для передачи сигнала остановки горутине.
func (w *Worker) stopByContext() {
	defer w.wg.Done()
	for {
		select {
		case <-w.ctx.Done():
			fmt.Println("Горутина остановлена (stopByContext)")
			return
		default:
			// Выполнение работы горутины.
			fmt.Println("Выполняется работа горутины (stopByContext)")
			time.Sleep(time.Second)
		}
	}
}
