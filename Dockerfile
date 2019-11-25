FROM golang:1.12.0-alpine3.9
ADD . /go/src/myapp
RUN apk update; \
    apk upgrade; \
    apk add git;
WORKDIR /go/src/myapp
RUN go get myapp
RUN go install
ENTRYPOINT ["/go/bin/myapp"]