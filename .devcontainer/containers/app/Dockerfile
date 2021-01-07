FROM golang:1.14
WORKDIR /go/src/gin_tutorial
COPY . .
RUN go get -u github.com/gin-gonic/gin
RUN go get github.com/codegangsta/gin
RUN go get github.com/Unknwon/com
RUN go get github.com/kelseyhightower/envconfig
RUN go get github.com/Unknwon/com
RUN go get -u github.com/jinzhu/gorm
# こっちのドライバーを使用する
RUN go get github.com/lib/pq
RUN eval $(cat .env)
