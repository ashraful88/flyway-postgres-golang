# Flywaydb PostgreSQL boilerplate  

Use Flyway(flywaydb.org) database migration tool with golang project

Read more about versioning sql files
https://flywaydb.org/documentation/migrations#naming

## Add database schema files

Add you sql files in ./sql
Use samvar version in filename
example: V1.1__anyname.sql V1.2__secondfile.sql

## Running flyway in local development environment

You need local installation of Golang and Docker.

Create .env from example.env

Run docker containers, this will create a local postgresql(you can update docker-compose.yaml to use any database) docker container and another docker container will run flyway migration to postgres. There is a container for PgAdmin(UI tool for postgres) url http://localhost:8080/browser/

 `docker-compose up -d`

Check migration status `docker logs flyway`

If migration successful then ready to work on Golang app in local.

Run migration after changes made in sql

`docker run --rm --network flyway-postgres-golang_flynet -v $PWD/sql:/flyway/sql flyway/flyway -url=jdbc:postgresql://$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB -user=$POSTGRES_USER -password=$POSTGRES_PASSWORD -connectRetries=60 migrate`

Delete database everything and migrate fresh

`docker run --rm --network flyway-postgres-golang_flynet -v $PWD/sql:/flyway/sql flyway/flyway -url=jdbc:postgresql://$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB -user=$POSTGRES_USER -password=$POSTGRES_PASSWORD -connectRetries=60 clean migrate`



## Running flyway in production environment

For production, Use Dockerfile
set ENV var to production in .env

In ./flyway/flyway.go flyway migration related functions are added. migration is triggered from main.go file. 
docker image flyway/flyway:8-alpine will run golang binary and database migration.





## Run flyway using Dockerfile only
 
Create .env file like example.env

   `docker build -t flywayk8 .`
   `docker run flywayk8:latest`
