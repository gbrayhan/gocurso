## Metodos de la libreria **strings** 

Se pueden implementar diferentes metodos de la libreria strings tal como se muestra a continuacion:  
  
```golang  

func strings2() {
	text := "Hello world, Hello CPA, Hello Go"

	fmt.Println(strings.ToUpper(text))
	fmt.Println(strings.ToLower(text))
	fmt.Println(strings.Replace(text,"Hello", "Hola",-1))
	fmt.Println(strings.Split(text, ","))
	
}
```  
  
## Uso de la sentencia **switch**

una de las diferencias en el swithc es que **golang** permite implementar condicionales dentro del switch, reemplazando los multiples else if que pudieran implementarse y así tener orden en las comparaciones y casos distintos:  

```golang  
func switchTest() {
	var number int
	fmt.Print("Type a number: ")
	fmt.Scanf("%d", &number)

	switch number {
		case1:
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

```  