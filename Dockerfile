FROM golang:1.7.1

RUN wget -O /usr/local/bin/dumb-init https://github.com/Yelp/dumb-init/releases/download/v1.2.0/dumb-init_1.2.0_amd64
RUN chmod +x /usr/local/bin/dumb-init

RUN mkdir -p /go/src/github.com/adamryman/ambition-model

COPY . /go/src/github.com/adamryman/ambition-model

WORKDIR /go/src/github.com/adamryman/ambition-model
	
RUN go install -v ./...
RUN ls $GOPATH
RUN ls $GOPATH/bin

EXPOSE 55440
EXPOSE 55439

ENV PORT 55440
ENV HTTPPORT 55439

ENTRYPOINT ["/usr/local/bin/dumb-init", "--"]
CMD ["/go/bin/ambition-server", "-grpc.addr", ":55440", "-http.addr", ":55439"]
