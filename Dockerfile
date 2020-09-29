FROM flyway/flyway:7.0-alpine AS migrator

COPY .env /flyway/.env
COPY ./run_migration.sh /flyway/run_migration.sh
COPY ./sql/ /flyway/sql/

USER root
RUN chmod +x /flyway/run_migration.sh
RUN cat /flyway/run_migration.sh
ENTRYPOINT ["/flyway/run_migration.sh"]
