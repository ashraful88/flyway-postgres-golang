#!/bin/sh

echo "exporting envs"

if [ -f .env ]
then
  export $(cat .env | sed 's/#.*//g' | xargs)
fi

docker run --rm --network flyway-postgres-golang_flynet -v $PWD/sql:/flyway/sql flyway/flyway -url=jdbc:postgresql://$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB -user=$POSTGRES_USER -password=$POSTGRES_PASSWORD -connectRetries=60 migrate