FROM golang:1.27-alpine AS builder

WORKDIR /app

COPY go.mod go.sum* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/build ./cmd

FROM scratch

COPY --from=builder /app/build /build

ENTRYPOINT ["/build"]

