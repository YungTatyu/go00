BINARY_NAME	:= main.out
SRCS_DIRS	:= src
TEST_DIRS	:= test
SRCS			:= $(shell find ${SRCS_DIRS} -name '*.go')
TEST_SRCS	:= $(shell find ${TEST_DIRS} -name '*.go')
RM				:= rm -f

all:	build

build:
	go build -o ${BINARY_NAME} ${SRCS} 

test:
	go test -v ${TEST_SRCS}
 
run:	build
	./${BINARY_NAME}
 
clean:
	go clean
	${RM} ${BINARY_NAME}

re:	all	re

format:
	go fmt ${SRCS}

.PHONY: all clean re test format
