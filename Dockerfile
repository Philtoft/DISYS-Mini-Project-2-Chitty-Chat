FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on

COPY go.mod /build
COPY go.sum /build/

RUN cd /build/ && git clone https://github.com/Philtoft/DISYS-Mini-Project-2-Chitty-Chat.git
RUN cd /build/DISYS-Mini-Project-2-Chitty-Chat/server && go build ./server.go

EXPOSE 9080

ENTRYPOINT  "/build/DISYS-Mini-Project-2-Chitty-Chat/server/server" ]