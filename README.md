# FARMERS MARKET
API for information about farmers market in the São Paulo city

## WEBAPI with GOLANG
In this API was used Go as main language, Postgresql to store data and Mux as web framework

### API Docs

- JSON file at docs/farmers_market.postman_collection.json

### Run Database

```
$ docker-compose up --build
```

### Run Server

First, it is necessary export ENV VARS to the system.

For linux systems run:
```
$ export PSQL_USERNAME=teste
$ export PSQL_PASSWORD=teste
$ export PSQL_DATABASE=markets
$ export PSQL_HOSTNAME=db
$ export PSQL_PORT=5432
```

Them, you  can run:
```
$ go run cmd/server/main.go
```

### Migrations

- Migrations is auto runned when start the application

### Seeds

- Run seeds is manual, you have to run 

```
$ go run internal/database/seeder/main.go -filename <some_csv_file.csv>
```

### Doker
If you desire run containerized API , in the docker-compose.yml uncomment the api service.

Then run `$ docker-compose up --build`

To run seeds : `$ docker-compose exec api sh -c ./app/seeder -filename=<some_csv_file.csv>`
