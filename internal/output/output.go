package output

import (
	"fmt"

	"github.com/damedelion/go_labyrinth/internal/utils"
)

func Print(points []utils.Point) {
	for _, point := range points {
		fmt.Println(point.Row, point.Col)
	}

	fmt.Println(".")
}
