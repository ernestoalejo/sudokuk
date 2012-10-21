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
	if err := strategies.KillerCombinations(s); err != nil {
		return err
	}

	s.Print()
	return nil
}
