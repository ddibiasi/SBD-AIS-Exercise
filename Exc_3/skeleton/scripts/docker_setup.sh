#!/bin/bash

cd ..

docker run --env-file debug.env -p 5432:5432 -d -v my_postgres_data:/var/lib/postgresql postgres:18

docker buildx build -t rest_drinks:latest .

docker run --env-file debug_2.env -p 8080:3000 rest_drinks:latest