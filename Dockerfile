FROM golang:1.21

WORKDIR /usr/src/app

COPY go.mod ./
RUN go mod download && go mod verify

COPY *.go ./

RUN go build -v -o /usr/local/bin/app ./

EXPOSE 8080

ENTRYPOINT ["app"]
