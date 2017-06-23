FROM golang:1.8

WORKDIR /go/src/github.com/tehsis/rabbitscore
ADD . /go/src/github.com/tehsis/rabbitscore
RUN go get github.com/tools/godep
RUN godep restore

EXPOSE 8080
