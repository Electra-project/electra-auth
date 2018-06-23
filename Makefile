NAME = electra-auth

ifeq ($(OS), Windows_NT)
	BINARY_NAME = ${NAME}.exe
else
	BINARY_NAME = ${NAME}
endif

install:
	go get -u golang.org/x/lint/golint
	go get -u github.com/tools/godep
	godep restore

lint:
	golint ./main.go
	golint ./src/...

start:
	go build && "./${BINARY_NAME}"

test:
	make lint
	go test -v ./...
