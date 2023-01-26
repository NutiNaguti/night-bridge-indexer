# Build stage
FROM golang:1.19-alpine3.17 as build-env

ADD . /dockerdev
WORKDIR /dockerdev
RUN go build -o indexer 

FROM alpine:latest 

WORKDIR /
COPY --from=build-env /dockerdev/indexer /

EXPOSE 1234

CMD ./indexer
