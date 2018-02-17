FROM alpine:3.7

ENV terraform_path /root/.tfversion/bin
ENV PATH "$PATH:$terraform_path"

COPY ./tfversion /usr/bin
