docker build -t hundred-go .
docker run --env-file .env -p 3000:3000 -it hundred-go
