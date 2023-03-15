test:
	ginkgo -r
lint:
	find . -type f -name "*.go" -exec gofmt -s -w {} \;
	find . -type f -name "*.go" -exec goimports -w {} \;
	golangci-lint run -c .golangci.yml
run:
	go run main.go
