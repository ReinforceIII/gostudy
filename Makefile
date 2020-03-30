TIME= $(shell date '+%y_%m_%d__%H_%M')

test_all: 
	go test -cover -coverprofile ./coverage/$(TIME).out ./...
	go tool cover -html=./coverage/$(TIME).out

build:
	go build -o main ./cmd/...