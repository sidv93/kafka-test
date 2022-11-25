FROM golang:1.18.4-alpine3.16

RUN mkdir /app

ADD . /app

WORKDIR /app/app

RUN go mod tidy

RUN go build -o main .

CMD ["/app/app/main"]
