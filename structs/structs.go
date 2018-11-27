package structs

import "fmt"

// GetArray imprime un array definido
func GetArray() {
	var arr1 [2]string
	arr1[0] = "array"
	arr1[1] = "array2"
	fmt.Println(arr1)
}

// GetSlice implime un slice definido
func GetSlice(name string) {
	var slice1[]string
	slice1 = append(slice1, "hola", "como", "estas")
	slice1 = append(slice1, name)
	fmt.Println(slice1)

}

// PlatziCourse es una estructura 
type PlatziCourse struct {
	Name   string
	Slug   string
	Skills []string
}

// PlatziCareer es una estrctura que toma los metodos de PlatziCourse
type PlatziCareer struct {
	PlatziCourse
}

// Subscribe es un metodo de PlatziCourse
func (p PlatziCourse) Subscribe(name string) {
	fmt.Printf("La persona %s se ha registrado al curso %s \n", name, p.Name)
}

// Subscribe es un metodo de PlatziCareer
func (p PlatziCareer) Subscribe(name string) {
	fmt.Printf("La persona %s se ha registrado a la carrera %s \n", name, p.Name)
}
