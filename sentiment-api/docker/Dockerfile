# Build the application from source
FROM golang:1.22 AS build-stage

WORKDIR /app

#RUN apt-get update && apt-get install -y librdkafka-dev

COPY go.mod go.sum ./
#RUN go get github.com/confluentinc/confluent-kafka-go/v2/kafka
#RUN go mod download

# Copy the entire application
COPY . .


# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -o sentiment-api ./cmd/sentiment-api/main.go


# Deploy the application binary into a lean image
FROM gcr.io/distroless/cc-debian12 AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/sentiment-api /sentiment-api

EXPOSE 8080

USER nonroot:nonroot


# Specify the command to run on container start
ENTRYPOINT ["/sentiment-api"]