package main

import (
	"flag"
	"fmt"

	"github.com/ernestokarim/sudokuk/domain"
	"github.com/ernestokarim/sudokuk/strategies"
)

var data = flag.String("data", "", "the data file containing the sudoku to solve")

type Strategy func(s *domain.Sudoku) (modified bool, err error)

func main() {
	flag.Parse()

	s := domain.NewSudoku()

	if err := s.ReadFrom(*data); err != nil {
		fmt.Println(err)
		return
	}

	if err := solve(s); err != nil {
		fmt.Println(err)
		return
	}
}

func solve(s *domain.Sudoku) error {
	// Always print the final sudoku before exiting
	defer s.Print()

	// Apply the cages combinations read from the data file
	if err := strategies.KillerCombinations(s); err != nil {
		return err
	}

	// List of strategies we can apply, ordered from simpler and quicker to
	// more complex and slower
	slist := []Strategy{
		strategies.HiddenSingles,
		strategies.NakedPairs,
	}

	for {
		// Check if the puzzle is solved
		if s.Solved() {
			return nil
		}

		// Check for solved squares
		if done, err := s.SolvedSquares(); err != nil {
			return err
		} else if done {
			continue
		}

		// Apply all the strategies to the sudoku
		// until one modifies the content
		i := 0
		for _, strat := range slist {
			i += 1
			if done, err := strat(s); err != nil {
				return err
			} else if done {
				i -= 1
				break
			}
		}

		// The puzzle can't be solved
		if i == len(slist) {
			break
		}
	}

	fmt.Println("ERROR: No more strategies to try. Puzzle unsolvable")
	return nil
	//return domain.NewErrorf("no more strategies to try. puzzle unsolvable")
}
