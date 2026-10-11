# Grafana with Prometheus (metrics data)

docker compose --profile go-authorizer-v2 --profile go-inventory-v2 --profile go-payment-v2 --profile go-order-v2 --profile py-mcp-inventory-v2 --profile py-mcp-order-v2 --profile tei-server up 

docker compose --profile go-authorizer-v2 --profile go-inventory-v2 --profile go-payment-v2 --profile go-order-v2 --profile py-mcp-inventory-v2 --profile py-mcp-order-v2 up 

docker compose --profile go-authorizer-v2 --profile go-inventory-v2 --profile go-payment-v2 --profile go-order-v2 --profile py-mcp-inventory-v2 --profile py-mcp-order-v2 --profile py-statistic-a2a-v2 up 

docker compose --profile go-authorizer-v2 --profile go-inventory-v2 --profile go-payment-v2 --profile go-order-v2 up --build

