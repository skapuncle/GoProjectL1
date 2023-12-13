package main

import "fmt"

// Human - структура, представляющая человека.
type Human struct {
	Name string // Имя человека.
	Age  int    // Возраст человека.
}

// Speak - метод, позволяющий человеку говорить.
func (h *Human) Speak() {
	fmt.Println(h.Name, "is speaking")
}

// Action - структура, представляющая действие, которое может выполнять человек.
type Action struct {
	Human             // Встраивание структуры Human.
	ActionType string // Тип действия.
}

// PerformAction - метод, позволяющий человеку выполнять действие.
func (a *Action) PerformAction() {
	fmt.Println(a.Name, "is performing", a.ActionType)
}

func main() {
	// Создание экземпляра структуры Action.
	action := Action{
		Human: Human{
			Name: "Danila", // Имя человека.
			Age:  22,       // Возраст человека.
		},
		ActionType: "working", // Тип действия.
	}

	// Вызов метода Speak.
	action.Speak()

	// Вызов метода PerformAction.
	action.PerformAction()
}
