package domain

import (
	"fmt"
	"os"
)

const (
	BOARD_COLS = 9
	BOARD_ROWS = 9
	BOARD_SIZE = BOARD_COLS * BOARD_ROWS
)

type Sudoku struct {
	Cages     []*Cage
	Answer    []int8
	Available [][]int8
}

func NewSudoku() *Sudoku {
	available := make([][]int8, BOARD_SIZE)
	for i, _ := range available {
		available[i] = []int8{1, 2, 3, 4, 5, 6, 7, 8, 9}
	}

	return &Sudoku{
		Answer:    make([]int8, BOARD_SIZE),
		Available: available,
	}
}

func (s *Sudoku) ReadFrom(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return NewError(err)
	}
	defer f.Close()

	if err := s.readCages(f); err != nil {
		return err
	}

	return nil
}

func (s *Sudoku) Print() {
	fmt.Print(" -----------------------------------------------------")
	fmt.Println("--------------------------------------------")
	fmt.Print("                                             AVAILABLE")
	fmt.Println("                                            ")

	for i := 0; i < BOARD_ROWS; i++ {
		if i%3 == 0 {
			fmt.Print(" -----------------------------------------------------")
			fmt.Println("--------------------------------------------")
		}

		for j := 0; j < BOARD_COLS; j++ {
			if j%3 == 0 {
				fmt.Print(" | ")
			} else {
				fmt.Print("-")
			}

			idx := i*BOARD_COLS + j
			var n int
			var k int8
			for ; k < 9; k++ {
				if k+1 == s.Available[idx][n] {
					fmt.Print(k + 1)

					if n+1 < len(s.Available[idx]) {
						n++
					}
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println(" |")
	}

	fmt.Print(" -----------------------------------------------------")
	fmt.Println("--------------------------------------------")
}

func (s *Sudoku) SetAvailable(row, col int8, av []int8) {
	s.Available[row*BOARD_COLS+col] = make([]int8, len(av))
	copy(s.Available[row*BOARD_COLS+col], av)
}

func (s *Sudoku) readCages(f *os.File) error {
	var ncages int
	fmt.Fscan(f, &ncages)

	s.Cages = make([]*Cage, ncages)

	for i := 0; i < ncages; i++ {
		var sum int8
		var ncells int
		fmt.Fscan(f, &sum, &ncells)

		cage := &Cage{
			Sum:   sum,
			Cells: make([]*Cell, ncells),
		}

		var row, col int8
		for j := 0; j < ncells; j++ {
			fmt.Fscanf(f, "%d,%d", &col, &row)
			cage.Cells[j] = &Cell{
				Row: row,
				Col: col,
			}
		}

		s.Cages[i] = cage
	}

	return nil
}
