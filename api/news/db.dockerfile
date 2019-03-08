FROM postgres:10.3

COPY news.sql /docker-entrypoint-initdb.d/news.sql

CMD ["postgres"]