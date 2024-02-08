FROM golang:1.16.3-alpine3.13 as builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod init platform
RUN go mod tidy
RUN go clean --modcache
RUN go build -o main .
CMD ["/app/main"]