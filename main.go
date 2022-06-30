package main

import "fmt"

func main() {
	var operator string
	var number1, number2 int

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&number1)
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&number2)
	fmt.Print("Выберите действие (+,-,/,%,*):")
	fmt.Scanln(&operator)

	output := 0
	switch operator {
	case "+":
		output = number1 + number2
		break
	case "-":
		output = number1 - number2
	case "*":
		output = number1 * number2
	case "/":
		output = number1 / number2
	case "%":
		output = number1 % number2
	default:
		fmt.Println("Неверно выбрано действие")
	}

	fmt.Printf("%d %s %d = %d", number1, operator, number2, output)
}
