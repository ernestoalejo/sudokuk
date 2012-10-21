package strategies

import (
	"github.com/ernestokarim/sudokuk/domain"
)

var combinations = map[int]map[int8][]int8{
	1: map[int8][]int8{
		1: []int8{1},
		2: []int8{2},
		3: []int8{3},
		4: []int8{4},
		5: []int8{5},
		6: []int8{6},
		7: []int8{7},
		8: []int8{8},
		9: []int8{9},
	},
	2: map[int8][]int8{
		3:  []int8{1, 2},
		4:  []int8{1, 3},
		5:  []int8{1, 2, 3, 4},
		6:  []int8{1, 2, 4, 5},
		7:  []int8{1, 2, 3, 4, 5, 6},
		8:  []int8{1, 2, 3, 5, 6, 7},
		9:  []int8{1, 2, 3, 4, 5, 6, 7, 8},
		10: []int8{1, 2, 3, 4, 6, 7, 8, 9},
		11: []int8{2, 3, 4, 5, 6, 7, 8, 9},
		12: []int8{3, 4, 5, 7, 8, 9},
		13: []int8{4, 5, 6, 7, 8, 9},
		14: []int8{5, 6, 8, 9},
		15: []int8{6, 7, 8, 9},
		16: []int8{7, 9},
		17: []int8{8, 9},
	},
	3: map[int8][]int8{
		6:  []int8{1, 2, 3},
		7:  []int8{1, 2, 4},
		8:  []int8{1, 2, 3, 4, 5},
		9:  []int8{1, 2, 3, 4, 5, 6},
		10: []int8{1, 2, 3, 4, 5, 6, 7},
		11: []int8{1, 2, 3, 4, 5, 6, 7, 8},
		12: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9},
		13: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9},
		14: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9},
		15: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9},
		16: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9},
		17: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9},
		18: []int8{1, 2, 3, 4, 5, 6, 7, 8, 9},
		19: []int8{2, 3, 4, 5, 6, 7, 8, 9},
		20: []int8{3, 4, 5, 6, 7, 8, 9},
		21: []int8{4, 5, 6, 7, 8, 9},
		22: []int8{5, 6, 7, 8, 9},
		23: []int8{6, 8, 9},
		24: []int8{7, 8, 9},
	},
}

// Only some combinations are possible with a given (sum,no of cells) pair.
// With that fact, remove all the other available answers.
func KillerCombinations(s *domain.Sudoku) error {
	for _, cage := range s.Cages {
		m, ok := combinations[len(cage.Cells)]
		if !ok {
			return domain.NewErrorf("needed combinations for %d cells", len(cage.Cells))
		}

		comb, ok := m[cage.Sum]
		if !ok {
			return domain.NewErrorf("needed combinations for %d sum", cage.Sum)
		}

		for _, cell := range cage.Cells {
			s.SetAvailable(cell.Row, cell.Col, comb)
		}
	}

	return nil
}
