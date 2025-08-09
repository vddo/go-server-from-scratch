# Build stage
FROM golang:1.24.6-alpine3.22 AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENDABLED=0 GOOS=linux go build -o /first-server

# Final stage
FROM alpine:3.22
COPY --from=builder /first-server /first-server
EXPOSE 8080
CMD [ "/first-server" ]

