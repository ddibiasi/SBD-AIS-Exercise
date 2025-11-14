# Software Architecture for Big Data - Exercise 4

Software Architecture for Big Data — Exercise 4

Topic: Docker Compose setup for orderservice and PostgreSQL

📘 Overview

In this exercise, we extend the previous implementation (Exercise 3) by containerizing the orderservice backend and PostgreSQL database within a single Docker Compose setup.
The goal is to:

Run both services in isolated containers

Connect them through a shared Docker network

Manage ports and environment variables properly

Verify that the backend connects successfully to the database

Exc_4/
│
├── skeleton/
│   ├── docs/                 # Documentation and code references
│   ├── frontend/             # Simple HTML interface (if applicable)
│   ├── model/                # Go data models
│   ├── repository/           # Database connection logic (db.go)
│   ├── rest/                 # REST API routes and handlers
│   ├── scripts/              # Utility and build scripts
│   ├── Dockerfile            # Build configuration for orderservice
│   ├── docker-compose.yml    # Multi-container setup (service + DB)
│   ├── go.mod / go.sum       # Go module dependencies
│   ├── main.go               # Application entry point
│   └── README.md             # Documentation (this file)

Setup Instructions
Build and start containers

From within the folder:
cd Exc_4/skeleton
docker compose up -d --build

This will:

Build the orderservice image from the provided Dockerfile

Pull the official postgres:16-alpine image

Create a shared network and persistent volume (skeleton_pgdata)

Start both containers automatically

2️⃣ Check running containers
docker compose ps

✅ Expected output:
NAME                  SERVICE             STATUS      PORTS
sbd3-postgres         postgres            Healthy     0.0.0.0:5433->5432/tcp
skeleton-orderservice orderservice        Up          0.0.0.0:3000->3000/tcp

View application logs
docker compose logs -f orderservice

Expected successful output:
INFO Order System starting
INFO Connecting to database
INFO ⚡⚡⚡ Order System is up and running ⚡⚡⚡

Environment Variables

Defined in .env file or directly in docker-compose.yml:

POSTGRES_DB=order
POSTGRES_USER=docker
POSTGRES_PASSWORD=docker
POSTGRES_TCP_PORT=5433
DB_HOST=sbd3-postgres

Note:
The host port was changed to 5433 to avoid conflicts with the existing local or previous exercise PostgreSQL containers.

services:
  postgres:
    image: postgres:16-alpine
    container_name: sbd3-postgres
    environment:
      POSTGRES_DB: order
      POSTGRES_USER: docker
      POSTGRES_PASSWORD: docker
    ports:
      - "5433:5432"
    volumes:
      - skeleton_pgdata:/var/lib/postgresql/data

  orderservice:
    build: .
    container_name: skeleton-orderservice
    depends_on:
      - postgres
    ports:
      - "3000:3000"
    environment:
      POSTGRES_DB: order
      POSTGRES_USER: docker
      POSTGRES_PASSWORD: docker
      POSTGRES_TCP_PORT: 5433
      DB_HOST: sbd3-postgres
    command: ["/app/ordersystem"]

volumes:
  skeleton_pgdata:

Verification

After setup, confirm:

Containers show Healthy / Up status

orderservice log prints ⚡⚡⚡ Order System is up and running

Access via browser or API test:
http://localhost:3000/

should return a valid JSON or HTML response from the Go backend.
to verify that the REST API or frontend is accessible.

Stop and clean up
docker compose down -v

This stops all containers and removes associated volumes and networks.


🧾 Result Summary
Component	Status	Port Mapping	Notes
PostgreSQL (DB)	✅ Healthy	5433 → 5432	Connected to orderservice
Orderservice	✅ Running	3000 → 3000	API reachable and DB connected
Network	✅ Created	skeleton_default	Isolated and functioning as expected
Logs Verified	✅ Yes		“⚡⚡⚡ Order System is up and running”


Troubleshooting
Issue	Possible Cause	Fix
❌ Port 5432 already allocated	Another PostgreSQL instance is already using port 5432.	Change host port in docker-compose.yml to 5433, or stop the conflicting container: docker ps → docker stop <container_id>.
❌ Connection refused / failed to initialize database	Backend starts before PostgreSQL is ready.	Ensure depends_on is set in compose, or restart: docker compose restart orderservice.
❌ Obsolete attribute 'version' warning	Docker Compose v2 no longer requires a version field.	Remove the line version: '3' from docker-compose.yml.
❌ Container already in use	Existing container names conflict.	Remove old containers: docker rm -f orders-db orderservice.
❌ Database not persisting data	Volume not attached or deleted.	Recreate volume: docker volume rm skeleton_pgdata then docker compose up -d.

Author

Mustapha Oluwatoyin Gali
Course: Software Architecture for Big Data (AIS-BA, WS25)
Professor: DI Dr. Daniele Dibiasi