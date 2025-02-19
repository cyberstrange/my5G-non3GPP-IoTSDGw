FROM golang:1.13-alpine

ENV PROJECT_PATH=/gw
ENV PATH=$PATH:$PROJECT_PATH/build
ENV CGO_ENABLED=0
ENV GO_EXTRA_BUILD_ARGS="-a -installsuffix cgo"

RUN apk add --no-cache ca-certificates tzdata make git bash

RUN mkdir -p $PROJECT_PATH
WORKDIR $PROJECT_PATH
COPY .  $PROJECT_PATH

#RUN make clean build


WORKDIR $PROJECT_PATH
