export GO111MODULE=on
MAKEFLAGS += -B

all: clean website build

release: windows linux website

linux-release: linux website

linux-arm:
	env GOOS=linux GOARCH=arm GOARM=5 go build

linux:
	env GOOS=linux GOARCH=amd64 GOARM=5 go build -o Spongebot.linux

windows:
	env GOOS=windows GOARCH=amd64 go build -o Spongebot.windows.exe

build:
	go build

run:
	go run main.go

website:
	cd ./website/; echo "I'm in some_dir"; \
	npm run-script "build go";

clean:
	rm -rf ./static/;
	rm ./Spongebot*;
