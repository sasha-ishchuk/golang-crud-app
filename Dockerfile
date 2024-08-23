FROM golang:1.23-alpine AS build
LABEL authors="alexa"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o golang-crud-app .

FROM alpine:3.13

WORKDIR /app

COPY --from=build /app/golang-crud-app .

EXPOSE 8080
CMD ["./golang-crud-app"]