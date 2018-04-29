# go-rest-starter

A starter for a Golang REST API.

## Install

Use official dependency manager, [dep](https://github.com/golang/dep)

```
dep ensure
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
docker-compose -f docker-compose.production.yml build
docker-compose -f docker-compose.production.yml up
```

## Credits

https://github.com/ilourt/gocker

## License

MIT
