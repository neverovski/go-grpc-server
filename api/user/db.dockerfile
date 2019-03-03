FROM postgres:10.3

COPY user.sql /docker-entrypoint-initdb.d/user.sql

CMD ["postgres"]