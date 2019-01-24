export GO111MODULE=on
MAKEFLAGS += -B

local: website build

release: windows linux

linux-release: linux website

linux-arm:
	env GOOS=linux GOARCH=arm GOARM=5 go build -o Spongebot.linux.arm

linux:
	env GOOS=linux GOARCH=amd64 GOARM=5 go build -o Spongebot.linux

windows:
	env GOOS=windows GOARCH=amd64 go build -o Spongebot.windows.exe

build:
	go build -o ./out/Spongebot

run:
	go run main.go

website:
	cd ./website/; npm i;npm run-script "build release"

website-static:
	cd ./website/; npm i;npm run-script "build go"

clean:
	rm -rf ./static/ || true;
	rm ./Spongebot* || true;
	rm -r ./out || true;
