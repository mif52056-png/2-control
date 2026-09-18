package main

import "fmt"

func main() {
	var mainBag, handBag, extraHandBag float64
	fmt.Print("Вес основного багажа: ")
	fmt.Scanln(&mainBag)
	fmt.Print("Вес ручной клади: ")
	fmt.Scanln(&handBag)
	fmt.Print("Вес доп. ручной клади: ")
	fmt.Scanln(&extraHandBag)

	fmt.Printf("Общий вес: %.2f\n", mainBag+handBag+extraHandBag)
}