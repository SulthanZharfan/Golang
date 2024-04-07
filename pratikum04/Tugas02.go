package main

import "fmt"

func main() {
	fmt.Println("Bubble Sort")
	arrayNumber := [20]int{10, 8, 4, 6, 9, 3, 5, 7, 2, 1, 20, 18, 14, 16, 19, 13, 15, 17, 12, 11}

	for i := 0; i < len(arrayNumber); i++ {
		for j := 0; j < len(arrayNumber)-1; j++ {
			if arrayNumber[j] > arrayNumber[j+1] {
				// Swap
				arrayNumber[j], arrayNumber[j+1] = arrayNumber[j+1], arrayNumber[j]
			}
		}
	}

	fmt.Println("Setelah dilakukan pengurutan:")
	fmt.Println(arrayNumber)
}
