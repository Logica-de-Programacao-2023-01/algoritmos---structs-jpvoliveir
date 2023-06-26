package main

import (
	"fmt"
)

type Viagem struct {
	Origem  string
	Destino string
	Data    string
	Preco   float64
}

func ViagemMaisCara(viagens []Viagem) Viagem {
	if len(viagens) == 0 {
		return Viagem{}
	}

	viagemMaisCara := viagens[0]
	for _, viagem := range viagens {
		if viagem.Preco > viagemMaisCara.Preco {
			viagemMaisCara = viagem
		}
	}

	return viagemMaisCara
}

func main() {
	viagens := []Viagem{
		{Origem: "São Paulo", Destino: "Rio de Janeiro", Data: "2023-07-15", Preco: 250.00},
		{Origem: "Belo Horizonte", Destino: "Salvador", Data: "2023-08-20", Preco: 350.00},
		{Origem: "Brasília", Destino: "Fortaleza", Data: "2023-09-10", Preco: 400.00},
	}

	viagemMaisCara := ViagemMaisCara(viagens)
	fmt.Println("Viagem mais cara:")
	fmt.Println("Origem:", viagemMaisCara.Origem)
	fmt.Println("Destino:", viagemMaisCara.Destino)
	fmt.Println("Data:", viagemMaisCara.Data)
	fmt.Println("Preço:", viagemMaisCara.Preco)
}
