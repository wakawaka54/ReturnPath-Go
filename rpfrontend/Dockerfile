FROM golang

ADD . /go/src/rpfrontend

RUN go get rpfrontend

RUN go install rpfrontend

ENV GO_HOME /go/src/rpfrontend/
ENV GO_PORT 1400

ENTRYPOINT /go/bin/rpfrontend

EXPOSE 1400
