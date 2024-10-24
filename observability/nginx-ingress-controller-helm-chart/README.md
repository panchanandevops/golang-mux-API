

# NGINX Ingress Controller Helm Chart

The NGINX Ingress Controller is a high-performance load balancer, reverse proxy, and web server that can be used to manage and route external traffic into a Kubernetes cluster.

This Helm chart deploys the NGINX Ingress Controller with configurable options for scaling, resource usage, and additional features such as security, logging, and metrics.

## Prerequisites

- Kubernetes 1.19+ (versions vary by the helm chart)
- Helm 3.0+
- A running Kubernetes cluster with access to create ingress resources.
- (Optional) Cert-manager if you plan to manage SSL/TLS certificates for your ingresses.

## Installation

To install the chart with the release name `my-ingress` in the `ingress-nginx` namespace, use the following commands:

```bash
# Create namespace (if not existing)
kubectl create namespace ingress-nginx


# Install the chart
helm install my-ingress ./ingress-nginx --namespace ingress-nginx
```

You can also install it with custom values:

```bash
helm install my-ingress ingress-nginx/ingress-nginx --namespace ingress-nginx --values values.yaml
```

## UnInstallation

To uninstall/delete the `my-ingress` deployment:

```bash
helm uninstall my-ingress --namespace ingress-nginx
```

The command removes all the Kubernetes components associated with the release, but keeps the configuration history.

## Configuration

The NGINX Ingress Controller Helm Chart is highly configurable. You can define custom settings by passing a custom `values.yaml` file or using the `--set` option with the `helm install` command.

Here are custom values which we are used.

### 1. **Service Configuration**

```yaml
controller:
  service:
    type: LoadBalancer # Change to LoadBalancer
    annotations:
      prometheus.io/scrape: "true"
      prometheus.io/port: "10254"
      prometheus.io/path: "/metrics"
```

- **`type: LoadBalancer`**:
  - The service type is set to `LoadBalancer`, which means the NGINX Ingress Controller service will expose its endpoints using a cloud provider's load balancer (such as AWS ELB, Google Cloud Load Balancer, or Azure Load Balancer). This is useful when you want to expose your service externally with a public IP.
  - If you're running Kubernetes on a cloud provider, the external load balancer will be automatically provisioned to route traffic to the Ingress Controller.

- **Annotations**:
  - These annotations are used to enable Prometheus scraping for metrics from the NGINX Ingress Controller.
  - **`prometheus.io/scrape: "true"`**: This annotation tells Prometheus that metrics should be scraped from this service.
  - **`prometheus.io/port: "10254"`**: This specifies the port where Prometheus can access the NGINX metrics. By default, the NGINX Ingress Controller exposes metrics on port 10254.
  - **`prometheus.io/path: "/metrics"`**: This tells Prometheus to collect metrics from the `/metrics` endpoint of the NGINX Ingress Controller.

These annotations allow Prometheus to automatically detect and scrape metrics from the NGINX Ingress Controller, which can then be visualized in monitoring tools like Grafana.

---

### 2. **Metrics Configuration**

```yaml
  metrics:
    enabled: true
    serviceMonitor:
      enabled: true
      additionalLabels:
        role: monitoring # Label for Prometheus ServiceMonitor
```

- **`metrics.enabled: true`**:
  - This setting enables metrics collection for the NGINX Ingress Controller. With this enabled, the controller will expose various performance and traffic metrics, such as request count, response times, and more.
  - The metrics are generally exposed on the `/metrics` endpoint (default is port `10254`).

- **`serviceMonitor.enabled: true`**:
  - This setting enables the creation of a `ServiceMonitor` resource, which is used by Prometheus Operator to scrape the metrics from the NGINX Ingress Controller.
  - A `ServiceMonitor` is a custom resource that configures Prometheus to monitor services, allowing for seamless integration between the NGINX Ingress Controller and Prometheus.

- **`additionalLabels`**:
  - The **`role: monitoring`** label is added to the ServiceMonitor resource. Labels are key-value pairs used to organize and select Kubernetes resources.
  - This label can be used by Prometheus Operator to find and select the `ServiceMonitor` for scraping metrics from the NGINX Ingress Controller.

---

### 3. **Resources Configuration**

```yaml
  resources:
    limits:
      cpu: 600m # 2 vCPUs
      memory: 600Mi # 1 GiB of memory
    requests:
      cpu: 300m # 1 vCPU for requests
      memory: 300Mi # 512 MiB for requests
```

- **Resources**:
  - Resource limits and requests allow Kubernetes to manage and allocate compute resources (CPU and memory) to your Ingress Controller effectively.

  - **Requests**:
    - **CPU**: `300m` means that 30% of 1 CPU core (or virtual CPU, vCPU) is requested for this container. This is the minimum amount of CPU resources that Kubernetes will allocate to this pod.
    - **Memory**: `300Mi` means that the container is requesting 300 MiB of memory. This is the minimum memory guaranteed to the pod.
    - Pods that request fewer resources are more likely to be scheduled quickly, but if the load increases, they might not have sufficient resources.

  - **Limits**:
    - **CPU**: `600m` means the container can use up to 60% of 1 CPU core (or vCPU). This is the maximum amount of CPU that Kubernetes will allow the container to use.
    - **Memory**: `600Mi` means that the container can use up to 600 MiB of memory. If the container exceeds this memory limit, it may be killed by Kubernetes.

These resource settings ensure that the NGINX Ingress Controller operates within the limits of the allocated CPU and memory, optimizing performance while ensuring resource efficiency.



### Further Customization

If you need to make additional configuration changes, such as enabling SSL/TLS termination, adjusting load balancer settings, or configuring custom ingress classes, you can modify more options in the `values.yaml` file provided with the `ingress-nginx` Helm chart.

