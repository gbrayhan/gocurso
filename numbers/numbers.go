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