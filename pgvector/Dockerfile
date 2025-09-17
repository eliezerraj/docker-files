# docker build -t pgvector:16 .
# docker run -d --name pgvector -p 5433:5432 -d pgvector:16
FROM pgvector/pgvector:pg16

# Set environment variables for the PostgreSQL database
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres
ENV POSTGRES_DB=postgres

# Expose the default PostgreSQL port
EXPOSE 5432
