# Flywaydb PostgreSQL boilerplate  

Flyway(flywaydb.org) is an open-source database migration tool

Read more about versioning sql files
https://flywaydb.org/documentation/migrations#naming



First init: 

 `docker-compose up -d`

Run migration after changes made in sql

   `docker run --rm --network flyway_flynet -v $PWD/sql:/flyway/sql flyway/flyway -url=jdbc:postgresql://pgdb/exampledb -user=postgres -password=password -connectRetries=60 migrate`

PgAdmin http://localhost:8080/
