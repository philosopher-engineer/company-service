# company-service
A typical golang backend service to handle API operations for a company repository 

The app is dockerized and along with pgsql and kafka can run under a single docker-compose file.

Zookeeper is used to serve the spawned containers as an orchestrator.

Run docker-compose up --build. 
