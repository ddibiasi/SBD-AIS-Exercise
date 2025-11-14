# Software Architecture for Big Data - Exercise 5
Today we are going to add two new microservices: A reverse proxy and a web server.

Typically the frontend (i.e. HTML, CSS, JS, static content) is served
by [nginx](https://nginx.org/), [Apache HTTP Server](https://httpd.apache.org/), or a similar. In this exercise we are going
to serve the frontend for the order service with [sws](https://github.com/static-web-server/static-web-server)
and add [traefik](https://doc.traefik.io/traefik/) as application proxy / load balancer.

## Todo 
- [ ] Add [sws](https://github.com/static-web-server/static-web-server) to docker compose and serve the `./frontend` folder
- [ ] Add [traefik](https://doc.traefik.io/traefik/) reverse proxy
  - sws should be reachable at http://localhost
  - The orderservice should be reachable at http://orders.localhost

## Tips and Tricks
The documentation for traefik provides a [quickstart example](https://doc.traefik.io/traefik/expose/docker/) for Docker compose.
This can be adapted to serve the orderservice and sws. Traefik must not use any other network than w

The Docker socket mapping on Windows might look different than on Linux and OSX. 
If your other Docker containers cannot be found, have a look [here](https://stackoverflow.com/questions/57466568/how-do-you-mount-the-docker-socket-on-windows/62176649#62176649),
[here](https://github.com/docker/for-win/issues/4642#issuecomment-567811455) or [here](https://community.traefik.io/t/how-to-run-on-windows-host-with-docker-provider/4834/4).

The orderservice uses the port 3000, keep in mind when setting up its traefik labels.

The static web server can be [configured via environment variables](https://static-web-server.net/configuration/environment-variables/),
to expose port 80 and serve the frontend folder.

Below are the steps I took in solving the above task:

Exercise 5 — Orderservice (Mustapha Oluwatoyin Gali)
1. Overview

This exercise implements the Orderservice using Go, Docker, PostgreSQL, and Traefik as the reverse proxy.
It includes API endpoints for managing drinks, orders, totals, and database interactions.

2. Technologies Used

Go 1.23+

Docker & Docker Compose

Traefik Reverse Proxy

PostgreSQL

Swagger / OpenAPI documentation

3. Folder Structure
Exc_5/
│── skeleton/
│   ├── main.go
│   ├── model/
│   ├── rest/
│   ├── repository/
│   ├── docs/
│   ├── docker-compose.yml
│   └── Dockerfile
│── README.md

4. How to Build & Run
Option A — Using Docker Compose
cd Exc_5/skeleton
docker-compose up --build

Option B — Manual Run
go mod tidy
go run main.go

5. API Documentation

After running the service, open Swagger UI:

http://localhost:8080/swagger/index.html


OR if behind Traefik:

http://orders.localhost/swagger/index.html

6. Author

Gali Mustapha Oluwatoyin
AIS — Software Architecture for Big Data
Johannes Kepler University Linz