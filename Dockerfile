FROM golang:1.16.7-alpine3.13

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o SuperListsAPI cmd/api/main.go
EXPOSE 8080
CMD ["./SuperListsAPI"]
