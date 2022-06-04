build:
	go build -o bin/hc -ldflags "-s -w" ./cmd

image:
	docker build -t hc:latest -f build/Dockerfile .
