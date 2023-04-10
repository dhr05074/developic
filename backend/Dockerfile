# Stage 1: Build stage
FROM golang:alpine AS build

# Disable CGO to make the executable more portable
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

RUN go mod download

# Build the executable
RUN go build -o main gateway/cmd/main.go

# Stage 2: Final stage
FROM scratch

# Copy the executable from the previous stage
COPY --from=build /app/main /main

# Set the entrypoint to the executable
ENTRYPOINT ["/main"]
