# go-rest-starter

A starter for a Golang REST API.

## Getting started

You can use Docker and docker-compose for development

```sh
cp .env.example .env
docker-compose up
```

- Try visiting http://localhost:8080/ for adminer (see .env.example for DB login)
- Click "SQL Command", and copy content from `./files/db/create_table.sql` there, and click "Execute"
- Finally, visit http://localhost:8000/api/users to get the users.

### SSH

```sh
$ docker exec -it <container-id> /bin/sh
$ docker-compose exec api sh
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
