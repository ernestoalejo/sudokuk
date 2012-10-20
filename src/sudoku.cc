

#include "sudoku.h"


using std::cout;
using std::endl;
using std::ostream;
using std::vector;


inline int index(int8_t* array, int row, int col) {
  return static_cast<int>(array[row*BOARD_COLS + col]);
}


inline void setindex(int8_t* array, int row, int col, int8_t value) {
  array[row*BOARD_COLS + col] = value;
}


inline int8_t* index(int8_t** array, int row, int col) {
  return array[row*BOARD_COLS + col];
}


Sudoku::Sudoku() {
  allocate();
}


Sudoku::~Sudoku() {
  free();
}


void Sudoku::allocate() {
  answer = new int8_t[BOARD_SIZE];
  for (int i = 0; i < BOARD_SIZE; i++) {
    answer[i] = 0;
  }

  sizes = new int8_t[BOARD_SIZE];
  for (int i = 0; i < BOARD_SIZE; i++) {
    sizes[i] = 9;
  }

  available = new int8_t*[BOARD_SIZE];
  for (int i = 0; i < BOARD_SIZE; i++) {
    available[i] = new int8_t[9];

    for (int j = 0; j < 9; j++) {
      available[i][j] = (int8_t) j+1;
    }
  }
}


void Sudoku::free() {
  delete [] answer;
  delete [] sizes;

  for (int i = 0; i < BOARD_SIZE; i++) {
    if (available[i] != NULL) {
      delete [] available[i];
    }
  }
  delete [] available;
}


ostream& operator<<(ostream& os, const Sudoku& s) {
  // Print the available values
  cout << " ----------------------------------------------------------";
  cout << "---------------------------------------" << endl;
  cout << "                                             AVAILABLE     ";
  cout << "                                       " << endl;
  for (int i = 0; i < BOARD_ROWS; i++) {
    if (i % 3 == 0) {
      cout << " ----------------------------------------------------------";
      cout << "---------------------------------------" << endl;
    }

    for (int j = 0; j < BOARD_COLS; j++) {
      if (j%3 == 0)
        cout << " | ";
      else
        cout << "-";

      int idx = 0;
      for (int k = 0; k < 9; k++) {
        if (k+1 == index(s.available, i, j)[idx]) {
          cout << k+1;

          if (idx < index(s.sizes, i, j)-1)
            idx++;
        } else {
          cout << " ";
        }
      }
    }

    cout << " |" << endl;
  }
  cout << " ----------------------------------------------------------";
  cout << "---------------------------------------" << endl;

  return os;

  // Print the currently answered values
  cout << " -------------------------" << endl;
  cout << "           ANSWER         " << endl;
  for (int i = 0; i < BOARD_ROWS; i++) {
    if (i % 3 == 0)
      cout << " -------------------------" << endl;

    for (int j = 0; j < BOARD_COLS; j++) {
      if (j%3 == 0)
        cout << " |";

      cout << " " << index(s.answer, i, j);
    }

    cout << " |" << endl;
  }
  cout << " -------------------------" << endl;

  return os;
}


void Sudoku::addCage(int8_t sum, const vector<Cell>& cells) {
  Cage cage;
  cage.sum = sum;
  cage.cells = cells;
  cages.push_back(cage);
}


void Sudoku::setSolvedCell(int x, int y, int8_t value) {
  setindex(answer, y, x, value);
  clearCol(x, y, value);
  clearRow(x, y, value);
  clearCage(x, y, value);

  int idx = y*BOARD_COLS + x;
  delete [] available[idx];
  available[idx] = new int8_t[1];
  available[idx][0] = value;
  sizes[idx] = 1;
}


void Sudoku::clearCol(int col, int row, int8_t value) {
  for (int i = 0; i < BOARD_ROWS; i++) {
    removeAvailable(col, i, value);
  }
}


void Sudoku::clearRow(int col, int row, int8_t value) {
  for (int i = 0; i < BOARD_COLS; i++) {
    removeAvailable(i, row, value);
  }
}


void Sudoku::clearCage(int col, int row, int8_t value) {
  int x = col/3;
  x *= 3;

  int y = row/3;
  y *= 3;

  for (int i = y; i < y+3; i++) {
    for (int j = x; j < x+3; j++) {
      removeAvailable(j, i, value);
    }
  }
}


void Sudoku::removeAvailable(int col, int row, int8_t value) {
  int size = index(sizes, row, col);
  for (int i = 0; i < size; i++) {
    if (index(available, row, col)[i] == value) {
      int8_t* modified = new int8_t[size-1];
      int idx = row*BOARD_COLS + col;

      for (int j = 0; j < size; j++) {
        if (j < i)
          modified[j] = available[idx][j];
        else if (j > i)
          modified[j-1] = available[idx][j];
      }

      delete [] available[idx];
      available[idx] = modified;
      sizes[idx] = size-1;
      return;
    }
  }

  if (size == 1) {
    setSolvedCell(col, row, value);
  }
}


void Sudoku::setAvailable(int col, int row, const vector<int>& values) {
  int idx = row*BOARD_COLS + col;

  delete [] available[idx];
  sizes[idx] = values.size();

  int8_t* modified = new int8_t[sizes[idx]];
  available[idx] = modified;

  for (int i = 0; i < sizes[idx]; i++) {
    available[idx][i] = values[i];
  }
}


bool Sudoku::checkSolved() {
  for (int i = 0; i < BOARD_SIZE; i++) {
    if (answer[i] == 0)
      return false;
  }
  return true;
}


void Sudoku::solve() {
  killerCombinations();

  return;

  while (1) {
    if (checkSolved()) break;

    cout << "Cannot solve this puzzle! We run out of strategies!" << endl;
    break;
  }
}
