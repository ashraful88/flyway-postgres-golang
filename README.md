# Flywaydb PostgreSQL Golang boilerplate  

Use Flyway(https://github.com/flyway/flyway) database schema migration tool with any Golang project.  

This repo uses PostgreSQL but can be used for any SQL database that flyway supports

Read more about versioning SQL files
https://flywaydb.org/documentation/concepts/migrations.html#naming

## Add database schema files

Add your .sql files in ./sql
Use the samvar version in the filename
example: V1.1__anyname.sql V1.2__secondfile.sql

## Running flyway in a local development environment

You need a local installation of Golang and Docker.

Create .env from example.env

Run docker containers, this will create a local PostgreSQL (you can update docker-compose.yaml to use any database) docker container and another docker container will run flyway migration to Postgres. There is a container for PgAdmin(UI tool for Postgres) URL http://localhost:8080/browser/

 `docker-compose up -d`

Check migration status `docker logs flyway`

If migration is successful then it's ready to work with Golang local installation.

Run migration after changes are made in the SQL file.

`docker run --rm --network flyway-postgres-golang_flynet -v $PWD/sql:/flyway/sql flyway/flyway -url=jdbc:postgresql://$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB -user=$POSTGRES_USER -password=$POSTGRES_PASSWORD -connectRetries=60 migrate`

Delete everything in the database and migrate fresh

`docker run --rm --network flyway-postgres-golang_flynet -v $PWD/sql:/flyway/sql flyway/flyway -url=jdbc:postgresql://$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB -user=$POSTGRES_USER -password=$POSTGRES_PASSWORD -connectRetries=60 clean migrate`



## Running flyway in the production environment

For production, Use Dockerfile. 
Set ENV var to `production` in .env

In production database migration will be executed from Golang.
Note that during CI/CD process database may not be accessible but once docker image is built and deploy to Kubernetes pods or any other environment database will be accessible for migration.

In ./flyway/flyway.go flyway migration-related functions are added. Migration is triggered from the main.go file. 
Docker image flyway/flyway:8-alpine will run Golang binary and database migration.





## Run flyway using Dockerfile only
 
Create a .env file like example.env

   `docker build -t flywayk8 .`
   `docker run flywayk8:latest`
