#!/bin/sh

echo "exporting envs"

if [ -f /flyway/.env ]
then
  export $(cat /flyway/.env | sed 's/#.*//g' | xargs)
fi

flyway migrate -url=jdbc:postgresql://$POSTGRES_HOST_PORT/$POSTGRES_DB -user=$POSTGRES_USER -password=$POSTGRES_PASSWORD -connectRetries=60
