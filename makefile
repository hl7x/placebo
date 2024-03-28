test:
	go test -v ./...

failed:
	go test -v ./... | grep FAIL
build:
	cd cmd/placebo/ && go build .
