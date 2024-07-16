#!/bin/sh

set -e

echo "Waiting for postgres to be ready..."

while ! pg_isready -h db -p 5432 ; do
  echo "still waiting"
  sleep 2
done

echo "Running migrations..."
goose -dir ./sql/schema/ postgres postgres://postgres:novek@db:5432/fechronizo up

echo "Migrations completed"