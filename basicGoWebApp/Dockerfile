FROM golang:1.10
WORKDIR /go/src/web
ADD . .
RUN go build -o web .
CMD ./web