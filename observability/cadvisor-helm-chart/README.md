# cAdvisor Helm Chart

### Overview

This Helm chart is designed to deploy **cAdvisor** as a DaemonSet for monitoring containerized applications across your Kubernetes nodes. cAdvisor provides container resource usage and performance metrics. The chart also integrates with Prometheus for scraping these metrics, making it a key part of your Kubernetes monitoring stack.


### Prerequisites

1. **Kubernetes cluster** (v1.18+)
2. **Helm** (v3+)
3. **Prometheus Operator** for scraping metrics
4. A **monitoring** namespace (or create it)

### Install Guide

Follow these steps to install the cAdvisor Helm chart in the `monitoring` namespace.

#### Step 1: Create the Monitoring Namespace

Ensure the `monitoring` namespace exists in your cluster:

```bash
kubectl create namespace monitoring
```

#### Step 2: Customize Values (Optional)

You can customize the chart using the provided `my-values.yaml` file. Below is an example of how to enable metrics and add custom labels.

**`my-values.yaml`**:
```yaml
metrics:
  enabled: true
  customLabels:
    role: monitoring
```

This configuration ensures that Prometheus scrapes cAdvisor metrics and adds a `role=monitoring` label for tracking.

#### Step 3: Install the Chart

Install the chart using the following Helm command:

```bash
helm install cadvisor ./cadvisor -f my-values.yaml --namespace monitoring
```

- `cadvisor` is the release name.
- `./cadvisor` is the path to the Helm chart directory.
- `-f my-values.yaml` applies your custom configurations.
- `--namespace monitoring` installs the chart in the `monitoring` namespace.

#### Step 4: Verify Installation

After installation, you can verify that the DaemonSet, Service, and RBAC resources are created successfully:

```bash
kubectl get daemonset -n monitoring
kubectl get service -n monitoring
kubectl get servicemonitor -n monitoring
```

### Key Components

- **DaemonSet**: Deploys cAdvisor on all nodes, ensuring each node is monitored.
- **RBAC**: Grants cAdvisor the required permissions for accessing necessary resources.
- **ServiceMonitor**: Configures Prometheus to scrape metrics from cAdvisor instances.
- **Service**: Exposes cAdvisor's metrics on the specified port.

### Customization

You can further customize the chart by modifying the `values.yaml` file or providing a custom `my-values.yaml`. Key configurations include:

- **Metrics Configuration**:
  - Enable or disable metrics scraping.
  - Add custom labels for tracking and monitoring purposes.

- **RBAC Settings**:
  - Modify RBAC settings such as ClusterRole and ClusterRoleBinding if needed.

### Uninstall the Chart

To remove the cAdvisor Helm release, run:

```bash
helm uninstall cadvisor --namespace monitoring
```

This will remove all associated resources from the `monitoring` namespace.




