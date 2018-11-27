## Paquetes en **golang**
  
En golang se puede usar paquetes y de una manera muy sencilla, se debe seguir la estructura de carpetas y de importacion de capetas, se utiliza la palabra reservada **package**.  
  
Ejemplo de un paquete:  
  
```golang  
package numbers

// GetVariables regresa tres variables para ser asignadas o mostradas
func GetVariables() (int, int32, int64) {
	return 1, 2, 20
}

// GetFloat regresa dos numeros flotantes
func GetFloat() (float32, float64) {
	return float32(0.1), float64(float32(0.1))
}

// Sum devuelve el valor de la suma de dos enteros
func Sum(a int, b int) int {
	return a + b
}
```  
Es necesario poner la funcion con la primera letra **mayuscula** para indicarle a go que es una funcion publica dentro del paquete.  
  
```golang  
import(
	"github.com/gbrayhan/gocurso/name"
	"github.com/gbrayhan/gocurso/structs"
)

age := numbers.Sum(20, 5)
a, b, c := numbers.GetVariables()
f32, f64 := numbers.GetFloat()

```   
  
## Maps en **golang**   
Los mapas en go nos facilitan implementar colecciones de datos por **llave-valor**  
  
```golang   
package maps

func GetMap(name string, edad int ) map[string]int {
	// Los mapas necesitan ser inicializados con la palabra make
	/*
	mapTest := make(map[string]int)
	mapTest["llave1"] = 1
	mapTest["llave2"] = 23
	*/

	mapTest := map[string]int {
		"Juan": 13,
		"Alejandro": 25,
		"Andrea": 19,
		"Juana": 32,
		name: edad,
	}

	delete(mapTest, "andrea")

	return mapTest
}
```  
## Structs en **go**  
  
Las estructuras nos permiten llevar a cabo colecciones de datos, bien definidas por un nombre. Dado que **go** no es un lenguaje orientado a objetos, la herramienta que permite delimitar un elemento por sus caracteristicas es **struct**.  
  
```golang   
type PlatziCourse struct {
	Name   string
	Slug   string
	Skills []string
}
type PlatziCareer struct {
	PlatziCourse
}
func main() {
	platziCourse := PlatziCourse{Name: "Go", Slug: "go", Skills: []string{"backend"}}
	platziCourse1 := new(PlatziCourse)
	platziCourse1.Name = "Go1"
	platziCourse1.Slug = "go1"
	platziCourse1.Skills = []string{"backend1"}
	fmt.Println(platziCourse)
}
```  

