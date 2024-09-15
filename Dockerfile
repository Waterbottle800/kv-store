FROM golang:1.23.1

WORKDIR /kv_project

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o main sample.go

EXPOSE 8080

CMD ["./main"]
