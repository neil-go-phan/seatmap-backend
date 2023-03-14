# Build stage
FROM golang:1.20-alpine3.16 AS builder
LABEL author="Neil GoldenOwl golang intern"
WORKDIR /app
# Copy all file in this foler to /app
COPY . . 
RUN go build -o main main.go
# EXPOSE: declare container's listen port 
EXPOSE 8080

# This img is super large, because it contain all the package that are require in pj. 
# Fix this by using multi-stage build 

# Run stage
FROM alpine:3.16

WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY db/migration ./db/migration
CMD [ "/app/main" ]