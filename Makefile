run:
	@go run cmd/main.go

build:
	mkdir -p build
	@go build -o build/solver cmd/main.go

clean:
	rm -r build

test:
	@go test ./pkg/vigenere