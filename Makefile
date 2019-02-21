export GO111MODULE=on
BINARY_NAME=drlm-cli

all: deps build
install:
	go install drlm-cli.go
build:
	go build drlm-cli.go
test:
	go test -v ./...
clean:
	go clean
	rm -f $(BINARY_NAME)
deps:
	go build -v ./...
upgrade:
	go get -u