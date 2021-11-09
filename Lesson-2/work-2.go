package main

import (
	"fmt"
	"math"
)

func main() {
	var err error
	var s float64
	fmt.Println("Введите значение площади круга")
	_, err := fmt.Scanln(&s)
	if err != nil {
		fmt.Println("Ошибка !!", err.Error())

		resultD := math.Sqrt(s/math.Pi) * 2 //формула вычисления площади круга D=√S/Пи

		fmt.Printf("Площадь круга равна: %f\n", resultD)

	}
}