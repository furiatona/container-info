### STAGE 1 : Build the go source code into binary
FROM golang:1.15 as builder

ENV APP_DIR /go/src/github.com/furiatona/container-info

## Copy source code from local machine into container
RUN mkdir -p ${APP_DIR}
COPY . ${APP_DIR}
RUN pwd

# Compile the binary and statically link
RUN cd $APP_DIR && CGO_ENABLED=0 go build && pwd
#Add Bash for exec and debugging
FROM golang:1.11-alpine3.8 as bash_builder
ENV APP_DIR /go/src/github.com/furiatona/container-info
RUN mkdir -p ${APP_DIR}
COPY . ${APP_DIR}
RUN apk add --no-cache bash

## STAGE 2 : Package the binary in a minimal alpine base image
FROM alpine:latest

ARG USER=container-info
ENV HOME /home/$USER
ENV APP_DIR /go/src/github.com/furiatona/container-info

RUN apk add --update sudo \
    && adduser -D $USER \
    && echo "$USER ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/$USER \
    && chmod 0440 /etc/sudoers.d/$USER

USER $USER
WORKDIR $HOME

COPY --from=bash_builder ${APP_DIR} .
COPY --from=builder ${APP_DIR} .
EXPOSE 8080
CMD ["./container-info"]
