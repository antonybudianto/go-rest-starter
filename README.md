# go-rest-starter

A starter for a Golang REST API.

## Getting started

You can use Docker and docker-compose for development

```sh
docker-compose up
```

- Try visiting http://localhost:8080/ for adminer and run `create_table.sql` there
- Finally, visit http://localhost:8000/users/ to get the users.

### IDE Integration

Set your GOPATH to `<current-directory>/go`

Example:

```sh
export GOPATH=/Users/antony/code/go-rest-starter/go
```

### SSH

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
