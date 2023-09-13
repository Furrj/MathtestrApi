FROM golang:1.21.0

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

ADD cmd ./cmd
ADD build ./build

WORKDIR /app/cmd

EXPOSE 5000

CMD ["go", "run", "."]