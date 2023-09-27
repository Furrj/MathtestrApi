FROM golang:1.21.1

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod downloadg

ADD cmd ./cmd
ADD build ./build

RUN go build -o server server.go

EXPOSE 5000

CMD ["./server"]