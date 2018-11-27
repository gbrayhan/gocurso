## Metodos en **go**
Se aplican a las estructuras y nos permiten trabajar de manera similar a la que lo hacemos con una clase en POO.
  
```golang   
func (p PlatziCourse) Subscribe(name string) {
	fmt.Printf("La persona %s se ha registrado al curso %s \n", name, p.Name)	
}
func main() {
	platziCourse1 := new(PlatziCourse)
	platziCourse1.Name = "Go"
	platziCourse1.Slug = "go"
	platziCourse1.Skills = []string{"backend"}
	platziCourse1.Subscribe("Arturo")
}
```    

