# workspace (GOPATH) configured at /go
FROM golang:1.22.0-alpine3.19 as builder    

RUN go install github.com/cespare/reflex@v0.3.1

COPY reflex.conf /

COPY start.sh /

ENTRYPOINT ["reflex", "-c", "/reflex.conf"]
