BINARY_NAME	:= main.out
SRCS_DIRS	:= .
SRCS			:= $(shell find ${SRCS_DIRS} -name '*.go')
RM				:= rm -f

all:	build

build:
	go build -o ${BINARY_NAME} ${SRCS} 

test:
	go test -v ./... 
 
run:	build
	./${BINARY_NAME}
 
clean:
	go clean
	${RM} ${BINARY_NAME}

re:	clean	all

format:
	go fmt ${SRCS} 

.PHONY: all clean re test format
