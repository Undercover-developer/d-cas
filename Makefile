build:
	@go build -o bin/ipfs

run: build
	@./bin/ipfs

test:
	@go test -v ./...