FROM postgres:10.3

COPY ./scripts/docker-entrypoint-initdb.d/user.sql /docker-entrypoint-initdb.d/user.sql

CMD ["postgres"]