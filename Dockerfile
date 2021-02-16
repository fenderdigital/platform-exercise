FROM golang:1.15.6
WORKDIR /app/
COPY src/go.mod src/go.sum /app/
RUN go mod download -x
COPY ./ ./
