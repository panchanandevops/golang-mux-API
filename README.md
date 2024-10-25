

# Golang Mux API

The **Golang Mux API** allows for seamless CRUD operations to manage user data, equipped with observability, a streamlined CI/CD pipeline using GitHub Actions, and robust deployments through Argo CD. Below are instructions for interacting with the API, deploying it on Kubernetes, and ensuring continuous integration and monitoring.

---

## API Endpoints and `curl` Commands

### 1. **POST**: Create New Users

- **Endpoint**: `POST /api/go/users`
- **Description**: Adds a new user with name and email attributes.

#### Example Commands:
```sh
# Create user "Alice Smith"
curl -X POST http://go-api.com/api/go/users \
     -H "Content-Type: application/json" \
     -d '{"name": "Alice Smith", "email": "alice.smith@example.com"}'

# Create user "Bob Johnson"
curl -X POST http://go-api.com/api/go/users \
     -H "Content-Type: application/json" \
     -d '{"name": "Bob Johnson", "email": "bob.johnson@example.com"}'

# Create user "Carol White"
curl -X POST http://go-api.com/api/go/users \
     -H "Content-Type: application/json" \
     -d '{"name": "Carol White", "email": "carol.white@example.com"}'
```

### 2. **GET**: Retrieve All Users

- **Endpoint**: `GET /api/go/users`
- **Description**: Retrieves a list of all users.

```sh
curl -X GET http://go-api.com/api/go/users
```

### 3. **PUT**: Update User Information

- **Endpoint**: `PUT /api/go/users/:id`
- **Description**: Updates a specific user's details based on their ID.

```sh
curl -X PUT http://go-api.com/api/go/users/1 \
     -H "Content-Type: application/json" \
     -d '{"name": "Alice Updated", "email": "alice.updated@example.com"}'
```

### 4. **GET**: Retrieve a User by ID

- **Endpoint**: `GET /api/go/users/:id`
- **Description**: Retrieves details of a user by their ID.

```sh
curl -X GET http://go-api.com/api/go/users/2
```

### 5. **DELETE**: Remove a User

- **Endpoint**: `DELETE /api/go/users/:id`
- **Description**: Deletes a specific user by their ID.

```sh
curl -X DELETE http://go-api.com/api/go/users/3
```

---

## CI/CD Pipeline with GitHub Actions

The GitHub Actions pipeline automates the process of building, testing, and deploying the Dockerized API to Docker Hub and Kubernetes.

### Workflow Overview: `build-docker-image.yaml`

- **Triggers**: Runs on tag pushes (e.g., `v1.0.0`) and changes to relevant directories.
- **Workflow Steps**:
  1. **Checkout & Setup**: Pulls the code and configures Docker Buildx.
  2. **Build & Push Docker Image**: Builds the Go API image, tags it, and pushes it to Docker Hub.
  3. **Update Deployment YAML**: Updates the Kubernetes deployment YAML to use the new image tag.
  4. **Create Pull Request**: Opens a PR to review and merge the updated image tag.

### Required GitHub Secrets

- `DOCKER_USERNAME`: Docker Hub username.
- `DOCKER_PASSWORD`: Docker Hub password.
- `PAT_TOKEN`: GitHub Personal Access Token to create pull requests.

---

## Deployment Using Argo CD

Argo CD supports three methods for managing application deployment to Kubernetes:

### Argo CD Applications

1. **Helm-based Deployment**:
   - **Config File**: `helm.yaml`
   - **Source**: `Deploy/go-helm-chart/`
   - **Destination**: `dev` namespace
   - **Sync Policy**: Automated self-healing and pruning

2. **Kubernetes Manifest Deployment**:
   - **Config File**: `k8s-manifesto.yaml`
   - **Source**: `Deploy/k8s/manifests`
   - **Sync Policy**: Self-healing and pruning enabled

3. **Kustomize-based Deployment**:
   - **Config File**: `kustomize.yaml`
   - **Source**: `Deploy/kustomize/overlays/prod`
   - **Sync Policy**: Automatic sync with self-healing and pruning

### Deploying with Argo CD
1. Clone the repository and apply the Argo CD configurations:
   ```bash
   git clone https://github.com/panchanandevops/golang-mux-API.git
   kubectl apply -f path/to/argo-application.yaml
   ```

---

## Observability and Monitoring Stack

This API application is equipped with observability tools using **Grafana**, **Prometheus**, and **Loki** to monitor, collect, and analyze data.

- **Grafana**: Configured to visualize time-series metrics and system health, with dashboard alerts for anomaly detection.
- **Prometheus**: Aggregates and stores time-series data from the application, enabling historical analysis and critical alerts.
- **Loki**: Manages log aggregation, allowing for efficient querying and debugging.

This setup ensures application stability, enabling quick diagnosis of potential issues and a proactive approach to monitoring.

