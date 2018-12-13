FROM alpine:3.8

LABEL MAINTAINER "Aurelien PERRIER <a.perrier89@gmail.com>"
LABEL APP "tfversion"

ENV TERRAFORM_PATH /root/.tfversion/bin
ENV PATH "$PATH:${TERRAFORM_PATH}"

COPY ./bin/tfversion /usr/bin
