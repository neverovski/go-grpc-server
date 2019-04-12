FROM postgres:10.3

COPY ./scripts/docker-entrypoint-initdb.d/news.sql /docker-entrypoint-initdb.d/news.sql

CMD ["postgres"]