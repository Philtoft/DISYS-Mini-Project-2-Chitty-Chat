FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on

COPY go.mod /build
COPY go.sum /build/

RUN cd /build/ && git clone https://github.com/Philtoft/DIS-mini-project-1.git
RUN cd /build/DIS-mini-project-1/server && go build ./server.go

EXPOSE 9080

ENTRYPOINT [ "/build/DIS-mini-project-1/server/server" ]
