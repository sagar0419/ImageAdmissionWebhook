FROM golang:1.24-bullseye AS builder
WORKDIR /k8sController
RUN apt-get update && apt-get install -y ca-certificates
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o imagecontroller
ENTRYPOINT ["./imagecontroller"]