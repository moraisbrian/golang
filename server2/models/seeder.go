package models

func Seed() *Products {
	produto1 := NewProduct()
	produto1.Name = "Martelo"

	produto2 := NewProduct()
	produto2.Name = "Prego"

	products := Products{}
	products.Add(*produto1)
	products.Add(*produto2)

	return &products
}
