# NEM Meter Readings

Inserts meter readings from NEM formatted data in CSV file.
Built using Go, Gin, Swagger, Gorm and Postgres

## Clone the project

```
$ git clone https://github.com/watahak/nem-meter-readings.git
```

## Development

Create `.env` from `.env.example`.

- `DB_HOST`: Host of database
- `DB_USER`: Username of database
- `DB_PASSWORD`: Password of database
- `DB_NAME`: Name of database

Running on machine

```
$ go mod tidy
$ go run main.go
```

Running on Docker

```
$ docker-compose build
$ docker-compose up
```

## Meter Readings
Meter reading API (`/meter-readings/upload`) can be viewed on Swagger at http://localhost:8080/swagger/index.html#/

Some CSVs are already included in `/data` directory to use for uploading
- `data.csv` provided dataset
- `dataInvalid200.csv` negative testcase
- `dataInvalid300.csv` negative testcase


## Testing

We are using Go's built in unit testing. Refer to docs on [how to add tests](https://go.dev/doc/tutorial/add-a-test)

To run tests in project

```
$ go test ./pkg/... -v
```
