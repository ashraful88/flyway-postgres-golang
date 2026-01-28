# Flywaydb PostgreSQL Golang boilerplate

Use [Flyway](https://github.com/flyway/flyway) database schema migration tool with any Golang project.

This repo uses PostgreSQL but can be used for any SQL database that Flyway supports.

Read more about versioning SQL files:
https://flywaydb.org/documentation/concepts/migrations.html#naming

## Prerequisites

- Go 1.20+
- Docker and Docker Compose

## Add database schema files

Add your .sql files in `./sql` directory.
Use semver versioning in the filename.

Example: `V1.1__anyname.sql`, `V1.2__secondfile.sql`

## Running Flyway in a local development environment

1. Create `.env` from `example.env`:
   ```bash
   cp example.env .env
   ```

2. Start Docker containers (PostgreSQL, Flyway, and PgAdmin):
   ```bash
   make run-docker
   # or
   docker-compose up -d
   ```

   This creates:
   - PostgreSQL database container
   - Flyway container that runs migrations automatically
   - PgAdmin UI at http://localhost:8080/browser/

3. Check migration status:
   ```bash
   docker logs flyway-postgres-golang-flyway-1
   ```

4. Run the Go application locally:
   ```bash
   make run
   ```

### Run migrations manually after SQL changes

```bash
make run-migrate
```

Or run directly with Docker:

```bash
source .env
docker run --rm \
  --network flyway-postgres-golang_flynet \
  -v $PWD/sql:/flyway/sql \
  flyway/flyway:9 \
  -url=jdbc:postgresql://pgdb:5432/$POSTGRES_DB \
  -user=$POSTGRES_USER \
  -password=$POSTGRES_PASSWORD \
  -connectRetries=60 \
  migrate
```

### Reset database and migrate fresh

```bash
source .env
docker run --rm \
  --network flyway-postgres-golang_flynet \
  -v $PWD/sql:/flyway/sql \
  flyway/flyway:9 \
  -url=jdbc:postgresql://pgdb:5432/$POSTGRES_DB \
  -user=$POSTGRES_USER \
  -password=$POSTGRES_PASSWORD \
  -connectRetries=60 \
  clean migrate
```

### Stop containers

```bash
make run-docker-down
# or
docker-compose down
```

## Running Flyway in the production environment

For production, use the Dockerfile. Set `ENV=production` in your environment.

In production, database migration is executed from the Go application on startup. Note that during CI/CD the database may not be accessible, but once the Docker image is deployed to Kubernetes or another environment, the database will be accessible for migration.

Migration functions are defined in `./flyway/flyway.go` and triggered from `main.go`. The Docker image uses `flyway/flyway:9-alpine` as the base image to run both the Go binary and database migrations.

## Run Flyway using Dockerfile only

1. Create a `.env` file from `example.env`

2. Build and run:
   ```bash
   docker build -t flywayk8 .
   docker run --env-file .env flywayk8:latest
   ```

## Available Make commands

| Command | Description |
|---------|-------------|
| `make run` | Run the Go server locally |
| `make run-docker` | Start Docker containers |
| `make run-docker-down` | Stop Docker containers |
| `make run-migrate` | Run database migrations |
| `make run-linter` | Run golangci-lint |
| `make run-test` | Run tests |
| `make run-test-coverage` | Run tests with coverage |

## Create project with gonew

Install gonew if not already installed:

```bash
go install golang.org/x/tools/cmd/gonew@latest
```

Create a new project with your project name:

```bash
gonew github.com/ashraful88/flyway-postgres-golang example.com/yourproject
```
