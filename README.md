# go-rest-starter

A starter for a Golang REST API.

## Getting started

You can use Docker and docker-compose for development

```sh
docker-compose up
```

- Try visiting http://localhost:8080/ for adminer and run `create_table.sql` there
- Finally, visit http://localhost:8000/users/ to get the users.

### SSH

```sh
$ docker exec -it <container-id> /bin/sh
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
