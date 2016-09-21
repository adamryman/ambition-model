FROM golang:1.7.1

RUN mkdir -p /go/src/github.com/adamryman/ambition-truss


COPY . /go/src/github.com/adamryman/ambition-truss

WORKDIR /go/src/github.com/adamryman/ambition-truss
RUN git remote set-url origin https://github.com/adamryman/ambition-truss && \
	go get -v -u ./...
	
RUN go install ./...

RUN mkdir -p /home/.config/ambition

ENV HOME /home

COPY ./config.json /home/.config/ambition

ENTRYPOINT ambition-server
