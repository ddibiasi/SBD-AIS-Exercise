# start-db.ps1
# Run a PostgreSQL container
docker run -d `
  --name orders-db `
  -e POSTGRES_DB=order `
  -e POSTGRES_USER=docker `
  -e POSTGRES_PASSWORD=docker `
  -p 5432:5432 `
  postgres:18

  # Start PostgreSQL for Exercise 3

$Network  = "sbd3-net"
$DbName   = "orders-db"
$Image    = "postgres:18"
$Volume   = "exc3_pgdata"

# Ensure network exists
if (-not (docker network ls --format '{{.Name}}' | Select-String -SimpleMatch $Network)) {
  docker network create $Network | Out-Null
}

# Stop old container (if any)
if (docker ps -a --format '{{.Names}}' | Select-String -SimpleMatch $DbName) {
  docker rm -f $DbName | Out-Null
}

# Start DB with env from current shell (.env you set earlier)
docker run -d `
  --name $DbName `
  --network $Network `
  -p 5432:5432 `
  -e POSTGRES_DB=$env:POSTGRES_DB `
  -e POSTGRES_USER=$env:POSTGRES_USER `
  -e POSTGRES_PASSWORD=$env:POSTGRES_PASSWORD `
  -v ${Volume}:/var/lib/postgresql/data `
  $Image

Write-Host "Postgres is starting as '$DbName' on network '$Network'."
