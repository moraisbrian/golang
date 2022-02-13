package interfaces

import "fmt"

type veiculo interface {
	start() string
}

func (c Car) start() string {
	return "Iniciou"
}

type Car struct {
	CarName string
	CarYear int
}

func ExemploInerface(c veiculo) {
	fmt.Printf("Carro com a interface veiculo: %v", c)
}

func UtilizandoInterface() {
	car := Car{
		CarName: "Fusca",
		CarYear: 2022,
	}

	ExemploInerface(car)
}
