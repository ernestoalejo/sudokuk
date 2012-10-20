
#ifndef INCLUDE_SUDOKU_H_
#define INCLUDE_SUDOKU_H_


#include <iostream>
#include <stdint.h>
#include <vector>

#include "common.h"


struct Cell {
  int x, y;
};


struct Cage {
  int sum;
  std::vector<Cell> cells;
};


class Sudoku {
public:
  Sudoku();
  ~Sudoku();

  void addCage(int8_t sum, const std::vector<Cell>& cells);

  void solve();

  friend std::ostream& operator<<(std::ostream& os, const Sudoku& s);

private:
  DISALLOW_COPY_AND_ASSIGN(Sudoku);

  int8_t* answer;
  int8_t** available;
  int8_t* sizes;

  std::vector<Cage> cages;

  void allocate();
  void free();

  void clearRow(int col, int row, int8_t value);
  void clearCol(int col, int row, int8_t value);
  void clearCage(int cage, int8_t value);

  int getCage(int x, int y);

  void removeAvailable(int col, int row, int8_t value);

  void setSolvedCell(int x, int y, int8_t value);
  bool checkSolved();

  bool killerCombinations();
};


#endif  // INCLUDE_SUDOKU_H_
