FROM golang:1.14
ARG USER_UID
ARG USER_GID
ARG USERNAME

RUN groupadd --gid ${USER_GID} ${USERNAME} \
    && useradd --uid ${USER_UID} --gid ${USER_GID} -m ${USERNAME} \
    #
    # [Optional] Add sudo support. Omit if you don't need to install software after connecting.
    && apt-get update \
    && apt-get install -y sudo \
    && echo ${USERNAME} ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/${USERNAME} \
    && chmod 0440 /etc/sudoers.d/${USERNAME}

ARG GO_PATH=/home/${USERNAME}/go
ARG SRC_PATH=$GO_PATH/src/gin_tutorial

WORKDIR $SRC_PATH
COPY . $SRC_PATH
RUN export GOPATH=$GO_PATH
RUN chown -R $USERNAME:$USERNAME $GO_PATH
RUN eval $(cat $SRC_PATH/.env)
