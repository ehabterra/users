FROM golang:1.15.2-alpine3.12

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build ./cmd/users && go build ./cmd/users-cli
CMD ["/app/users"]

EXPOSE 8080
EXPOSE 8000