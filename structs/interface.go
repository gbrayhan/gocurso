package structs

// Platzi es una interface de PlatziCourse y PlatziCareer
type Platzi interface {
	Subscribe(name string)
}

func InterfaceTest() {
	platziCourse := PlatziCourse{Name: "Go", Slug: "go", Skills: []string{"backend"}}

	platziCareer := new (PlatziCareer)
	platziCareer.Name = "Go"
	platziCareer.Slug = "go"

	callSubscribe(platziCourse)
	callSubscribe(platziCareer)
}

func callSubscribe(p Platzi) {
	p.Subscribe("Alejandro")
}