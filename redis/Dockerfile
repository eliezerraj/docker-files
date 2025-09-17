# docker build -t redis .
# docker run -d --name redis-memory -p 6379:6379 redis
# docker exec -it redis-memory redis-cli // test-it $PING

# Use the official Redis base image
FROM redis:latest

# Expose the default Redis port
EXPOSE 6379

# You can add a custom redis.conf if needed
# COPY redis.conf /usr/local/etc/redis/redis.conf
# CMD ["redis-server", "/usr/local/etc/redis/redis.conf"]

# If no custom config is provided, the default Redis server command will run
CMD ["redis-server"]
