# Grafana with Prometheus (metrics data)

  docker compose --profile go-data-sink --profile go-inventory --profile go-cart --profile go-clearance --profile go-worker-event --profile go-order up
  docker compose --profile all up -d --build 
  docker compose --profile all up -d --force-recreate
