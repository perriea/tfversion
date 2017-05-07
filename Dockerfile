FROM golang:1.8.1-alpine

ENV tfversion_path /go/src/github.com/perriea/tfversion/
ENV terraform_path /root/terraform/bin
ENV PATH "$PATH:$terraform_path"

RUN mkdir -p $tfversion_path && \
    mkdir -p /root/repo && \
    mkdir -p /root/terraform/bin

ADD . $tfversion_path

RUN apk -Uuv add groff less python py-pip py-virtualenv git openssh make && \
	  pip install awscli && \
	  apk --purge -v del py-pip && \
	  rm /var/cache/apk/*

RUN cd $tfversion_path && \
    go build . && \
    cp $tfversion_path/tfversion $terraform_path

VOLUME ['/root/.aws', '/root/.ssh']

WORKDIR /root/repo

RUN tfversion install 0.9.4
