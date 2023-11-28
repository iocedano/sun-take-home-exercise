init:
	yarn --cwd ./client install

build:
	yarn --cwd ./client build

run: 
	make build
	go run main.go

test-go:
	go test -v ./...