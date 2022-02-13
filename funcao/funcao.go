package funcao

import "errors"

func InvocarFuncoes() {
	a, b, err := Nome("brian", 10)

	if err != nil {
		print(err.Error())
	} else {
		print(a, " ", b)
	}

	print("\n")

	ref := ObterValorReferencia(&b)
	print(ref)
}

func Nome(a string, b int) (string, int, error) {
	if b > 10 {
		return "", 0, errors.New("Error")
	}

	return a, b, nil
}

func ObterValorReferencia(valor *int) int {
	return *valor
}
