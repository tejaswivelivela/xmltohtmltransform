FROM golang:alpine
RUN apk update
RUN apk add git
RUN apk add build-base libxml2-dev
COPY . /go/src/xmltohtmltransform
WORKDIR /go/src/xmltohtmltransform
RUN go get -t -v ./...
RUN go build -o app;
RUN apk del git
CMD ["./app"]
