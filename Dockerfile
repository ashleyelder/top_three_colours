FROM golang:1.12.0-alpine3.9
ADD . /go/src/myapp
WORKDIR /go/src/myapp
RUN go get myapp
RUN apk add git
RUN go install
ENTRYPOINT ["/go/bin/myapp"]
