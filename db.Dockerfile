FROM mysql:8.0.23

COPY ./migrations/*.sql /docker-entrypoint-init.d/

