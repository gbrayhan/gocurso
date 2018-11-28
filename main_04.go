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