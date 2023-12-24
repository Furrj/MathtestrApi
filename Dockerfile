FROM golang:1.21.1

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

ADD cmd ./cmd
ADD internal ./internal
ADD client ./client

ENV MODE=PROD
ENV DB_URL=postgres://postgres:password@localhost:5432/mathtestr
ENV DB_URL_TEST=postgres://postgres:password@localhost:5432/test
ENV GIN_MODE=release

RUN go build -o server cmd/server.go

EXPOSE 5000

CMD ["./server"]

