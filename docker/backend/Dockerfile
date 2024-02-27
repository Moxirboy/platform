# FROM ubuntu:latest

# # Install Bash
# RUN apt-get update && \
#     apt-get install -y bash

# # Set the default command to bash
# CMD ["bash"]
# # workspace (GOPATH) configured at /go

FROM golang:1.22.0-alpine3.19 as builder    

RUN apk add --no-cache bash
# RUN go install github.com/cespare/reflex@v0.3.1

# COPY reflex.conf /

# COPY start.sh /


ENTRYPOINT ["reflex", "-c", "/reflex.conf"]
