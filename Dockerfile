#syntax=docker/dockerfile:1

FROM golang:1.19-alpine3.17
WORKDIR /app
COPY ./src ./src
RUN cd src && go mod tidy && go build -o indexer 
EXPOSE 1234
CMD ./src/indexer
