FROM golang:1.24-bullseye
WORKDIR /k8sController
COPY go.mod go.sum ./
COPY . .
RUN go mod tidy 
ENTRYPOINT ["go", "run", "/k8sController/main.go"]