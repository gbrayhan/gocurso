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

func main() {

	fmt.Println( maps.GetMap("Alejandra",21))

	platziCourse := PlatziCourse{Name: "Go", Slug: "go", Skills: []string{"backend"}}
	platziCourse1 := new(PlatziCourse)
	platziCourse1.Name = "Go1"
	platziCourse1.Slug = "go1"
	platziCourse1.Skills = []string{"backend1"}
	fmt.Println(platziCourse)
	
}