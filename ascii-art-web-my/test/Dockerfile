FROM golang 
#:onbuild

MAINTAINER mini-alem

RUN mkdir /ascii-docker

ADD . /ascii-docker
WORKDIR /ascii-docker/main
RUN go build main.go

CMD ["/ascii-docker/main/main"]
