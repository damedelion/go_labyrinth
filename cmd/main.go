package main

import (
	"log"

	"github.com/damedelion/go_labyrinth/internal/input"
	"github.com/damedelion/go_labyrinth/internal/output"
	"github.com/damedelion/go_labyrinth/pkg/dijkstra"
)

func main() {
	labyrinth, start, finish, err := input.Read()
	if err != nil {
		log.Fatal(err.Error())
	}

	path, err := dijkstra.Dijkstra(labyrinth, start, finish)
	if err != nil {
		log.Fatal(err.Error())
	}

	output.Print(path)
}
