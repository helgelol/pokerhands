# Run with air for development
FROM golang:1.23-alpine as api-dev
WORKDIR /app
RUN go install github.com/air-verse/air@latest
CMD [ "air main.go" ]

# Build clean binary for production use
# Consider doing this with ko for smaller footprint
FROM golang:1.23-alpine as api-prod
WORKDIR /app
COPY /api /app
RUN go build -o main
CMD [ "./main" ]
