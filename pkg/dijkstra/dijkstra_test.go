package dijkstra_test

import (
	"testing"

	"github.com/damedelion/go_labyrinth/internal/utils"
	"github.com/damedelion/go_labyrinth/pkg/dijkstra"
	"github.com/stretchr/testify/assert"
)

func TestDijkstra(t *testing.T) {
	tests := []struct {
		name      string
		labyrinth [][]int
		start     utils.Point
		finish    utils.Point
		expected  []utils.Point
		expectErr bool
	}{
		{
			name: "Example",
			labyrinth: [][]int{
				{1, 2, 0},
				{2, 0, 1},
				{9, 1, 0},
			},
			start:  utils.Point{Row: 0, Col: 0},
			finish: utils.Point{Row: 2, Col: 1},
			expected: []utils.Point{
				{Row: 0, Col: 0}, {Row: 1, Col: 0}, {Row: 2, Col: 0}, {Row: 2, Col: 1},
			},
			expectErr: false,
		},
		{
			name: "Basic Path",
			labyrinth: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			start:  utils.Point{Row: 0, Col: 0},
			finish: utils.Point{Row: 2, Col: 2},
			expected: []utils.Point{
				{Row: 0, Col: 0}, {Row: 1, Col: 0}, {Row: 2, Col: 0}, {Row: 2, Col: 1}, {Row: 2, Col: 2},
			},
			expectErr: false,
		},
		{
			name: "No Path",
			labyrinth: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
			start:     utils.Point{Row: 0, Col: 0},
			finish:    utils.Point{Row: 2, Col: 2},
			expected:  nil,
			expectErr: true,
		},
		{
			name: "Single Cell",
			labyrinth: [][]int{
				{1},
			},
			start:     utils.Point{Row: 0, Col: 0},
			finish:    utils.Point{Row: 0, Col: 0},
			expected:  []utils.Point{{Row: 0, Col: 0}},
			expectErr: false,
		},
		{
			name: "Large Labyrinth",
			labyrinth: [][]int{
				{1, 9, 8, 7, 6},
				{1, 0, 0, 0, 1},
				{1, 1, 1, 0, 1},
				{1, 0, 1, 1, 1},
				{1, 1, 1, 0, 1},
			},
			start:  utils.Point{Row: 0, Col: 0},
			finish: utils.Point{Row: 4, Col: 4},
			expected: []utils.Point{
				{Row: 0, Col: 0}, {Row: 1, Col: 0}, {Row: 2, Col: 0}, {Row: 2, Col: 1}, {Row: 2, Col: 2},
				{Row: 3, Col: 2}, {Row: 3, Col: 3}, {Row: 3, Col: 4}, {Row: 4, Col: 4},
			},
			expectErr: false,
		},
		{
			name: "Long Path",
			labyrinth: [][]int{
				{1, 9, 1},
				{1, 0, 1},
				{1, 0, 1},
				{1, 1, 1},
				{0, 1, 0},
			},
			start:  utils.Point{Row: 0, Col: 0},
			finish: utils.Point{Row: 0, Col: 2},
			expected: []utils.Point{
				{Row: 0, Col: 0}, {Row: 1, Col: 0}, {Row: 2, Col: 0}, {Row: 3, Col: 0}, {Row: 3, Col: 1},
				{Row: 3, Col: 2}, {Row: 2, Col: 2}, {Row: 1, Col: 2}, {Row: 0, Col: 2},
			},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := dijkstra.Dijkstra(tt.labyrinth, tt.start, tt.finish)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
