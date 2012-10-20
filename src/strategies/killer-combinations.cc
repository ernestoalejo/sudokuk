
#include "sudoku.h"


using std::cout;
using std::endl;
using std::vector;


vector<int> getCombinations1(int sum) {
  std::vector<int> v;
  v.push_back(sum);
  return v;
}


vector<int> getCombinations2(int sum) {
  vector<int> v;

  for (int i = 1; i < 9; i++) {
    for (int j = 9; j > 0 && j > i; j--) {
      if (j+i == sum) {
        v.push_back(i);
        v.push_back(j);
      }
    }
  }

  return v;
}


vector<int> getCombinations3(int sum) {
  vector<int> v;

  for (int i = 1; i < 8; i++) {
    for (int j = 9; j > 0 && j > i; j--) {
      for (int k = 9; k > 0 && k > j; k--) {
        if (j+i+k == sum) {
          v.push_back(i);
          v.push_back(j);
          v.push_back(k);
        }
      }
    }
  }

  return v;
}


vector<int> getCombinations(int n, int sum) {
  switch (n) {
  case 1:
    return getCombinations1(sum);

  case 2:
    return getCombinations2(sum);

  case 3:
    return getCombinations3(sum);

  default:
    cout << "ERROR: searching combinations of " << n << " to sum up "
        << sum << endl;
  }

  throw "shouldn't reach here";
}


void Sudoku::killerCombinations() {
  for (unsigned int i = 0; i < cages.size(); i++) {
    vector<int> v = getCombinations(cages[i].cells.size(), cages[i].sum);

    cout << cages[i].cells.size() << " " << cages[i].sum << " ";
    for (unsigned int c = 0; c < v.size(); c++) {
      cout << v[c] << " ";
    }
    cout << endl;

    for (unsigned int j = 0; j < cages[i].cells.size(); j++) {
      setAvailable(cages[i].cells[j].x, cages[i].cells[j].y, v);
    }
  }
}

