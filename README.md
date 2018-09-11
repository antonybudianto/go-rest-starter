# go-rest-starter

A starter for a Golang REST API.

## Docker

You can use Docker and docker-compose for development

```
docker-compose up
```

It'll be served at http://localhost:8000

Try visiting http://localhost:8000/users

## SSH
```sh
$ docker exec -it 7bb8b885ea2d /bin/sh
```

## Production

```sh
docker-compose -f docker-compose.production.yml build
docker-compose -f docker-compose.production.yml up
```

## Credits

https://github.com/ilourt/gocker

## License

MIT
