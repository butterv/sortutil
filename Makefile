init:
	GO111MODULE=on go mod download

build:
	go build

test:
	go test -v

benchmark:
	go test -bench . -benchmem

lint:
	GO111MODULE=on golint ./...
