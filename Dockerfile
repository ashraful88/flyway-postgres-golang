#####################################
#   STEP 1: build golang executable #
#####################################
FROM golang:1.19-alpine3.16 AS build

#ENV GO111MODULE on
RUN mkdir -p /etc/secret/
COPY  .env /etc/secret/.env

ARG ENV
ENV ENV $ENV

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/release

ADD . .

RUN CGO_ENABLED=0 go build -a -o main

#####################################
#   STEP 2 build flyway image      #
#####################################
FROM flyway/flyway:8-alpine

COPY .env /app/.env
COPY . /app

COPY --from=build /go/release/main /app/main

WORKDIR /app/
RUN chmod +x /app/main
ENTRYPOINT ["./main"]

