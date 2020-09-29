.PHONY: test
test:
	./bin/run-redis.sh
	@go test -v -short ./...
