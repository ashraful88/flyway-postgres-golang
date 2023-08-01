#!/bin/sh

# run flyway migration from local machine to docker container
# this file is not intent to be used in production
echo "exporting envs"

if [ -f .env ]
then
  export $(cat .env | sed 's/#.*//g' | xargs)
  echo $(cat .env | sed 's/#.*//g' | xargs)
fi


docker run --rm --network "${PWD##*/}_flynet" -v $PWD/sql:/flyway/sql flyway/flyway -url=jdbc:postgresql://pgdb:$POSTGRES_PORT/$POSTGRES_DB -user=$POSTGRES_USER -password=$POSTGRES_PASSWORD -connectRetries=1 migrate

