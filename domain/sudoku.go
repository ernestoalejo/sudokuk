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

const (
	AV_1 = 1 << iota
	AV_2
	AV_3
	AV_4
	AV_5
	AV_6
	AV_7
	AV_8
	AV_9

	AV_ALL = AV_1 | AV_2 | AV_3 | AV_4 | AV_5 | AV_6 | AV_7 | AV_8 | AV_9
)

var (
	AV_LIST = [9]uint{AV_1, AV_2, AV_3, AV_4, AV_5, AV_6, AV_7, AV_8, AV_9}
)

type Sudoku struct {
	Cages     []*Cage
	Answer    []int8
	Available []uint
}

func NewSudoku() *Sudoku {
	return &Sudoku{
		Answer:    make([]int8, BOARD_SIZE),
		Available: make([]uint, BOARD_SIZE),
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
	fmt.Println("                                             AVAILABLE")

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

			var k uint
			for ; k < 9; k++ {
				var mask uint = 1 << k
				if s.Available[idx]&mask == mask {
					fmt.Print(k + 1)
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println(" |")
	}

	fmt.Print(" -----------------------------------------------------")
	fmt.Println("--------------------------------------------")

	fmt.Println(" -------------------------")
	fmt.Println("          ANSWER")

	for i := 0; i < BOARD_ROWS; i++ {
		if i%3 == 0 {
			fmt.Println(" -------------------------")
		}

		for j := 0; j < BOARD_COLS; j++ {
			if j%3 == 0 {
				fmt.Print(" | ")
			} else {
				fmt.Print(" ")
			}

			n := s.Answer[i*BOARD_COLS+j]
			if n == 0 {
				fmt.Print(" ", "")
			} else {
				fmt.Print(n, "")
			}
		}
		fmt.Println(" |")
	}

	fmt.Println(" -------------------------")
}

func (s *Sudoku) RemoveAvailable(row, col int, value uint) error {
	idx := row*BOARD_COLS + col

	s.Available[idx] &^= 1 << (value - 1)

	if s.Available[idx] == 0 {
		return NewErrorf("cannot remove the last available number: %dx%d", col, row)
	}

	return nil
}

func (s *Sudoku) Solved() bool {
	for _, c := range s.Answer {
		if c == 0 {
			return false
		}
	}
	return true
}

func (s *Sudoku) SolvedSquares() (modified bool, e error) {
	for i, av := range s.Available {
		if s.Answer[i] == 0 && bitsSet(av) == 1 {
			fmt.Printf(" * Found solved square %d: %dx%d: %d\n", i, i/BOARD_COLS, i%BOARD_COLS, av)

			if err := s.SolveCell(i/BOARD_COLS, i%BOARD_COLS, bitSet(av)+1); err != nil {
				return false, err
			}

			modified = true
		}
	}

	return modified, nil
}

func (s *Sudoku) SolveCell(row, col int, sol uint) error {
	idx := row*BOARD_COLS + col

	s.Answer[idx] = int8(sol)
	s.Available[idx] = 1 << (sol - 1)

	// Clear the row
	for i := 0; i < BOARD_ROWS; i++ {
		if i == row {
			continue
		}

		if err := s.RemoveAvailable(i, col, sol); err != nil {
			return err
		}
	}

	// Clear the col
	for i := 0; i < BOARD_COLS; i++ {
		if i == col {
			continue
		}

		if err := s.RemoveAvailable(row, i, sol); err != nil {
			return err
		}
	}

	// Clear the cage
	var x int = col / 3
	x *= 3
	var y int = row / 3
	y *= 3

	for i := x; i < x+3; i++ {
		for j := y; j < y+3; j++ {
			if i == col && j == row {
				continue
			}

			if err := s.RemoveAvailable(j, i, sol); err != nil {
				return err
			}
		}
	}

	return nil
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
