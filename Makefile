.PHONY: dep build test cov covline dbcreate dbstatus dbmigrate run clean

dep:
	bundle install

build:
	go build -o bin/kisakita_stories_api app/api.go

test:
	go test -cover -coverprofile=coverage.out $$(go list ./... | grep -Ev "app|bin|db|doc|entity")

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
	go run app/api.go

clean:
	rm -rf bin
