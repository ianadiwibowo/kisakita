.PHONY: build test cov covline run clean

build:
	go build -o bin/kisakita_stories_api app/api.go

test:
	go test -cover -coverprofile=coverage.out $$(go list ./... | grep -Ev "app|entities")

cov:
	go tool cover -func coverage.out

covline:
	go tool cover -html=coverage.out

run:
	go run app/api.go

clean:
	rm -rf bin
