#Flywaydb PostgreSQL boilerplate  

Flyway(flywaydb.org) is an open-source database migration tool

https://flywaydb.org/documentation/migrations#naming


   `docker run --rm --network flyway_flynet -v $PWD/sql:/flyway/sql flyway/flyway -url=jdbc:postgresql://pgdb/exampledb -user=postgres -password=password -connectRetries=60 migrate`