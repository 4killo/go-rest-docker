FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD 123
ENV MYSQL_DATABASE customers
ENV MYSQL_USER all
ENV MYSQL_PASSWORD 123

ADD setup.sql /docker-entrypoint-initdb.d

EXPOSE 3306
