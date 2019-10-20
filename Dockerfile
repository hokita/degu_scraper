FROM golang:1.13

RUN apt-get update
RUN apt-get upgrade -y

ENV GOBIN /go/bin

RUN go get github.com/PuerkitoBio/goquery
RUN go get github.com/nlopes/slack

WORKDIR /go
ADD . /go

CMD ["go", "run", "main.go"]
