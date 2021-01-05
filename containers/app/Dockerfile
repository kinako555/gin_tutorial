FROM golang:1.14.2
COPY . .
WORKDIR /go/src/gin_tutorial
RUN go get github.com/gin-gonic/gin
RUN go get github.com/codegangsta/gin
ADD . /go/src/gin_tutorial