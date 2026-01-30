# Grafana with Prometheus (metrics data)

  docker-compose up
  docker-compose up -d --force-recreate

## Access Prometheus

  http://localhost:9090/targets

## Access Grafana

  http://localhost:3000/

## Access Loki

  http://loki:3100/ready

  Grafana loki datasource (http://172.17.0.1:3100)

## Alloy 

  curl http://alloy:12345/-/healthy

## Test Loki

  Loki standart

  curl -s -X POST "http://localhost:3100/loki/api/v1/push" \
    -H "Content-Type: application/json" \
    -d '{"streams":[{"stream":{"job":"test"},"values":[["'"$(date +%s%N)"'","{\"message\":\"hello\"}"]]}]}' \
    -w "\nHTTP_CODE:%{http_code}\n"

  Otel standart

  curl -X POST http://localhost:3100/otlp/v1/logs \
    -H "Content-Type: application/json" \
    -d '{
      "resourceLogs": [
        {
          "resource": {
            "attributes": [
              {"key": "service.name", "value": {"stringValue": "demo-service"}}
            ]
          },
          "scopeLogs": [
            {
              "scope": {"name": "demo-logger"},
              "logRecords": [
                {
                  "timeUnixNano": "'"$(($(date +%s%N)))"'",
                  "severityText": "INFO",
                  "body": {"stringValue": "Hello from curl via OTLP 3!"},
                  "attributes": [
                    {"key": "trace_id", "value": {"stringValue": "abc123"}}
                  ]
                }
              ]
            }
          ]
        }
      ]
    }'
  