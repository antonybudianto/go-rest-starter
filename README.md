# go-rest-starter

A starter for a Golang REST API.

## Install

Use official dependency manager, [dep](https://github.com/golang/dep)

```
dep ensure
```

## Start

```
go run *.go
```

## Docker

You can use Docker and docker-compose for development

```
docker-compose up
```

It'll be served at http://localhost:8000

Try visiting http://localhost:8000/users

## Production

```sh
docker build -t starter-prod:latest -f ./Dockerfile.prod .
docker run starter-prod:latest
```

## Credits

https://github.com/ilourt/gocker

## License

MIT
