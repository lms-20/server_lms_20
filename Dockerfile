# build stage
FROM golang:alpine AS build
WORKDIR /go/src/app
COPY . .


# RUN go mod init eclaim-api
RUN go mod tidy
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN swag init
# RUN apk add --no-cache git
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/api ./main.go

# final stage
FROM alpine:latest

# Make docker-compose wait for container dependencies be ready
# Solution 1: use dockerize tool -----------------------------
RUN apk add --no-cache openssl
ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

WORKDIR /usr/app



COPY --from=build /go/src/app/bin /go/bin
COPY --from=build /go/src/app/ ./
COPY ./.env.development /
EXPOSE 3030
ENTRYPOINT ENV=DEV /go/bin/api