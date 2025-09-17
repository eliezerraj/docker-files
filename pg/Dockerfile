# docker build -t postgre:16 .
# docker run -d --name postgre-16 -p 5432:5432 postgre-16
FROM postgres:16

# Set environment variables for the PostgreSQL database
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres
ENV POSTGRES_DB=postgres

# Expose the default PostgreSQL port
EXPOSE 5432
