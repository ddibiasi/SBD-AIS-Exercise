#!/bin/sh
set -e   

#delete the containers if they exists
docker rm -f orderservice orders-postgres

#Start PostgreSQL container
docker run -d \
  --name orders-postgres \
  -p 5432:5432 \
  -p 3000:3000 \
  --env-file debug.env \
  -e PGDATA=/var/lib/postgresql/18/docker \
  -v pg18:/var/lib/postgresql/18/docker \
  postgres:18

#build the go application image using dokerfile
docker build -t orderservice .


#run the go application contaner
docker run -d \
  --name orderservice \
  --network container:orders-postgres \
  --env-file debug.env \
  orderservice

echo "All Ready!"
