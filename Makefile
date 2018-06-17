NAME = electra-auth

ifeq ($(OS), Windows_NT)
	BINARY_NAME = ${NAME}.exe
else
	BINARY_NAME = ${NAME}
endif

download-binary:
	go run ./tasks/download_binary.go

install:
	godep restore
	make download-binary

start:
	go build && "./${BINARY_NAME}"

test:
	go test -v ./...
