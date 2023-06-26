package main

import (
	"fmt"
)

type Filme struct {
	Titulo     string
	Diretor    string
	Ano        int
	Avaliacoes []int
}

func (f *Filme) AdicionarAvaliacao(avaliacao int) {
	f.Avaliacoes = append(f.Avaliacoes, avaliacao)
}

func (f *Filme) RemoverAvaliacao(index int) {
	if index >= 0 && index < len(f.Avaliacoes) {
		f.Avaliacoes = append(f.Avaliacoes[:index], f.Avaliacoes[index+1:]...)
	}
}

func (f Filme) CalcularMediaAvaliacoes() float64 {
	total := 0
	for _, avaliacao := range f.Avaliacoes {
		total += avaliacao
	}
	media := float64(total) / float64(len(f.Avaliacoes))
	return media
}

func (f Filme) ImprimirInformacoes() {
	fmt.Println("Título:", f.Titulo)
	fmt.Println("Diretor:", f.Diretor)
	fmt.Println("Ano:", f.Ano)
	fmt.Println("Média de Avaliações:", f.CalcularMediaAvaliacoes())
}

func main() {
	filme := Filme{
		Titulo:     "Inception",
		Diretor:    "Christopher Nolan",
		Ano:        2010,
		Avaliacoes: []int{9, 8, 10},
	}

	fmt.Println("Informações do filme:")
	filme.ImprimirInformacoes()

	filme.AdicionarAvaliacao(7)
	fmt.Println("Informações do filme após adicionar uma avaliação:")
	filme.ImprimirInformacoes()

	filme.RemoverAvaliacao(1)
	fmt.Println("Informações do filme após remover uma avaliação:")
	filme.ImprimirInformacoes()
}
