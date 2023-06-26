package main

import (
	"fmt"
)

type Aluno struct {
	Nome  string
	Idade int
	Notas []float64
}

func (a *Aluno) AdicionarNota(nota float64) {
	a.Notas = append(a.Notas, nota)
}

func (a *Aluno) RemoverNota(index int) {
	if index >= 0 && index < len(a.Notas) {
		a.Notas = append(a.Notas[:index], a.Notas[index+1:]...)
	}
}

func (a Aluno) CalcularMedia() float64 {
	soma := 0.0
	for _, nota := range a.Notas {
		soma += nota
	}
	media := soma / float64(len(a.Notas))
	return media
}

func (a Aluno) ImprimirInformacoes() {
	fmt.Println("Nome:", a.Nome)
	fmt.Println("Idade:", a.Idade)
	fmt.Println("Média:", a.CalcularMedia())
}

func main() {
	aluno := Aluno{
		Nome:  "João",
		Idade: 20,
		Notas: []float64{8.5, 7.8, 9.2},
	}

	fmt.Println("Informações do aluno:")
	aluno.ImprimirInformacoes()

	aluno.AdicionarNota(6.9)
	fmt.Println("Informações do aluno após adicionar uma nota:")
	aluno.ImprimirInformacoes()

	aluno.RemoverNota(1)
	fmt.Println("Informações do aluno após remover uma nota:")
	aluno.ImprimirInformacoes()
}
