FROM golang:1.8.3-alpine

ENV tfversion_path /go/src/github.com/perriea/tfversion/
ENV terraform_path /root/.tfversion/bin
ENV PATH "$PATH:$terraform_path"
#ENV GOOGLE_APPLICATION_CREDENTIALS=/root/.gcloud/project.json

RUN mkdir -p $tfversion_path && \
    mkdir -p /root/repo && \
    mkdir -p /root/.tfversion/bin

ADD . $tfversion_path

RUN apk -Uuv add python py-pip py-virtualenv && \
	  pip install awscli && \
	  apk --purge -v del py-pip && \
	  rm /var/cache/apk/*

RUN cd $tfversion_path && \
    go build . && \
    cp $tfversion_path/tfversion /go/bin

VOLUME ['/root/.aws', '/root/.gcloud', '/root/.ssh']

WORKDIR /root/repo

RUN tfversion install --version 0.9.6
