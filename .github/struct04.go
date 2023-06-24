package main

import (
	"fmt"
	"time"
)

type Musica struct {
	Titulo  string
	Artista string
	Duracao time.Duration
}

type Playlist struct {
	Nome    string
	Musicas []Musica
}

func imprimirPlaylist(p Playlist) {
	fmt.Println("Playlist:", p.Nome)

	tempoTotal := time.Duration(0)
	for _, musica := range p.Musicas {
		tempoTotal += musica.Duracao
		fmt.Println("Música:", musica.Titulo)
		fmt.Println("Artista:", musica.Artista)
		fmt.Println("Duração:", musica.Duracao)
		fmt.Println("---------------------------------")
	}

	fmt.Println("Tempo total da playlist:", tempoTotal)
}

func main() {
	musicas := []Musica{
		{Titulo: "Música 1", Artista: "Artista 1", Duracao: 3 * time.Minute},
		{Titulo: "Música 2", Artista: "Artista 2", Duracao: 4 * time.Minute},
		{Titulo: "Música 3", Artista: "Artista 3", Duracao: 5 * time.Minute},
	}

	playlist := Playlist{
		Nome:    "Minha Playlist",
		Musicas: musicas,
	}

	imprimirPlaylist(playlist)
}
