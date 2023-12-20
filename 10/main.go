package main

import (
	"fmt"
	"math"
)

func main() {
	// Исходные данные.
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Создаем map для хранения групп температур.
	groups := make(map[int][]float64)

	// Проходим по каждой температуре.
	for _, temp := range temperatures {
		// Округляем температуру вниз до ближайшего числа, кратного 10.
		group := int(math.Floor(temp/10)) * 10
		// Если температура меньше следующего шага, вычитаем 10.
		if temp < 0 {
			group += 10
		}
		// Добавляем температуру в соответствующую группу.
		groups[group] = append(groups[group], temp)
	}

	// Выводим результат.
	for group, temperatures := range groups {
		fmt.Printf("%d: %v\n", group, temperatures)
	}
}
