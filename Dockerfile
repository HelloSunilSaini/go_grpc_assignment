FROM golang:alpine as base 

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/grpc/server/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=base /app/main .

ENV GRPC_PORT=8080

EXPOSE 8080

CMD ["./main"]