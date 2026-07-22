# Grafana with Prometheus (metrics data)

sudo docker compose --profile go-data-sink --profile go-inventory --profile go-cart --profile go-clearance --profile go-worker-event --profile go-order --profile py-stat-inference-a2a  --profile go-federated-registry --profile tei-server --profile py-mcp-sales-server --profile py-mcp-inventory-server up --build

--profile py-planner-a2a  --profile py-mcp-inventory-server up 

--build

docker compose --profile go-data-sink --profile go-inventory --profile go-cart up --build 

docker compose --profile all up -d --force-recreate

sudo docker compose --profile go-data-sink up -d --force-recreate