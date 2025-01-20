package input

import (
	"fmt"

	"github.com/damedelion/go_labyrinth/internal/utils"
)

func Read() ([][]int, utils.Point, utils.Point, error) {
	var m, n int
	_, err := fmt.Scan(&m, &n)
	if m <= 0 || n <= 0 || err != nil {
		return nil, utils.Point{}, utils.Point{}, fmt.Errorf("incorrect input")
	}

	labyrinth := make([][]int, m)
	for i := 0; i < m; i++ {
		labyrinth[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			_, err := fmt.Scan(&labyrinth[i][j])
			if err != nil {
				return nil, utils.Point{}, utils.Point{}, fmt.Errorf("incorrect input")
			}
		}
	}

	var rowStart, colStart, rowFinish, colFinish int
	fmt.Scan(&rowStart, &colStart, &rowFinish, &colFinish)
	start := utils.Point{Row: rowStart, Col: colStart}
	finish := utils.Point{Row: rowFinish, Col: colFinish}

	fmt.Println()

	return labyrinth, start, finish, nil
}
