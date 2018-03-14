FROM golang:1.10-alpine
RUN apk update && apk add git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
CMD ["pitcher"]
