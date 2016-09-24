FROM golang:1.7.1

RUN mkdir -p /go/src/github.com/adamryman/ambition-truss/ambition

COPY . /go/src/github.com/adamryman/ambition-truss/ambition

WORKDIR /go/src/github.com/adamryman/ambition-truss/ambition
RUN go get -v -u -f github.com/adamryman/ambition-truss/...
	

RUN go install ./...

RUN mkdir -p /home/.config/ambition
ENV HOME /home
COPY ./config.json /home/.config/ambition

ENTRYPOINT ambition-server
