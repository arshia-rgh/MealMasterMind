# Used postgres docker file to run the migrations
FROM postgres:latest

WORKDIR /docker-entrypoint-initdb.d

COPY migrations/ /docker-entrypoint-initdb.d/

