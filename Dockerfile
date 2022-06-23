# Go example
# (c) 2022

FROM golang:alpine

WORKDIR /app
ADD . /app

ENV GIN_MODE release
ENV PUBLISH_ADDRESS ":8013"

RUN cd /app/cmd/main && go build
RUN go clean -cache

ENTRYPOINT /app/cmd/main/main

EXPOSE 8013
