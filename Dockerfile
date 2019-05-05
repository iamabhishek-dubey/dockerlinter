FROM golang:latest as builder
MAINTAINER Abhishek Dubey
COPY ./ /go/src/github.com/iamabhishek-dubey/dockerlinter/
WORKDIR /go/src/github.com/iamabhishek-dubey/dockerlinter/
RUN go get -v -t -d ./... \
    && go build

FROM alpine:latest
MAINTAINER Abhishek Dubey
RUN apk add --no-cache libc6-compat bash
COPY --from=builder /go/src/github.com/iamabhishek-dubey/dockerlinter/dockerlinter /bin/
COPY --from=builder /go/src/github.com/iamabhishek-dubey/dockerlinter/reports /dockerlinter/reports/
WORKDIR /dockerlinter/
