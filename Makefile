NAME = electra-auth

ifeq ($(OS), Windows_NT)
	BINARY_NAME = ${NAME}.exe
else
	BINARY_NAME = ${NAME}
endif

install:
	godep restore
	go get -u golang.org/x/lint/golint

lint:
	golint ./main.go
	golint ./src/...

start:
	go build && "./${BINARY_NAME}"

test:
	make lint
	go test -v ./...
