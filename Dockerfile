FROM golang:1.12-alpine

# Prepare the app directory
WORKDIR /go/src/github.com/milesalex/tracer-demo/

# Ensure dependancies are installed
RUN apk add --update git && rm -rf /var/cache/apk/*
RUN go get github.com/golang/dep/cmd/dep

# Prepare dev build
ADD . .
RUN dep ensure --vendor-only
RUN go build -o main
CMD ["./main"]
