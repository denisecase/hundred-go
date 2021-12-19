#!/usr/bin/env bash
docker build -t hundred-go .
docker run --env-file .env.empty -p 3000:3000 -it hundred-go
