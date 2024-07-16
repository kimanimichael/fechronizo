# Fechronizo
Fechronizo is an RSS aggregator project. This fetches RSS feeds from the internet depending on all users' subscriptions and updates these feeds in the database accordingly.

## Prerequisites
The following are required to run this service:

1. [Docker](https://docs.docker.com/get-docker/) (version 26.0.0)

To develop fechronizo, you will need:

1. [Go](https://go.dev/doc/install) (version 1.21)
2. [PostgreSQL](https://www.postgresql.org/download/) (version 14.12)

## Installation
Ensure all  pre-requisites are satisfied before carrying out installation and clone the repo

### Using Docker
Execute this command from the project root to run the service

`docker compose -f docker/docker-compose.yml up`

### Without Docker
#### Database Creation

Create a PostgreSQL database and name it appropriately. Check this [reference](https://github.com/Mike-Kimani/fechronizo/blob/master/.env#L2) .The database is named `fechronizo` in this case

#### Apply all Database Migrations
``
cd sql/schema
``

``
goose postgres postgres://{userName}:{password}@localhost:5432/{databaseName} up
``

#### Build and Start the Server
``
go build && ./fechronizo
``