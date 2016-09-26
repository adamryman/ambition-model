FROM golang:1.7.1

RUN mkdir -p /go/src/github.com/adamryman/ambition-truss


COPY . /go/src/github.com/adamryman/ambition-truss

WORKDIR /go/src/github.com/adamryman/ambition-truss
RUN git remote set-url origin https://github.com/adamryman/ambition-truss && \
	go get -v ./...
	
RUN go install ./...

ENV PORT 55440

ENTRYPOINT ambition-server -grpc.addr=:$PORT
