FROM golang:latest

COPY ./app /go/src/app

WORKDIR /go/src/app

RUN go get -d -v
RUN go install -v

CMD ["app"]