FROM golang:1.15.0-alpine3.12

RUN mkdir /app

ADD . /app

WORKDIR /app

#Download chromium for the scraper
RUN apk update && apk add chromium

RUN go mod download
RUN go build -o main ./cmd/.

EXPOSE 3001:3001

CMD ["/app/main"]