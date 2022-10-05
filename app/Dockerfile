FROM golang:1.18.4-alpine

RUN mkdir /build
ADD . /build
WORKDIR /build

RUN go build -o main cmd/main.go

EXPOSE 8080

CMD ["/build/main"]