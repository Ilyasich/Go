package main

import (
	"fmt"
	"os"
)

func main() {
	var err error
	var a, b float64

	fmt.Println("Введите значение высоты: ")
	_, err = fmt.Scanln(&a)
	if err != nil {
		fmt.Println("Ошибка!" , err.Error())
		os.Exit(1)
	}

	fmt.Println("Введите значение ширины: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Println("Ошибка! Неверный ввод.", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Площадь прямоугольника равна: %f\n", a*b)

}
