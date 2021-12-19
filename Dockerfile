# Start from golang base image
FROM golang:1.17-alpine3.14

# Define current working directory in the container
WORKDIR /hundred-go

# Download modules to local cache so we can skip re-
# downloading on consecutive docker build commands
COPY go.mod .
COPY go.sum .
RUN go mod download

# Add sources from current directory to container working directory
COPY . .

# Build the go app and tell where to put it
RUN go build -o out/hundred-go .

# Expose port 3000 for our web app binary to the world
EXPOSE 3000

CMD ["/hundred-go/out/hundred-go"]
