package numbers

import "errors"

// GetVariables regresa tres variables para ser asignadas o mostradas
func GetVariables() (int, int32, int64) {
	return 1, 2, 20
}

// GetFloat regresa dos numeros flotantes
func GetFloat() (float32, float64) {
	return float32(0.1), float64(float32(0.1))
}

// Sum suma dos números enteros y devuelve el resultado
func Sum(a interface{}, b interface{}) (int, error) {
	switch a.(type) {
	case string:
		return 0, errors.New("El valor A es un string")
	}
	switch b.(type) {
	case string:
		return 0, errors.New("El valor B es un string")
	}
	return a.(int) + b.(int), nil
}