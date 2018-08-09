FROM golang:1.10
ADD . .
RUN CGO_ENABLED=0 go build -a -o api .

FROM alpine:3.7
COPY --from=0 /go/api .
CMD ./api