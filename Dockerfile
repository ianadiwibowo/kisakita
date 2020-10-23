FROM golang:1.15-alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o bin/kisakita cmd/main.go

EXPOSE 8080

CMD ["./bin/kisakita"]
