FROM golang:1.21.1

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

ADD cmd ./cmd
ADD internal ./internal
ADD build ./build

RUN go build -o server cmd/server.go

EXPOSE 5000

CMD ["./server"]
