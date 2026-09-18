package main

import "fmt"

func main() {
	marked := []int{9, 10, 11, 12, 13, 14, 25, 26}
	total := 0

	for _, day := range marked {
		weekday := (day - 1) % 7 // 0=Пн, 1=Вт, ..., 6=Вс
		if weekday <= 3 {
			total += 2100
		} else {
			total += 2850
		}
	}
	fmt.Print(total) // 19800
}