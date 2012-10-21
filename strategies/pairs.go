package strategies

import (
	"fmt"

	"github.com/ernestokarim/sudokuk/domain"
)

func NakedPairs(s *domain.Sudoku) (bool, error) {
	if done, err := scanNakedPairsRows(s); err != nil {
		return false, err
	} else if done {
		return true, nil
	}

	if done, err := scanNakedPairsCols(s); err != nil {
		return false, err
	} else if done {
		return true, nil
	}

	return false, nil
}

func scanNakedPairsRows(s *domain.Sudoku) (bool, error) {
	modified := false

	for i := 0; i < domain.BOARD_ROWS; i++ {
		for j := 0; j < domain.BOARD_COLS; j++ {
			// Extract the available value
			v := s.Available[i*domain.BOARD_COLS+j]

			// Ignore if there aren't two unique available answers
			if domain.BitsSet(v) != 2 {
				continue
			}

			for scanX := j + 1; scanX < domain.BOARD_COLS; scanX++ {
				// Ignore the answer if it's not the pair of the currently
				// detected one
				if s.Available[i*domain.BOARD_COLS+scanX] != v {
					continue
				}

				s.Print()

				logged := false
				for applyX := 0; applyX < domain.BOARD_COLS; applyX++ {
					idx := i*domain.BOARD_COLS + applyX

					// Ignore the two answers containing the pair, as well
					// as cells where the strategy has no effect.
					if applyX == j || applyX == scanX || s.Available[idx]&v != v {
						continue
					}

					if !logged {
						modified = true
						logged = true
						fmt.Printf(" * Solving naked pair [row]: %dx%d + %dx%d\n", j, i, scanX, i)
					}

					s.Available[idx] &^= v
				}
			}
		}
	}

	return modified, nil
}

func scanNakedPairsCols(s *domain.Sudoku) (bool, error) {
	modified := false

	for i := 0; i < domain.BOARD_ROWS; i++ {
		for j := 0; j < domain.BOARD_COLS; j++ {
			// Extract the available value
			v := s.Available[i*domain.BOARD_COLS+j]

			// Ignore if there aren't two unique available answers
			if domain.BitsSet(v) != 2 {
				continue
			}

			for scanY := i + 1; scanY < domain.BOARD_ROWS; scanY++ {
				// Ignore the answer if it's not the pair of the currently
				// detected one
				if s.Available[scanY*domain.BOARD_COLS+j] != v {
					continue
				}

				logged := false
				for applyY := 0; applyY < domain.BOARD_ROWS; applyY++ {
					idx := applyY*domain.BOARD_COLS + j

					// Ignore the two answers containing the pair, as well
					// as cells where the strategy has no effect.
					if applyY == i || applyY == scanY || s.Available[idx]&v != v {
						continue
					}

					if !logged {
						modified = true
						logged = true
						fmt.Printf(" * Solving naked pair [col]: %dx%d + %dx%d\n", j, i, j, scanY)
					}

					s.Available[idx] &^= v
				}
			}
		}
	}

	return modified, nil
}
