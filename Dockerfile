FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /wbgoodsfeed ./cmd/wbgoodsfeed

FROM alpine:3
COPY --from=builder /wbgoodsfeed /wbgoodsfeed
ENTRYPOINT ["/wbgoodsfeed"]