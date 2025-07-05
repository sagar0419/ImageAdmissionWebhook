FROM golang:1.24-bullseye
WORKDIR /ImageAdmissionWebhook
COPY go.mod go.sum ./
COPY . .
RUN go mod tidy 
ENTRYPOINT ["go", "run", "/ImageAdmissionWebhook/main.go"]