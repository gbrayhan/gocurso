# Tipos de variables en **golang**, uso de arrays, slices y condicionales
  
En el siguiente codigo se muestra como es que se declaran las variables de 32 y 64 bits, aunque se recomienda que se debe considerar si nuestra aplicacion se requiere correr en sistemas de 32 y 64 bits.  
  
```golang  
func main() {
	name := getName()
	lastname := "Gabriel"
	age := sum(20, 5)
	a, b, c := getVariables()
	f32, f64 := getFloat()

	fmt.Printf(holaMundo, name, lastname)
	fmt.Printf("Tu edad es: %d \n", age)
	fmt.Printf("Los valores de las variables que se declararon son: %d %d %d \n", a, b, c)
	fmt.Println("A continuacion se muestra la presicion de las variables float32 y float64.")
	fmt.Printf("float32: %.20f y float62 %.20f \n", f32, f64)
	fmt.Println("Hasta pronto!")

}

func getVariables() (int, int32, int64) {
	return 1, 2, 20
}

func getFloat() (float32, float64) {
	return float32(0.1), float64(float32(0.1))
}
```  
  
Go soporta multiples lenguajes nativamente, tambien hay algunas funciones de string que permiten trabajar con datos del texto, en el siguiente codigo se muestra la declaracion 


```golang   

func main() {

	fmt.Println(string ("Hasta pronto!"[3]))
	fmt.Println("cantidad de letras que tiene hola: ", len("hola"))
	// imprimir()
}
func imprimir() { // para imprimir los caracteres individuales
	str := "こんにちは世界"
	for i := 0; i < len(str); i++ {
		fmt.Println(string(str[i]))
	}
}  

```   
  
En **golang** se pueden definir arrays y slices, que permiten trabajar con **colecciones** de datos, la diferencia el uno con el otro es que el slice es dinamico y su implementacion es como se muestra a continuacion:  

```golang
func getArray() {
	var arr1 [2]string
	arr1[0] = "array"
	arr1[1] = "array2"
	fmt.Println(arr1)
}

func getSlice(name string) {
	var slice1[]string
	slice1 = append(slice1, "hola", "como", "estas")
	slice1 = append(slice1, name)
	fmt.Println(slice1)

}

```
  
Las condicionales en **go** se realizan de la siguiente manera, podemos hacer una asignacion o tarea antes de entrar al condicional.  
  
```golang   
func ifTest() {
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
```    
