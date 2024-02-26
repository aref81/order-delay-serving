# Order Delay Reporting System

## Overview

This README provides instructions for deploying,
running, and using the Order Delay Reporting System,
which is designed to manage the delay reports and their reviews.
<br>
<br>
**We utilized Docker to simplify the deployment and setup process.**

## Prerequisites

ensure you have met the following requirements:

- Docker and Docker Compose are installed on your system.
- You have knowledge of RESTful API interactions.
- cURL, Postman or any other tools that allow you to call web API

## Deployment

To deploy this system, follow these steps:

1. **Clone the repository** <br>
   Clone the source code to your local machine.
2. **Docker Compose** <br>
   Navigate to the root directory which contains the docker-compose.yaml file.
   This file defines the services that make up your application, in this case,
   postgres, rabbitmq, and the main service.
3. **Building and Running Containers** <br>
   Run the following command to build and start the containers in the background:
    ``` bash
   docker-compose up -d --build
   ```
   This will pull the required images, build the service from the Dockerfile, and start the services as defined.
4. **Verifying the Deployment** <br>
   Verify that all the containers are up and running:
     ``` bash
   docker-compose ps
   ```
5. **Stopping the Services** <br>
   When you're done, you can stop the services using:
     ``` bash
   docker-compose down
   ```

## Using the System

The Order Delay Reporting System provides several endpoints for handling delay reports and other related entities.

### Agents Endpoints

- `POST /agents` : This endpoint is used to create a new agent. The request body should contain the agent details in JSON format.
### Reports Endpoints

- `POST /reports` : Report a new delay for an order. Requires details of the delay report in the request body.
- `POST /reports/:agentID` : Retrieve a queued report for a specific agent. The agentID should be provided in the URL.
- `GET /reports` : Get a summary of vendors delay in last week. No request body is required.
- `POST /reports/review` : Submit a review for a delayed order report. Requires the review submission details in the request body.
### Orders Endpoints

- `POST /orders` : Creates a new order in the system.
- `/orders/:id` : Retrieves the details of an order by its unique identifier.
### Trips Endpoints

- `POST /trips` : Creates a new trip in the system.
- `/trips/:id` : Retrieves the details of a trip by its unique identifier.
### Vendors Endpoints

- `POST /vendors` : Creates a new vendor in the system.
- `/vendors/:id` : Retrieves the details of a vendor by its unique identifier.
## Interacting with Endpoints
To interact with these endpoints, you can use tools like curl or Postman.
Here is an example of how to use curl to create a new agent:

```bash
curl -X POST http://localhost:8000/agents \
     -H "Content-Type: application/json" \
     -d '{"name":"John Doe"}'
```

## Configs
In the `docker-compose.yaml` file, the postgres service uses environment variables for the database user,
password, and database name, also you can custom the default config of rabbitMQ Service in the file too.
<br>
the service itself uses a `config.yaml` file attached as a volume to the container in order to
set the desired configs for the service, a sample config us shown below:

```yaml
server:
  port: 8000
  address: localhost

database:
  port: 5432
  address: postgres
  user: postgres
  password: postgres
  dbname: postgres

rabbit:
  host: rabbitmq
  port: 5672
  user: guest
  password: guest
  queue: sf
```



