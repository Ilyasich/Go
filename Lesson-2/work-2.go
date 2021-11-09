package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var err error
	fmt.Println("Введите значение площади круга")
	var s float64
	_, err = fmt.Scanln(&s)
	if err != nil {
		fmt.Println("Ошибка !!", err.Error())
		os.Exit(1)
	}
		resultD := math.Sqrt(s/math.Pi) * 2 //формула вычисления площади круга D=√S/Пи

		fmt.Printf("Площадь круга равна: %f\n", resultD)

	}
