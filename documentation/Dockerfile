FROM    		ubuntu:latest
MAINTAINER      Jorden Lowe <supermitsuba@gmail.com>

RUN     apt-get -qq update
RUN		apt-get -qq install golang
RUN		apt-get -qq install git

RUN		mkdir /tmp/GO
RUN		mkdir /tmp/GO/src
ENV		GOPATH /tmp/GO

RUN		git clone https://github.com/supermitsuba/RestApiDiscovery.git /tmp/GO/src/RestApiDiscovery

RUN		apt-get -qq install mercurial

WORKDIR /tmp/GO/src/RestApiDiscovery
RUN		go get
RUN		go test
RUN		go build