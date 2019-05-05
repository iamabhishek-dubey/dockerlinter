FROM alpine:latest

MAINTAINER Abhishek Dubey <abhishek.dubey@opstree.com>

ARG PACKER_VERSION="1.2.5"

ARG TERRAFORM_VERSION="0.11.7"

WORKDIR /root/hashicorp

RUN apk update && \
    apk add wget unzip git

RUN wget https://releases.hashicorp.com/packer/${PACKER_VERSION}/packer_${PACKER_VERSION}_linux_amd64.zip && \
    unzip packer_${PACKER_VERSION}_linux_amd64.zip && \
    mv packer /usr/bin/ && \
    rm packer_${PACKER_VERSION}_linux_amd64.zip

RUN wget https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
    unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
    mv terraform /usr/bin/ && \
    rm terraform_${TERRAFORM_VERSION}_linux_amd64.zip

WORKDIR /root/hashicorp
