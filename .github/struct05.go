package main

import (
	"fmt"
	"time"
)

type Funcionario struct {
	Nome    string
	Salario float64
	Idade   int
}

func (f *Funcionario) AumentarSalario(porcentagem float64) {
	aumento := f.Salario * (porcentagem / 100)
	f.Salario += aumento
}

func (f *Funcionario) DiminuirSalario(porcentagem float64) {
	desconto := f.Salario * (porcentagem / 100)
	f.Salario -= desconto
}

func (f *Funcionario) TempoDeServico() int {
	idadeInicioTrabalho := 18
	idadeAtual := time.Now().Year() - f.Idade
	tempoServico := idadeAtual - idadeInicioTrabalho
	return tempoServico
}

func main() {
	funcionario := Funcionario{
		Nome:    "João",
		Salario: 3000,
		Idade:   25,
	}

	fmt.Printf("Salário inicial do funcionário %s: %.2f\n", funcionario.Nome, funcionario.Salario)

	funcionario.AumentarSalario(10)
	fmt.Printf("Salário após aumento: %.2f\n", funcionario.Salario)

	funcionario.DiminuirSalario(5)
	fmt.Printf("Salário após desconto: %.2f\n", funcionario.Salario)

	tempoServico := funcionario.TempoDeServico()
	fmt.Printf("Tempo de serviço do funcionário: %d anos\n", tempoServico)
}
