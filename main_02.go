package main

import (
	"fmt"
	"github.com/gbrayhan/gocurso/maps"
	"github.com/gbrayhan/gocurso/numbers"
	"github.com/gbrayhan/gocurso/structs"
	"time"
)

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

func main() {

	fmt.Println( maps.GetMap("Alejandra",21))
	structs.InterfaceTest()

	//structs.InterfaceTest()
	number, err := numbers.Sum("50", 50)
	if err != nil {
		//panic(err)
		fmt.Println(err)
	}
	fmt.Println(number)
	pointerTest()

	go forGo(500)
	go forGo(400)
	time.Sleep(10000 * time.Millisecond)
}

func helloGo(index int) {
	fmt.Println("Hola soy un print en la Gorutine #",index)
}

func forGo(n int) {
	for i := 0 ; i < n ; i++  {
		go helloGo(i)
	}
}