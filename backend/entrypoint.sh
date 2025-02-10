#!/bin/sh
set -e

cd "src/db/schema"
goose postgres "postgres://postgres:postgres@db:5432/gofeed" status
goose postgres "postgres://postgres:postgres@db:5432/gofeed" up
cd "../../.."

# Execute the main command passed via CMD
exec "$@"