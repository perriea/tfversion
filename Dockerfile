FROM alpine:3.7

LABEL MAINTAINER "Aurelien PERRIER <a.perrier89@gmail.com>"
LABEL APP "tfversion"

ENV TERRAFORM_PATH /root/.tfversion/bin
ENV PATH "$PATH:${TERRAFORM_PATH}"

COPY ./tfversion /usr/bin
