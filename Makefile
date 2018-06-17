NAME = electra-auth

ifeq ($(OS), Windows_NT)
	BINARY_NAME = ${NAME}.exe
else
	BINARY_NAME = ${NAME}
endif

install:
	godep restore

start:
	go build && "./${BINARY_NAME}"

test:
	go test -v ./...
