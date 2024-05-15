build:
	@go build -o bin/api

run: build
	@./bin/api

test:
	@go test -v ./...

docker:
	@docker build -t api .
	@docker run -p 3000:3000 api