FROM golang:latest

RUN go get github.com/codegangsta/negroni
RUN go get github.com/gorilla/mux
RUN go get github.com/globalsign/mgo
RUN go get github.com/smartystreets/goconvey/convey
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/stretchr/testify/mock

ADD . /go/src/API_MVC
WORKDIR /go/src/API_MVC

RUN go build

ENTRYPOINT ./API_MVC

EXPOSE 8000
