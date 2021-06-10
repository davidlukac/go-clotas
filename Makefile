.PHONY: build
build:
	go build -o bin/clotas main.go

run:
	go run main.go

clean:
	go clean

install:
	go install
