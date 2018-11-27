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
## Implementando interfaces en **go**
Cuando hay herencia entre structs y ambos tienen un mismo metodo que realiza diferentes acciones, usamos **interfcaes** en go.

```golang
package structs
// Platzi es una interface de PlatziCourse y PlatziCareer
type Platzi interface {
	Subscribe(name string)
}
// InterfaceTest muestra como es implementada la interface
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
```
## Defer
Los defer en go son instrucciones que se implementan en una funcion y que se ejecutan solamente al final de dicha funcion.
```golang
func deferTest() {
	fmt.Println("La funcion InterfaceTest ha terminado")
}
// InterfaceTest muestra como es implementada la interface
func InterfaceTest() {
	defer deferTest()
	platziCourse := PlatziCourse{Name: "Go", Slug: "go", Skills: []string{"backend"}}
}
```
