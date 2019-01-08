include .env
export
MAKEFLAGS += --silent

build:
	go build
run:
	go run main.go
