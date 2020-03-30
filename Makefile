TIME= $(shell date '+%y_%m_%d__%H_%M')

daily_test: 
	go test -coverprofile ./coverage/$(TIME).out ./...
	go tool cover -html=./coverage/$(TIME).out

test:
	go test -v -cover ./...

build:
	go build -o main ./cmd/...