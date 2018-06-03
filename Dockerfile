# pulling a lightweight version of golang
FROM golang:1.8-alpine
MAINTAINER Mauricio Ribeiro <maumribeiro@gmail.com>
RUN apk --update add --no-cache git
RUN apk add --no-cache bash

# Copy the local package files to the container's workspace.
ADD . /go/src/cv-server-rest-go
WORKDIR /go/src/cv-server-rest-go

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go get github.com/gorilla/mux
RUN go get gopkg.in/mgo.v2
RUN go get cv-server-rest-go

# Run the command by default when the container starts.
ENTRYPOINT ["/go/bin/cv-server-rest-go"]

# Document that the service listens on port 9000.
EXPOSE 9000