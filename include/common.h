
#ifndef INCLUDE_COMMON_H_
#define INCLUDE_COMMON_H_


#define BOARD_COLS 9
#define BOARD_ROWS 9

#define BOARD_SIZE BOARD_COLS*BOARD_ROWS


// A macro to disallow the copy constructor and operator= functions
// This should be used in the private: declarations for a class
#define DISALLOW_COPY_AND_ASSIGN(TypeName) \
  TypeName(const TypeName&);               \
  void operator=(const TypeName&)


#endif  // INCLUDE_COMMON_H_
