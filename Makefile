BIN_OUTPUT := nostdglobals

.PHONY: clean test build

default: clean test build

build: clean
	go build -o ${BIN_OUTPUT} ./cmd/nostdglobals/...

install: clean
	go install ./cmd/nostdglobals/...

clean:
	rm -rf nostdglobals

test: clean
	go test -v -cover ./...