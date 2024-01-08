package main

import "fmt"

// Интерфейс для принтера A
type PrinterA interface {
	PrintDocument(document Document) error
}

// Документ
type Document struct {
	Content string
}

// Принтер A
type ConcretePrinterA struct{}

// Печать документа на принтере A
func (p *ConcretePrinterA) PrintDocument(document Document) error {
	fmt.Println("Printing document on printer A:", document.Content)
	return nil
}

// Интерфейс для принтера B
type PrinterB interface {
	Print(content string) error
}

// Принтер B
type ConcretePrinterB struct{}

// Печать на принтере B
func (p *ConcretePrinterB) Print(content string) error {
	fmt.Println("Printing content on printer B:", content)
	return nil
}

// Адаптер для принтера B
type PrinterBAdapter struct {
	printerB PrinterB
}

// Реализация метода PrintDocument для адаптера принтера B
func (p *PrinterBAdapter) PrintDocument(document Document) error {
	return p.printerB.Print(document.Content)
}

func main() {
	printerA := &ConcretePrinterA{}
	printerB := &ConcretePrinterB{}

	// Создаем адаптер для принтера B
	printerBAdapter := &PrinterBAdapter{printerB: printerB}

	// Теперь мы можем печатать документы на обоих принтерах, используя один и тот же интерфейс
	printerA.PrintDocument(Document{Content: "Hello, World!"})
	printerBAdapter.PrintDocument(Document{Content: "Hello, Universe!"})
}
