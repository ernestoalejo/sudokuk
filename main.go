package main

import (
	"flag"
	"fmt"

	"github.com/ernestokarim/sudokuk/domain"
	"github.com/ernestokarim/sudokuk/strategies"
)

var data = flag.String("data", "", "the data file containing the sudoku to solve")

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
	defer s.Print()

	if err := strategies.KillerCombinations(s); err != nil {
		return err
	}

	for {
		if s.Solved() {
			break
		}

		if done, err := s.SolvedSquares(); err != nil {
			return err
		} else if done {
			continue
		}

		fmt.Println("No more strategies to do! Puzzle can't be resolved!")
		break
	}

	return nil
}
