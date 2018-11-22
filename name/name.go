package name

import "fmt"

// GetName regresa el nombre que se ingresa desde teclado
func GetName() string {
	var name string = "Default"
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name)

	return name
}