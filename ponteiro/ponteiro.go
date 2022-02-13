package ponteiro

import "fmt"

var z *int

func ExemploPonteiro() {
	x := 10
	y := &x
	z = &x

	// Valor da variável x
	fmt.Println(x)

	// Valor da variável y (endereço de memória da variável x)
	fmt.Println(y)

	// valor da variável z (endereço de memória da variável x)
	fmt.Println(z)

	// Tipo das variáveis
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)

	// Acessando valor da variável x pelos ponteiros (variáveis y e z)
	fmt.Println(*y)
	fmt.Println(*z)
}
