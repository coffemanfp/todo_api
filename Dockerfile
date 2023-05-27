FROM golang:1.20.0

WORKDIR /usr/src/app


COPY . .

RUN go mod tidy