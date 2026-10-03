#!/bin/bash
BASE_PATH=http://localhost:8083

sleep 30

until [ "$(curl -s -o /dev/null -w "%{http_code}" "$BASE_PATH/connectors")" -eq 200 ]; do
  echo "Connectors is not yet active. Retrying in 5 seconds..."
  sleep 5
done

STATUS=$(curl -s -o /dev/null -w "%{http_code}" -X POST -H "Content-Type: application/json" \
  $BASE_PATH/connectors \
  -d '{
  "name": "opensearch-connector-go-ecommerce",
  "config": {
    "name": "opensearch-connector-go-ecommerce",
    "connector.class": "io.aiven.kafka.connect.opensearch.OpensearchSinkConnector",
    "tasks.max": "1",
    "topics": "kfk.logs.go.inventory,kfk.logs.go.order,kfk.logs.go.payment,kfk.logs.go.authorizer",
    "connection.url": "http://opensearch:9200",
    "connection.username": "admin",
    "connection.password": "9a!39cbE3",
    "connection.timeout.ms": "3000",
    "key.converter.schemas.enable": "false",
    "value.converter.schemas.enable": "false",
    "value.converter": "org.apache.kafka.connect.json.JsonConverter",
    "schema.ignore": "true",
    "key.ignore": "true",
    "type.name": "_doc",
    "behavior.on.null.values": "ignore",
    "behavior.on.malformed.documents": "report",
    "external.resource.usage": "index",
    "errors.tolerance": "all",
    "errors.retry.timeout": "60000",
    "errors.retry.delay.max.ms": "5000",
    "errors.deadletterqueue.topic.name": "kfk.logs.go.ecommerce.dlt",
    "errors.deadletterqueue.context.headers.enable": "true",
    "errors.log.enable": "true",
    "retry.backoff.ms": "1000",
    "max.retries": "5",
    "flush.timeout.ms": "10000",
    "linger.ms": "3000",
    "max.buffered.records": "2000",
    "max.in.flight.requests": "5",
    "batch.size": "10",
    "transforms": "IndexReplace,TimestampRouter",
    "transforms.IndexReplace.type": "org.apache.kafka.connect.transforms.RegexRouter",
    "transforms.IndexReplace.regex": "^kfk\\.logs\\.go\\.(.*)$",
    "transforms.IndexReplace.replacement": "logs-go-$1",

    "transforms.TimestampRouter.type": "org.apache.kafka.connect.transforms.TimestampRouter",
    "transforms.TimestampRouter.topic.format": "${topic}-${timestamp}",
    "transforms.TimestampRouter.timestamp.format": "yyyy.MM.dd"
  }
}')

echo "Opensearch connector status $STATUS"