# Attach the same grafana network

    docker network ls

    docker inspect grafana_grafana_network

# Remove phrase private key

    openssl rsa -in private.key -out private-no-key.key

# copy the certs

    certificate.crt
    private.key

Remove phrase 

    openssl rsa -in private.key -out private-no-pass.key 
