FROM golang:1.18.1-alpine3.15

EXPOSE 8080

WORKDIR /usr/src/og

COPY . .

CMD [ "go", "run", "./cmd/onyxforum-og/", "-host", "0.0.0.0" ]
