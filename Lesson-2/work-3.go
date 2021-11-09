package main

import "fmt"

func main() {
	fmt.Println("Задание 3 - Выведите цифры, соответствующие количество сотен, десятков и единиц  в этом числе.")
	fmt.Println("Введите трехзначное число")
	var numbers string
	_, err := fmt.Scanln(&numbers)
	if err != nil {
		fmt.Println("Ошибка !!", err.Error())
	}
	if len(numbers) == 3 { //проверяем длинну числа должна равняться 3
		fmt.Println("Число:", numbers)
		fmt.Println("Сотни: ", string(numbers[0]))
		fmt.Println("Десятки: ", string(numbers[1]))
		fmt.Println("Еденицы: ", string(numbers[2]))
	} else {
		fmt.Println("Вы ввели не верное згначение. !! Число не трехзначное !!")
	}
}
