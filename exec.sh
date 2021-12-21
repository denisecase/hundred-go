#!/bin/sh
docker build --tag docker-hundred-go:latest .
docker run --env-file .env --publish 3000:3000 -it docker-hundred-go
