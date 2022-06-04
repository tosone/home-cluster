build:
	go build -o bin/hc -ldflags "-s -w -X main.version=${VERSION}" ./cmd
