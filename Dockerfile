FROM golang:1.10-alpine AS builder

ARG VERSION
ARG REVISION
ARG PACKAGE=github.com/nomkhonwaan/myblog-server

ENV CGO_ENABLED=0
ENV PACKAGE=${PACKAGE}

WORKDIR "${GOPATH}/src/${PACKAGE}"

COPY . .

RUN apk add -u --no-cache curl git && \
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && \
    dep ensure && \
    dep ensure -add github.com/jteeuwen/go-bindata && \
    go build \
      -ldflags " \
        -X ${PACKAGE}/cmd/myblog/app.version=${VERSION} \
        -X ${PACKAGE}/cmd/myblog/app.revision=${REVISION}" \
      -o /usr/local/bin/myblog-server \
      cmd/myblog/main.go

FROM alpine:latest

WORKDIR /usr/local/bin

COPY --from=builder /usr/local/bin/myblog-server .

EXPOSE 8080 

ENTRYPOINT [ "/usr/local/bin/myblog-server" ]