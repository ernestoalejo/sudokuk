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
	AVAILABLE_1 = 1 << iota
	AVAILABLE_2
	AVAILABLE_3
	AVAILABLE_4
	AVAILABLE_5
	AVAILABLE_6
	AVAILABLE_7
	AVAILABLE_8
	AVAILABLE_9

	AVAILABLE_ALL = AVAILABLE_1 | AVAILABLE_2 | AVAILABLE_3 | AVAILABLE_4 |
		AVAILABLE_5 | AVAILABLE_6 | AVAILABLE_7 | AVAILABLE_8 | AVAILABLE_9
)

type Sudoku struct {
	Cages     []*Cage
	Answer    []int8
	Available []int
}

func NewSudoku() *Sudoku {
	available := make([]int, BOARD_SIZE)
	for i, _ := range available {
		available[i] = AVAILABLE_ALL
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

			var n uint
			var k int8
			for ; k < 9; k++ {
				mask := 1 << n
				if s.Available[idx]&mask == mask {
					fmt.Print(k + 1)
					n++
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

/*
func (s *Sudoku) SetAvailable(row, col int8, av []int8) {
	s.Available[row*BOARD_COLS+col] = make([]int8, len(av))
	copy(s.Available[row*BOARD_COLS+col], av)
}

func (s *Sudoku) RemoveAvailable(row, col int, value int8) error {
	idx := row*BOARD_COLS + col

	for i, n := range s.Available[idx] {
		if n == value {
			if len(s.Available[idx]) == 1 {
				return NewErrorf("cannot remove the last available number: %dx%d", row, col)
			}

			s.Available[idx] = append(s.Available[idx][:i], s.Available[idx][i+1:]...)
			return nil
		}
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
	for i, available := range s.Available {
		if s.Answer[i] == 0 && len(available) == 1 {
			fmt.Printf(" * Found solved square: %dx%d: %d\n", i/BOARD_COLS,
				i%BOARD_COLS, available[0])

			if err := s.SolveCell(i/BOARD_COLS, i%BOARD_COLS, available[0]); err != nil {
				return false, err
			}

			modified = true
		}
	}

	return modified, nil
}

func (s *Sudoku) SolveCell(row, col int, sol int8) error {
	idx := row*BOARD_COLS + col

	s.Answer[idx] = sol
	if len(s.Available[idx]) != 1 {
		s.Available[idx] = []int8{sol}
	}

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
*/
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
