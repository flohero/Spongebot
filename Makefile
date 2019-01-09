include .env
include .env.sample
export
MAKEFLAGS += --silent

build:
	go build
run:
	go run main.go
