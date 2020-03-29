test: 
	@go test -v ./api/...

build:
	@go build ./command/main.go