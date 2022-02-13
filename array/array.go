package array

import "fmt"

func InvocarFuncoes() {
	ExemploSlice()
	ExemploMap()
}

func ExemploSlice() {
	s := make([]int, 5)

	s = append(s, 1, 2, 3)

	fmt.Println(s)

	sliceString := []string{
		"Carro",
		"Moto",
		"Bicicleta",
	}

	sliceString = append(sliceString, "Aviao")

	fmt.Println(sliceString)
}

func ExemploMap() {
	m := make(map[string]int)

	m["Fulano"] = 10
	m["Ciclano"] = 11

	fmt.Println(m["Fulano"])
}
