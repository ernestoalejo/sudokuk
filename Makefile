#
# 'make depend' uses makedepend to automatically generate dependencies 
#               (dependencies are added to end of Makefile)
# 'make'        build executable file 'myCXX'
# 'make clean'  removes all .o and executable files
#

APP = sudokuk

BIN = bin
OBJ = obj
INCLUDE = include
SRC = src

TARGET = dev

CXX = g++
CFLAGS = -Wall -Wextra -Werror -Wno-unused-parameter -g
PROD_CFLAGS = -Wall -Wextra -Werror -O2
INCLUDES = -I$(INCLUDE)
LFLAGS = 
LIBS =
SRCS = $(shell find $(SRC) -name '*.cc')
OBJS = $(patsubst $(SRC)/%.cc,$(OBJ)/%.o, $(SRCS))
MAIN = $(addprefix $(BIN)/, $(APP))


ifeq ($(TARGET),prod)
	CFLAGS = $(PROD_CFLAGS)
endif


.PHONY: all depend clean lint


all: $(MAIN)


$(MAIN): $(OBJS) 
	$(CXX) $(CFLAGS) -o $(MAIN) $(OBJS) $(LFLAGS) $(LIBS)


obj/%.o: src/%.cc
	$(CXX) $(CFLAGS) $(INCLUDES) -c $< -o $@


clean:
	$(RM) bin/* obj/* make.depend


lint:
	./lint.sh


make.depend: $(SRCS) 
	$(CXX) $(INCLUDES) -M $^ > $@ 


include make.depend 
