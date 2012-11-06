all:
	@go build gologger.go
	@go test
	@rm -f foo
