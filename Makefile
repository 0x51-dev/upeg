.PHONY: test gen fmt

test:
	go test -v -cover ./... --count=5

gen:
	go generate

fmt:
	go mod tidy
	gofmt -s -w .
	goarrange run -r .
	golangci-lint run ./...
