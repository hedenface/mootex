FROM golang:1.19-alpine3.15 AS builder

RUN mkdir -p src/github.com/hedenface/mootex

COPY go.mod   $GOPATH/src/github.com/hedenface/mootex/
COPY Makefile $GOPATH/src/github.com/hedenface/mootex/
COPY pkg/     $GOPATH/src/github.com/hedenface/mootex/pkg/
COPY src/     $GOPATH/src/github.com/hedenface/mootex/src/

WORKDIR $GOPATH/src/github.com/hedenface/mootex

RUN apk add --no-cache bash make \
    && make


FROM alpine:3.15

RUN apk add --no-cache bash curl

COPY --from=builder /go/src/github.com/hedenface/mootex/mootex /bin/mootex

ENTRYPOINT ["/bin/mootex"]
