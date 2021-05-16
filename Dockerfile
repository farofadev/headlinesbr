FROM golang:1.16.4-buster

RUN go install github.com/mitranim/gow@latest

WORKDIR /go/src/app

COPY . .

RUN go get

EXPOSE 8080

CMD ["go", "run", "main.go"]