FROM golang:alpine

WORKDIR /go/src

COPY . .

RUN go build -o scraper .

CMD ["./scraper"]