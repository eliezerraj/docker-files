# docker build -t neo4j-db .
# docker run -d --name neo4j-db -p 7474:7474 -p 7687:7687 neo4j-db

FROM neo4j:latest

# Set environment variables for configuration (optional)
ENV NEO4J_AUTH="neo4j/admin1234"

# Expose ports (optional, as the base image usually does this)
EXPOSE 7474 7687

# Define the default command to run when the container starts
# The official Neo4j image handles this, so it's often not needed here.
# CMD ["neo4j"]