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