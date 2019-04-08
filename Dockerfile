FROM golang

RUN go get github.com/gorilla/mux

WORKDIR /go/src/github.com/heaptracetechnology/microservice-gmail

ADD . /go/src/github.com/heaptracetechnology/microservice-gmail

RUN go install github.com/heaptracetechnology/microservice-gmail

ENTRYPOINT microservice-gmail

EXPOSE 3000