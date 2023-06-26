package main

import "fmt"

type Animal struct {
	Nome    string
	Especie string
	Idade   int
	Som     string
}

func (a *Animal) AlterarSom(novoSom string) {
	a.Som = novoSom
}

func (a Animal) ImprimirInformacoes() {
	fmt.Println("Nome:", a.Nome)
	fmt.Println("Espécie:", a.Especie)
	fmt.Println("Idade:", a.Idade)
	fmt.Println("Som:", a.Som)
}

func main() {
	animal := Animal{
		Nome:    "Bob",
		Especie: "Cachorro",
		Idade:   3,
		Som:     "Latido",
	}

	fmt.Println("Informações do animal:")
	animal.ImprimirInformacoes()

	novoSom := "Miado"
	animal.AlterarSom(novoSom)

	fmt.Println("Novas informações do animal:")
	animal.ImprimirInformacoes()
}
