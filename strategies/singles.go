package strategies

/*
import (
	"fmt"

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

	if done, err := scanCage(s); err != nil {
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

		// Scan the cols
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

			if s.Answer[i*domain.BOARD_COLS+lastCol[j]] != 0 {
				continue
			}

			fmt.Printf(" * Solving hidden single (row): %dx%d (%d)\n", lastCol[j], i, j+1)

			if err := s.SolveCell(i, lastCol[j], int8(j+1)); err != nil {
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

		// Scan the rows
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

			fmt.Printf(" * Solving hidden single (col): %dx%d (%d)\n", i, lastRow[j], j+1)

			if err := s.SolveCell(lastRow[j], i, int8(j+1)); err != nil {
				return false, err
			}

			return true, nil
		}
	}

	return false, nil
}

func scanCage(s *domain.Sudoku) (bool, error) {
	rows := domain.BOARD_ROWS / 3
	cols := domain.BOARD_COLS / 3

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			count := make([]int, 9)
			lastRow := make([]int, 9)
			lastCol := make([]int, 9)

			for row := i * 3; row < i*3+3; row++ {
				for col := j * 3; col < j*3+3; col++ {
					for _, a := range s.Available[row*domain.BOARD_COLS+col] {
						count[a-1]++
						lastRow[a-1] = row
						lastCol[a-1] = col
					}
				}
			}

			// A hidden single in the cage has been found, solve it
			for j, c := range count {
				if c != 1 {
					continue
				}

				if s.Answer[lastRow[j]*domain.BOARD_COLS+lastCol[j]] != 0 {
					continue
				}

				fmt.Printf(" * Solving hidden single (cage): %dx%d (%d)\n", lastCol[j], lastRow[j], j+1)

				if err := s.SolveCell(lastRow[j], lastCol[j], int8(j+1)); err != nil {
					return false, err
				}

				return true, nil
			}
		}
	}

	return false, nil
}
*/
