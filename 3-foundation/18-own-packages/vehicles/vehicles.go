package vehicles

type Carro struct {
	Marca string
}

func (c Carro) Andar() string {
	return "Carro andando"
}
