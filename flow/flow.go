package flow

import "fmt"

// IfTest evalua si un numero es par o imprar y si su valor es 3
func IfTest() {
	var number = 0
	fmt.Print("Ingresa un número del 1 al 10: ")
	fmt.Scanf("%d", &number)
	if number%2 == 0 {
		fmt.Println("El número es par")
	} else {
		fmt.Println("El número es impar")
	}

	if number2 := 3; number2 == 3 {
		fmt.Println("Entró al condicional")
	}
}

// SwitchTest evalua si un numero es 1 y si es o no par
func SwitchTest() {
	var number int
	fmt.Print("Type a number: ")
	fmt.Scanf("%d", &number)

	switch number {
		case 1:
		fmt.Println("The number is 1")
		default:
			fmt.Println("The number isn't 1")
	}

	switch {
		case number % 2 == 0:
		fmt.Println("The number is pair")
		default:
			fmt.Println("The number is odd")
	}
}