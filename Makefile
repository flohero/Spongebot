export GO111MODULE=on
MAKEFLAGS += -B

all: clean website build

linux-release: linux website

linux-arm: clean website
	env GOOS=linux GOARCH=arm GOARM=5 go build

linux: clean website
	env GOOS=linux GOARCH=amd64 GOARM=5 go build

build:
	go build

run:
	go run main.go

website:
	cd ./website/; echo "I'm in some_dir"; \
	npm run-script "build go";

clean:
	rm -rf ./static/;
	rm ./Spongebot.*;
