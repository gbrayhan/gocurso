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
