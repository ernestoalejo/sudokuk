
#include <fstream>
#include <iostream>

#include "sudoku.h"


using std::cerr;
using std::cout;
using std::endl;
using std::ifstream;
using std::vector;


int main(int argc, char* argv[]) {
  if (argc != 2) {
    cerr << "usage: " << argv[0] << " [data_filename]" << endl << endl;
    return 1;
  }

  Sudoku sudoku;

  ifstream data(argv[1]);
  if (!data) {
    cerr << "cannot open the data file" << endl;
    return 2;
  }

  int cages;
  data >> cages;

  for (int i = 0; i < cages; i++) {
    int sum;
    int ncells;
    data >> sum >> ncells;

    vector<Cell> cells(ncells);

    for (int j = 0; j < ncells; j++) {
      int x, y;
      char c;
      data >> x >> c >> y;

      if (c != ',') {
        cerr << "coordinates without a colon between them" << endl;
        return 3;
      }

      Cell cell;
      cell.x = x;
      cell.y = y;
      cells[j] = cell;
    }

    sudoku.addCage(sum, cells);
  }

  data.close();

  sudoku.solve();
  cout << sudoku << endl;

  return 0;
}
