FROM golang:1.18.1

RUN mkdir /api

ADD . /api

WORKDIR /api

RUN go build -o main .

EXPOSE 8080

CMD ["/api/main"]