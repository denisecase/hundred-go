# syntax=docker/dockerfile:1
# Put the above parser directive as first line 
# (optional, but recommended)

# A multi-stage build helps keep the final image smaller
 
# Start from base image - most recent alpine (nice and small)
# AS build_base for multistage
FROM golang:1.17-alpine3.15 AS build_base
ENV GO111MODULE=on
ENV CGO_ENABLED=0 

# Create a new working directory in the image being built
WORKDIR /app

# Install some helpful tools 
RUN apk add --no-cache bash

# Download modules to local cache so we can skip re-
# downloading on consecutive docker build commands
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Add sources from current directory to container working directory
COPY . ./
COPY web ./web/

# Build the first go app and tell where to put it
RUN go build -o docker-hundred-go

# Start fresh from a smaller image for second stage
FROM alpine:3.15
WORKDIR /
COPY --from=build_base /app/docker-hundred-go /docker-hundred-go

# Copy above only copies binary. Explictly copy .env
COPY .env /

# Expose port 8080 for our web app binary to the world
EXPOSE 3000

#  tell Docker what command to execute when image starts a container
CMD ["/docker-hundred-go"]
#USER nonroot:nonroot
#ENTRYPOINT ["/docker-hundred-go"]

