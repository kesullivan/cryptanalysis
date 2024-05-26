run:
	@go run cmd/main.go

build:
	mkdir -p build
	@go build -o build/solver cmd/main.go

clean:
	if [ -d "build" ]; then rm -r build; fi

test:
	@go test ./pkg/vigenere