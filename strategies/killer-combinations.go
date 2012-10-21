package strategies

import (
	"github.com/ernestokarim/sudokuk/domain"
)

var combinations = map[int]map[int8]uint{
	1: map[int8]uint{
		1: AV_1,
		2: AV_2,
		3: AV_3,
		4: AV_4,
		5: AV_5,
		6: AV_6,
		7: AV_7,
		8: AV_8,
		9: AV_9,
	},
	2: map[int8]uint{
		3:  AV_1 | AV_2,
		4:  AV_1 | AV_3,
		5:  AV_1 | AV_2 | AV_3 | AV_4,
		6:  AV_1 | AV_2 | AV_4 | AV_5,
		7:  AV_1 | AV_2 | AV_3 | AV_4 | AV_5 | AV_6,
		8:  AV_1 | AV_2 | AV_3 | AV_5 | AV_6 | AV_7,
		9:  AV_1 | AV_2 | AV_3 | AV_4 | AV_5 | AV_6 | AV_7 | AV_8,
		10: AV_1 | AV_2 | AV_3 | AV_4 | AV_6 | AV_7 | AV_8 | AV_9,
		11: AV_2 | AV_3 | AV_4 | AV_5 | AV_6 | AV_7 | AV_8 | AV_9,
		12: AV_3 | AV_4 | AV_5 | AV_7 | AV_8 | AV_9,
		13: AV_4 | AV_5 | AV_6 | AV_7 | AV_8 | AV_9,
		14: AV_5 | AV_6 | AV_8 | AV_9,
		15: AV_6 | AV_7 | AV_8 | AV_9,
		16: AV_7 | AV_9,
		17: AV_8 | AV_9,
	},
	3: map[int8]uint{
		6:  AV_1 | AV_2 | AV_3,
		7:  AV_1 | AV_2 | AV_4,
		8:  AV_1 | AV_2 | AV_3 | AV_4 | AV_5,
		9:  AV_1 | AV_2 | AV_3 | AV_4 | AV_5 | AV_6,
		10: AV_1 | AV_2 | AV_3 | AV_4 | AV_5 | AV_6 | AV_7,
		11: AV_1 | AV_2 | AV_3 | AV_4 | AV_5 | AV_6 | AV_7 | AV_8,
		12: AV_1 | AV_2 | AV_3 | AV_4 | AV_5 | AV_6 | AV_7 | AV_8 | AV_9,
		13: AV_1 | AV_2 | AV_3 | AV_4 | AV_5 | AV_6 | AV_7 | AV_8 | AV_9,
		14: AV_1 | AV_2 | AV_3 | AV_4 | AV_5 | AV_6 | AV_7 | AV_8 | AV_9,
		15: AV_1 | AV_2 | AV_3 | AV_4 | AV_5 | AV_6 | AV_7 | AV_8 | AV_9,
		16: AV_1 | AV_2 | AV_3 | AV_4 | AV_5 | AV_6 | AV_7 | AV_8 | AV_9,
		17: AV_1 | AV_2 | AV_3 | AV_4 | AV_5 | AV_6 | AV_7 | AV_8 | AV_9,
		18: AV_1 | AV_2 | AV_3 | AV_4 | AV_5 | AV_6 | AV_7 | AV_8 | AV_9,
		19: AV_2 | AV_3 | AV_4 | AV_5 | AV_6 | AV_7 | AV_8 | AV_9,
		20: AV_3 | AV_4 | AV_5 | AV_6 | AV_7 | AV_8 | AV_9,
		21: AV_4 | AV_5 | AV_6 | AV_7 | AV_8 | AV_9,
		22: AV_5 | AV_6 | AV_7 | AV_8 | AV_9,
		23: AV_6 | AV_8 | AV_9,
		24: AV_7 | AV_8 | AV_9,
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
			s.Available[cell.Row*domain.BOARD_ROWS+cell.Col] |= comb
		}
	}

	return nil
}
