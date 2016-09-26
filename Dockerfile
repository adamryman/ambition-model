FROM golang:1.7.1

RUN mkdir -p /go/src/github.com/adamryman/ambition-truss

COPY . /go/src/github.com/adamryman/ambition-truss

WORKDIR /go/src/github.com/adamryman/ambition-truss
RUN git remote set-url origin https://github.com/adamryman/ambition-truss && \
	go get -v ./...
	
RUN go install ./...

EXPOSE 55440
EXPOSE 55439

ENV PORT 55440
ENV HTTPPORT 55439

ENTRYPOINT ambition-server -grpc.addr=:$PORT -http.addr=:$HTTPPORT
