# Grafana with Prometheus (metrics data)

  docker compose --profile go-data-sink --profile go-inventory --profile go-cart --profile go-clearance --profile go-worker-event --profile go-order --profile py-stat-inference-a2a --profile py-cart-a2a --profile py-inventory-a2a --profile py-planner-a2a up --build
  docker compose --profile all up -d --build 
  docker compose --profile all up -d --force-recreate
