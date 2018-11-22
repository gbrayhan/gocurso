# Empezando con GO  
  
Debemos definir el entrono de trabajo en **Go**, esto lo declaramos con el siguiente comando:  
`export GOPATH=/var/www/go`  
  
Luego de esto es necesario definir algunas carperas para la correcta importacion de las librerias en go, ejemplo:    
**/var/www/go/src/github.com/gbrayhan/gocurso**  
  
Para ejecutar un programa de **Go** sin crear el ejecutable debemos de usar el siguiente comando:  
`go run main.go`   
  
Ejemplo **Hola mundo** de la primera Clase:  
  
```golang  
package main

import "fmt"

func main() {
	var name string = "Default"
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name)
	fmt.Printf("Hola %s \n", name)
	fmt.Println("Hasta pronto!")

} 
``` 
  
Se pueden declarar de diferente manera las variables, tal como se muestra a continuacion:  
  
```golang  
package main  
  
import "fmt"  
  
const holaMundo string = "Hola %s %s bienvenido al curso de golang.\n"  
  
func main() {  
	var name string = "Default"  
	name = "mi nombre"  
	lastname := "Gabriel"  
	var age = 25  
  
	var (  
		a = 12  
		b = 15  
		c = 45  
	)  
  
	fmt.Print("Ingresa tu nombre: ")  
	fmt.Scanf("%s", &name)  
	fmt.Printf(holaMundo, name, lastname)  
	fmt.Printf("Tu edad es: %d \n", age)  
	fmt.Printf("Los valores de las variables que se declararon son: %d %d %d \n", a, b, c)  
	fmt.Println("Hasta pronto!")  
}
```  
  
Funciones en **golang**   
   
Se puede imlementar en las funciones un multiple retorno tal como se muestra a continuacion en la funcion **getVariables()**  
    
```golang  
package main

import "fmt"

const holaMundo string = "Hola %s %s bienvenido al curso de golang.\n"

func main() {
	name := getName()
	lastname := "Gabriel"
	age := sum(20, 5)
	a, b, c := getVariables()

	fmt.Printf(holaMundo, name, lastname)
	fmt.Printf("Tu edad es: %d \n", age)
	fmt.Printf("Los valores de las variables que se declararon son: %d %d %d \n", a, b, c)
	fmt.Println("Hasta pronto!")

}

func getName() string {
	var name string = "Default"
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name)

	return name
}

func getVariables() (int, int, int) {
	return 1, 2, 3
}

func sum(a int, b int) int {
	return a + b
}
```  

