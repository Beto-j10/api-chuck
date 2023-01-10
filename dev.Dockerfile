FROM golang:1.19.4-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# RUN go build -o main

EXPOSE 8080

# CMD ["./main"]

CMD ["air", "-c", ".air.toml"]