# scripts/start-orderservice.ps1
param(
  [string]$ImageName     = "orderservice",
  [string]$ContainerName = "orderservice",
  [string]$Network       = "sbd3-net",
  [string]$EnvFile       = "$PSScriptRoot/../.env"
)

$ErrorActionPreference = "Stop"

# Build image from the project root (one level up from /scripts)
docker build -t $ImageName "$PSScriptRoot/.."

# Remove any previous container (ignore errors)
docker rm -f $ContainerName 2>$null | Out-Null

# Run the service on the same network as Postgres, loading .env
docker run -d `
  --name $ContainerName `
  --network $Network `
  --env-file $EnvFile `
  -p 3000:3000 `
  $ImageName

Write-Host "`nOrderservice is running at http://localhost:3000" -ForegroundColor Green
