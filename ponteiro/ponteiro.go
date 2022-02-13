package ponteiro

import "fmt"

// & -> Referência de memória, obtem o endereço de memória
// * -> Ponteiro, aponta para o valor da variável passada por referência

var z *int

func ExemploPonteiro() {
	x := 10
	y := &x
	z = &x

	// Valor da variável x
	fmt.Println(x)

	// Valor da variável y (endereço de memória da variável x)
	fmt.Println(y)

	// Valor da variável z (endereço de memória da variável x)
	fmt.Println(z)

	// Tipo das variáveis
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)

	// Acessando valor da variável x pelos ponteiros (variáveis y e z)
	fmt.Println(*y)
	fmt.Println(*z)
}
