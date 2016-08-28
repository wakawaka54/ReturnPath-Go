FROM golang

ADD . /go/src/rpapi

RUN go get rpapi

RUN go install rpapi

ENV GO_HOME /go/src/rpapi/
ENV GO_PORT 1337

ENTRYPOINT /go/bin/rpapi

EXPOSE 1337
