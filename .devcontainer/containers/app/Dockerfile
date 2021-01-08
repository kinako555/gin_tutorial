FROM golang:1.14
ARG USER_UID
ARG USER_GID
ARG USERNAME

RUN echo ${USER_UID}

RUN groupadd --gid ${USER_GID} ${USERNAME} \
    && useradd --uid ${USER_UID} --gid ${USER_GID} -m ${USERNAME} \
    #
    # [Optional] Add sudo support. Omit if you don't need to install software after connecting.
    && apt-get update \
    && apt-get install -y sudo \
    && echo ${USERNAME} ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/${USERNAME} \
    && chmod 0440 /etc/sudoers.d/${USERNAME}
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
