FROM golang:1.7

ADD . /go/src/github.com/tehsis/rabbitscore
WORKDIR /go/src/github.com/tehsis/rabbitscore
RUN go get github.com/tools/godep
RUN godep restore
RUN go install
ENTRYPOINT /go/bin/rabbitscore

EXPOSE 8080
