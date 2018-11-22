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
