package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{1, 3, 4, 7, 10, 13, 15, 17, 18}
	target := 10

	sort.Ints(numbers)

	index := sort.SearchInts(numbers, target)

	if index < len(numbers) && numbers[index] == target {
		fmt.Printf("Number found at index: %d\n", index)
	} else {
		fmt.Println("Number not found.")
	}

}
