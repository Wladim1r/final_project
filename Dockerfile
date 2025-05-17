## Build state
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/scheduler ./main.go


## Final state
FROM alpine:3.21

WORKDIR /sched

COPY --from=builder /app/bin/scheduler .
COPY --from=builder /app/web ./web

EXPOSE 7540

CMD ["./scheduler"]
