package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Введите трехзначное число")
	var numbers string
	_, err := fmt.Scanln(&numbers)
	if err != nil {
		fmt.Println("Ошибка !!", err.Error())
		os.Exit(1)
	}
	if len(numbers) == 3 {
		fmt.Println("Сотни: ", string(numbers[0]))
		fmt.Println("Десятки: ", string(numbers[1]))
		fmt.Println("Еденицы: ", string(numbers[2]))
	} else {
		fmt.Println("Ошибка!!! Введите трехзначное число")
	}
}
