# GoFeed

## Development Commands

### Generates type-safe Go code from SQL

```bash
sqlc generate
```

### Apply all available migrations

```bash
goose postgres postgres://postgres:postgres@localhost:5432/gofeed up
```

### Roll back a migration

```bash
goose postgres postgres://postgres:postgres@localhost:5432/gofeed down
```

### Print the status of all migrations

```bash
goose status
```
