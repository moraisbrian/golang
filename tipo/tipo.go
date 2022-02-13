package tipo

import (
	"encoding/json"
	"fmt"
)

// type CarName string
// type CarYear int

type Car struct {
	CarName string `json:"car"`
	CarYear int    `json:"year"`
}

func InvocarFuncoes() {
	ManipulacaoJson()
	UtilizandoFuncaoStruct()
}

func ManipulacaoJson() {
	car := Car{
		CarName: "Fusca",
		CarYear: 2022,
	}

	result, _ := json.Marshal(car)
	fmt.Println(string(result))

	j := []byte(`{"car": "Celta", "year": 2021}`)
	var car2 Car
	json.Unmarshal(j, &car2)
	fmt.Println(car2)
}

// Se utilizado ponteirio (*) o valor do objeto é alterado por referência
// Se não utilizado ponteiro (*) o valor é alterado apenas para a execução da função
func (c *Car) drive() {
	c.CarName = "Brasilha"
	fmt.Println(c.CarName, "Andou")
}

func UtilizandoFuncaoStruct() {
	car1 := Car{
		CarName: "Fusca",
		CarYear: 2022,
	}

	car1.drive()
	fmt.Println(car1)
}
