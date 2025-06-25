FROM golang:1.24 AS builder
WORKDIR /k8sController

COPY go.mod go.sum ./
RUN go mod tidy
COPY . .

FROM gcr.io/distroless/static
COPY --from=builder /k8sController/controller /controller
COPY tls /tls
ENTRYPOINT ["/controller"]