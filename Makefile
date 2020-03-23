.PHONY: build test cov covline run

build:
	go build -o bin/kisakita-api app/api/main.go

test:
	go test -cover -coverprofile=coverage.out $$(go list ./... | grep -Ev "app|entities")

cov:
	go tool cover -func coverage.out

covline:
	go tool cover -html=coverage.out

run:
	go run app/api/main.go
