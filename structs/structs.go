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