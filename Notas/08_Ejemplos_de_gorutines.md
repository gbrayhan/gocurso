## Ejemplos de gorutines  
 
A continuacion veremos  el programa de ping-pong usando **gorutines** y chanels para esperar la respuesta de nuestras rutinas.  

```golang   
package main

import (
	"fmt"
	"time"
)

// [chan de escritura] chan <- para evitar errores futuros en el código al tratar de leer el canal en este código
func ping(ball chan <- int, action chan <- string) {
	time.Sleep(4000 * time.Millisecond)
	ball <- 1
	action <- "Player ping"
}

func pong(ball chan <- int, action chan <- string) {
	ball <- 2
	action <- "Player pong"
}
// referee es un loop que siempre pregunta la accion
func referee(action <- chan string) {
	for {
		fmt.Println("Aqui es donde espera la respuesta del job")
		fmt.Println("Action: ", <- action)
		fmt.Println("Aqui acaba de revibir la respuesta")
	}
}

func pingpong() {
	ball := make(chan int)
	action := make(chan string)

	go referee(action)
	go ping(ball, action) // Primer Saque

	for {
		value := <- ball
		switch value {

		case 1:
			go pong(ball, action)
		case 2:
			go ping(ball, action)
		default: // No entra nunca aqui
			fmt.Println("No vale nada")
		}
	}
}

func main() {
	fmt.Printf("Comienza el juego de ping pong por 200 segundos\n")
	go pingpong()
	time.Sleep(200 * time.Second)
}
```  
Ejemplo de la documentacion oficial de golang de la sincronizacion de canales con la gorutine.  

```golang  
/**
 * Podemos usar canales para sincronizar la ejecución a través de goroutines.
 * Aquí hay un ejemplo del uso de una recepción de bloqueo para
 * esperar a que un goroutine termine.
 */
package main

import "fmt"
import "time"

/**
 * Esta es la función que ejecutaremos en una goroutine.
 * El canal 'done' se usará para notificar a otra goroutine
 * que el trabajo de esta función está hecho.
 */
func worker(done chan bool) {
	fmt.Print("Trabajando...")
	time.Sleep(2 * time.Second)
	fmt.Println("Hecho")

	// Enviar un valor para notificar que hemos terminado.
	done <- true
}

func main() {

	// Poner en marcha la gorutine, dandole un canal (make) para notificar
	done := make(chan bool, 1)
	go worker(done)

	// Bloquear hasta que recibamos una notificación del trabajo terminado en el canal.
	<-done
} 
``` 




