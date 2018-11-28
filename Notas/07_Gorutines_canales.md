## Rutinas en **go**  
  
Las rutinas en go nos permiten implementar procesamiento en paralelo, las gorutines se implementan de manera nativa lo cual garantiza la efectividad de las mismas.  
```golang  
func helloGo(index int) {
	fmt.Println("Hola soy un print en la Gorutine #",index)
}
func forGo(n int) {
	for i := 0 ; i < n ; i++  {
		go helloGo(i)
	}
}
func main() {
	go forGo(500)
	go forGo(400)
	time.Sleep(10000 * time.Millisecond)
}
```  
  
## Canales en **go**  
Nos permiten trabajar de una manera correcta con las **gorutines**.  


```golang
package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		defer close(ch) // Evita que el programa sobreescriba algún valor erroneo
		ch <- "Hola Chanel"
	}()
	fmt.Println("El valor de ch es: ", <- ch)

	ch1 := make(chan int)
	go func() {
		defer close(ch1)
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
		ch1 <- 4
		ch1 <- 5
	}()
	fmt.Println(ch1)
	for n := range ch1 {
		fmt.Printf("El valor de ch1 es: %d \n",n)
	}

	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 23

	fmt.Println("El valor de ch2 es: ", <- ch2)
	fmt.Println("El valor de ch2 es: ", <- ch2)
	close(ch2)
}
```  
# Canales en **go**


Los canales nos ayudan a trabar mejor con las **gorutines**, a continuacion mostraremos como se implementan los chanels en nuestras rutinas de go.  
  
```golang
package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		defer close(ch) // Evita que el programa sobreescriba algún valor erroneo
		ch <- "Hola Chanel"
	}()
	fmt.Println("El valor de ch es: ", <- ch)

	ch1 := make(chan int)
	go func() {
		defer close(ch1)
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
		ch1 <- 4
		ch1 <- 5
	}()
	fmt.Println(ch1)
	for n := range ch1 {
		fmt.Printf("El valor de ch1 es: %d \n",n)
	}

	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 23

	fmt.Println("El valor de ch2 es: ", <- ch2)
	fmt.Println("El valor de ch2 es: ", <- ch2)
	close(ch2)

}
```  