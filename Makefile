all: test build
build:
	go build -o beaver github.com/andremedeiros/beaver/cmd/beaver
test:
	go test -v ./...
clean:
	rm -f beaver
deps:
	go mod download


