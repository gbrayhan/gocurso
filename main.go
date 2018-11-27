package main

import(
	"github.com/gbrayhan/gocurso/maps"
	"fmt"
)

type PlatziCourse struct {
	Name   string
	Slug   string
	Skills []string
}

type PlatziCareer struct {
	PlatziCourse
}

func (p PlatziCourse) Subscribe(name string) {
	fmt.Printf("La persona %s se ha registrado al curso %s \n", name, p.Name)
	
}

func main() {

	fmt.Println( maps.GetMap("Alejandra",21))
	platziCourse := PlatziCourse{Name: "Go", Slug: "go", Skills: []string{"backend"}}
	platziCourse1 := new(PlatziCourse)
	platziCourse1.Name = "Go"
	platziCourse1.Slug = "go"
	platziCourse1.Skills = []string{"backend1"}
	
	platziCourse1.Subscribe("Arturo")

	fmt.Println(platziCourse)
	
}