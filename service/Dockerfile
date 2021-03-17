# Build the Go Binary.
FROM golang:1.14.4 as build_api
ENV CGO_ENABLED 0

# Create a location in the container for the source code. Using the
# default GOPATH location.
RUN mkdir -p /service

COPY go.* /service/
WORKDIR /service
RUN go mod download


# Copy the source code into the container.
WORKDIR /service
COPY . .

# Build the admin tool so we can have it in the container. This should not change
# often so do this first.
WORKDIR /service/app/admin
RUN go build -o admin main.go

# Build the service binary. We are doing this last since this will be different
# every time we run through this process.
WORKDIR /service/app/api
RUN go build -o api main.go


# Run the Go Binary in Alpine.
FROM alpine:3.7
COPY --from=build_api /service/app/private.pem /app/private.pem
COPY --from=build_api /service/app/admin/admin /app/admin
COPY --from=build_api /service/app/api/api /app/api
WORKDIR /app

# This container exposes port 8080 to the outside world
EXPOSE 8080

CMD ["/app/api"]

LABEL org.opencontainers.image.title="api"