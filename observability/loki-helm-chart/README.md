
# Loki Helm Chart

This repository contains the Helm chart for deploying Loki, a log aggregation system designed for efficiently storing and querying logs from Kubernetes.



## Charts

- **Loki:** The main chart for deploying Loki, either in single-binary mode or as a scalable distributed system.
- **Grafana Agent Operator:** Used for managing Grafana agents and collecting logs.
- **Minio:** A dependency for providing object storage when using Loki in certain environments.

## Templates

Templates are defined for various components of Loki including:

- `admin-api`
- `backend`
- `compactor`
- `distributor`
- `ingester`
- `gateway`
- `index-gateway`

Each component has Kubernetes resources like deployments, services, and HPA configurations.

## Values Files

There are different `values.yaml` files provided for various deployment scenarios:

- **distributed-values.yaml:** For deploying Loki in a distributed mode.
- **simple-scalable-values.yaml:** For scalable but simple deployments.
- **single-binary-values.yaml:** For deploying Loki as a single binary, useful in small environments.

