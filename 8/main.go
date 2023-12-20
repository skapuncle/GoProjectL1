package main

import "fmt"

// setBit - устанавливает i-й бит в 1.
func setBit(n int64, i uint) int64 {
	return n | (1 << i)
}

// clearBit - устанавливает i-й бит в 0.
func clearBit(n int64, i uint) int64 {
	mask := ^(1 << i)
	return n & int64(mask)
}

func main() {
	var n int64
	var i uint

	fmt.Scan(&n, &i)

	// Выводим исходное число в двоичной системе.
	fmt.Printf("Original number: %b\n", n)

	// Устанавливаем i-й бит в 1.
	n = setBit(n, i)
	fmt.Printf("After setting bit to 1: %b\n", n)

	// Устанавливаем i-й бит в 0.
	n = clearBit(n, i)
	fmt.Printf("After setting bit to 0: %b\n", n)

}
