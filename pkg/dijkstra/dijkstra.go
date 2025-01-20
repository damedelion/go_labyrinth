package dijkstra

import (
	"container/heap"
	"fmt"
	"math"

	"github.com/damedelion/go_labyrinth/internal/utils"
	"github.com/damedelion/go_labyrinth/pkg/priorq"
)

func Dijkstra(labyrinth [][]int, start, finish utils.Point) ([]utils.Point, error) {
	rows, cols := len(labyrinth), len(labyrinth[0])
	directions := []utils.Point{
		{Row: -1, Col: 0}, {Row: 1, Col: 0}, {Row: 0, Col: -1}, {Row: 0, Col: 1},
	}

	distance := make([][]int, rows)
	for i := range distance {
		distance[i] = make([]int, cols)
		for j := range distance[i] {
			distance[i][j] = math.MaxInt
		}
	}

	distance[start.Row][start.Col] = 0
	pq := &priorq.PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &utils.Node{Point: start, Distance: 0})

	previous := make(map[utils.Point]*utils.Point)

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*utils.Node)

		if current.Point == finish {
			path := []utils.Point{}
			for p := &current.Point; p != nil; p = previous[*p] {
				path = append([]utils.Point{*p}, path...)
			}
			return path, nil
		}

		for _, d := range directions {
			next := utils.Point{Row: current.Row + d.Row, Col: current.Col + d.Col}
			if next.Row >= 0 && next.Row < rows && next.Col >= 0 && next.Col < cols && labyrinth[next.Row][next.Col] > 0 {
				newDist := current.Distance + labyrinth[next.Row][next.Col]
				if newDist < distance[next.Row][next.Col] {
					distance[next.Row][next.Col] = newDist
					heap.Push(pq, &utils.Node{Point: next, Distance: newDist})
					previous[next] = &current.Point
				}
			}
		}
	}

	return nil, fmt.Errorf("path not found")
}
