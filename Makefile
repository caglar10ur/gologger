all:
	@go build gologger.go
	@go test
format:
	@gofmt -s -w *.go
vet:
	`which go` vet .
