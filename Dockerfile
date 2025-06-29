FROM golang:1.24-bullseye
WORKDIR /k8sController
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
ENTRYPOINT ["go run main.go"]