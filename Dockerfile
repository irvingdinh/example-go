FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

RUN apk add --no-cache curl

WORKDIR /app

COPY --from=0 /app/main ./

ADD config.yaml config.yaml

CMD ["./main", "serve"]
