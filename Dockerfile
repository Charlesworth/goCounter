FROM alpine:3.2
MAINTAINER Charles Cochrane <cochrane26@gmail.com>

ENV GOROOT=/usr/lib/go \
    GOPATH=/gopath \
    GOBIN=/gopath/bin \
    PATH=$PATH:$GOROOT/bin:$GOPATH/bin

WORKDIR /gopath/src/app
ADD . /gopath/src/app
RUN mkdir data

RUN apk add -U git go && \
  go get -v app && \
  apk del git go && \
  rm -rf /gopath/pkg && \
  rm -rf /gopath/src && \
  rm -rf /var/cache/apk/*

ENTRYPOINT ["/gopath/bin/app"]
