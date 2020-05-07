.PHONY: dep build test cov covline dbcreate dbstatus dbmigrate run clean

dep:
	bundle install

build:
	go build -o bin/kisakita cmd/main.go

test:
	go test -cover -coverprofile=coverage.out $$(go list ./... | grep -Ev "cmd|bin|db|doc|entity|mocks")

cov:
	go tool cover -func coverage.out

covline:
	go tool cover -html=coverage.out

dbcreate:
	rake db:create

dbstatus:
	rake db:migrate:status

dbmigrate:
	rake db:migrate

run:
	go run cmd/main.go

clean:
	rm -rf bin
