FROM golang:1.24 as builder
WORKDIR /k8sController
COPY . .
RUN go mod init k8sController && go mod tidy
RUN go build -o controller

FROM gcr.io/distroless/static
COPY --from=builder /k8sController/controller /controller
COPY tls /tls
ENTRYPOINT ["/controller"]