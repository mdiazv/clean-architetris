FROM golang:1.12 AS builder

ENV GO111MODULE=on
WORKDIR '/app'

COPY go.mod .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM alpine:latest
WORKDIR '/app'
COPY --from=builder /app/clean-architetris . 

EXPOSE 8080
CMD ["/app/clean-architetris"]
