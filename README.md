# golang-web-api-postgres-simple

## Description
This returns a data dump of the table `pop`
as a JSON encoded string. An ORM is used to
create and seed the database; even though scripts
are included, `db/sql`, they are for reference only.

## Tech stack
- bash
- golang 1.13

## Docker stack
- ubuntu:latest
- postgres

## Build notes
The service takes about 1 min before callable.

## To run
`sudo ./install.sh -u`
Available at http://localhost:8080

## To stop
`sudo ./install.sh -d`

## To see help
`sudo ./install.sh -h`

## Credits
- https://www.mindbowser.com/golang-go-with-gorm/
- https://www.enterprisedb.com/postgres-tutorials/postgresql-and-golang-tutorial
- https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/
