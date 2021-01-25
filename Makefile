.PHONY: run
run:
	go run main.go

.PHONY: build
build:
	go build -ldflags "-s -w" -o server main.go