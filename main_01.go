package main

import(
	"fmt"
	"strings"

	"github.com/gbrayhan/gocurso/name"
	"github.com/gbrayhan/gocurso/structs"
	"github.com/gbrayhan/gocurso/flow"
	"github.com/gbrayhan/gocurso/numbers"

)

const holaMundo string = "Hola %s %s bienvenido al curso de golang.\n"

func main() {
	firstName := name.GetName()
	lastname := "Gabriel"
	age := numbers.Sum(20, 5)
	a, b, c := numbers.GetVariables()
	f32, f64 := numbers.GetFloat()

	array := [3]int{2,1,3} // no se puede exceder

	fmt.Printf(holaMundo, firstName, lastname)
	fmt.Printf("Tu edad es: %d \n", age)
	fmt.Printf("Los valores de las variables que se declararon son: %d %d %d \n", a, b, c)
	fmt.Println("A continuacion se muestra la presicion de las variables float32 y float64.")
	fmt.Printf("float32: %.20f y float62 %.20f \n", f32, f64)
	fmt.Println(string ("Hasta pronto!"[3]))
	fmt.Println("cantidad de letras que tiene hola: ", len("hola"))
	fmt.Println(array)

	// imprimir()

	structs.GetArray()
	structs.GetSlice(firstName)

	flow.IfTest()
	strings2()

	flow.SwitchTest()

}




func imprimir() { 
	str := "こんにちは世界"
	for i := 0; i < len(str); i++ {
		fmt.Println(string(str[i]))
	}
}

func strings2() {
	text := "Hello world, Hello CPA, Hello Go"

	fmt.Println(strings.ToUpper(text))
	fmt.Println(strings.ToLower(text))
	fmt.Println(strings.Replace(text,"Hello", "Hola",-1))
	fmt.Println(strings.Split(text, ","))
	
}

