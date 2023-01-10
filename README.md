# Chuck
## How to run
```sh
docker-compose up
```
## To access the API
```sh
http://localhost:8085/api/v1/chuck
```

## Development Environment
If you do not want to use **`air`**, you will need to uncomment the following lines in the **`dev.Dockerfile file:`**
```sh
# RUN go build -o main
# CMD ["./main"]
```
and comment or remove the following lines:
```sh
RUN go install github.com/cosmtrek/air@latest
CMD ["air", "-c", ".air.toml"]
```
Additionally, you will need to comment the following lines in the **`docker-compose.yml`** file:
```sh
volumes:
  - .:/app
```