all: build

GO111MODULE=on

build:
	go build -o bin/service/live/resource/cmd app/service/live/resource/cmd/main.go
	go build -o bin/service/main/music/cmd app/service/main/music/cmd/main.go
	go build -o bin/interface/main/user/cmd app/interface/main/user/cmd/main.go
	
test:
	go test --race -timeout 2m ./...
	
clean:
	go clean -i ./...
	@rm -rf ./bin

.PHONY: build test clean