.PHONY: build
build:
	go build -o bin/clotas cmd/clotas/main.go

run:
	go run cmd/clotas/main.go

clean:
	go clean

install:
	go install
