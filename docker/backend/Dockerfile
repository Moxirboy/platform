FROM golang:1.22.0-alpine3.19 as builder

RUN apk --no-cache add bash

RUN go install github.com/cespare/reflex@v0.3.1

COPY reflex.conf /
COPY start.sh /
RUN chmod 777 start.sh

ENTRYPOINT ["/bin/bash", "-c", "/start.sh"]
