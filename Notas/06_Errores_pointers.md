## Errores en **go**  
El manejo de errrores se puede establecer de una manera más sencilla debido al multiple retorno de las funciones o metodos, ya que juntamente con el valor esperado, podemos retornar el error si es que hay alguno.  
  
```golang  
// Sum suma dos números enteros y devuelve el resultado
func Sum(a interface{}, b interface{}) (int, error) {
	switch a.(type) {
	case string:
		return 0, errors.New("El valor A es un string")
	}
	switch b.(type) {
	case string:
		return 0, errors.New("El valor B es un string")
	}
	return a.(int) + b.(int), nil
}
```  
Se manda a llamar de la siguiente manera y se trabaja con el error.  
  
```golang  
number, err := numbers.Sum("50", 50)
if err != nil {
	//panic(err)
	fmt.Println(err)
}
fmt.Println(number)
```  
## Punteros en **go**  
Nos permiten establecer variables que apuntan a direcciones de memoria para optimizar recursos.  
  
```golang  
func pointerTest(){
	a := 100
	var b * int
	b = &a
	fmt.Println("Sin modificar")
	fmt.Println(a, *b)
	fmt.Println(&a, b)
	pointerModify(b)
	fmt.Println("Con una modificación")
	fmt.Println(a, *b)
	fmt.Println(&a, b)
}

func pointerModify(c  *int) {
	*c = 10
}
```  

