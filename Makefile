export GO111MODULE=on
MAKEFLAGS += --silent

build:
	go build
run:
	go run main.go
