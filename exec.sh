#!/usr/bin/env bash
docker build -t auth0-golang-web-app .
docker run  -it auth0-golang-web-app
