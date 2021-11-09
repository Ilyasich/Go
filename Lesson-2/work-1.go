package main

import (
	"fmt"
)

func main() {

	var a, b float64

	fmt.Println("Введите значение высоты: ")
	fmt.Scanln(&a)

	fmt.Println("Введите значение ширины: ")
	fmt.Scanln(&b)

	fmt.Printf("Площадь прямоугольника равна: %f\n", a*b)

}
