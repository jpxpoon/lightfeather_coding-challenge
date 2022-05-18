FROM golang:1.17-alpine AS builder
WORKDIR /home/server/
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o server .

FROM alpine
COPY --from=builder /home/server/server /usr/local/bin/server
EXPOSE 3001
ENTRYPOINT server
