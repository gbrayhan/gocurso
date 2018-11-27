package main

import (
	"fmt"
	"github.com/gbrayhan/gocurso/maps"
	"github.com/gbrayhan/gocurso/numbers"
	"github.com/gbrayhan/gocurso/structs"
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

}