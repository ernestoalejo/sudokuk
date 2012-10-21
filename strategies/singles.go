package strategies

import (
	"github.com/ernestokarim/sudokuk/domain"
)

// Sometimes a number it's present once in its row, col or cage. Discard the
// rest of answers because that one can be there only.
func HiddenSingles(s *domain.Sudoku) (bool, error) {
	if done, err := scanRow(s); err != nil {
		return false, nil
	} else if done {
		return true, nil
	}

	if done, err := scanCol(s); err != nil {
		return false, nil
	} else if done {
		return true, nil
	}

	return false, nil
}

func scanRow(s *domain.Sudoku) (bool, error) {
	for i := 0; i < domain.BOARD_ROWS; i++ {
		count := make([]int, 9)
		lastCol := make([]int, 9)

		// Scan the rows
		for j := 0; j < domain.BOARD_COLS; j++ {
			for _, a := range s.Available[i*domain.BOARD_COLS+j] {
				count[a-1]++
				lastCol[a-1] = j

			}
		}

		// A hidden single in the row has been found, solve it
		for j, c := range count {
			if c != 1 {
				continue
			}

			if s.Answer[j*domain.BOARD_COLS+lastCol[j]] != 0 {
				continue
			}

			if err := s.SolveCell(i, lastCol[j], int8(c)); err != nil {
				return false, err
			}

			return true, nil
		}
	}

	return false, nil
}

func scanCol(s *domain.Sudoku) (bool, error) {
	for i := 0; i < domain.BOARD_COLS; i++ {
		count := make([]int, 9)
		lastRow := make([]int, 9)

		// Scan the cols
		for j := 0; j < domain.BOARD_ROWS; j++ {
			for _, a := range s.Available[j*domain.BOARD_COLS+i] {
				count[a-1]++
				lastRow[a-1] = j
			}
		}

		// A hidden single in the col has been found, solve it
		for j, c := range count {
			if c != 1 {
				continue
			}

			if s.Answer[lastRow[j]*domain.BOARD_COLS+i] != 0 {
				continue
			}

			if err := s.SolveCell(lastRow[j], i, int8(c)); err != nil {
				return false, err
			}

			return true, nil
		}
	}

	return false, nil
}
