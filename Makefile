.PHONY: test
test:
	./run-redis.sh
	@go test -v -short ./...
