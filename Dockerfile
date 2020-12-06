FROM golang:1.15-alpine AS builder
RUN apk update && apk add git
WORKDIR /go/src/pitcher
COPY . .
RUN go get -v ./...
RUN go build -v ./cmd/pitcher
FROM alpine:3.12
COPY --from=builder /go/src/pitcher/pitcher .
CMD ["/pitcher"]