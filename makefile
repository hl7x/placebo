test:
	go test -v ./...

failed:
	go test -v ./... | grep FAIL
