

# kube-prometheus-stack

This Helm chart sets up a comprehensive monitoring solution for Kubernetes clusters. It includes Prometheus, Grafana, and Alertmanager, along with other necessary exporters and services to monitor the health of your cluster.

## Prerequisites

- Kubernetes 1.16+
- Helm 3.0+
- Your local `kube-prometheus-stack` Helm chart is required.
- An ingress controller (e.g., NGINX) set up in your cluster.

## Components

This chart includes the following components:

- **Prometheus**: A monitoring system and time-series database.
- **Alertmanager**: Manages alerts sent by Prometheus.
- **Grafana**: Visualization tool for metrics and dashboards.
- **Prometheus Node Exporter**: Exposes system-level metrics to Prometheus.
- **kube-state-metrics**: Exposes Kubernetes resource metrics to Prometheus.
- **Prometheus Operator**: Manages Prometheus, Alertmanager, and related resources.

## Installing the Chart in the `monitoring` Namespace

To install the Helm chart in the `monitoring` namespace with Ingress enabled for Grafana, use the following command:

```bash
helm install [RELEASE_NAME] ./kube-prometheus-stack --namespace monitoring
```

Replace `[RELEASE_NAME]` with your desired release name for the Helm chart.

If the namespace `monitoring` does not exist, you can create it:

```bash
kubectl create namespace monitoring
```

## Ingress Configuration for Grafana

Ensure that Ingress is configured in your `values.yaml` for Grafana. A sample configuration might look like this:

```yaml
grafana:
  ingress:
    enabled: true
    annotations:
      kubernetes.io/ingress.class: "nginx"
    hosts:
      - grafana.example.com
    paths:
      - /
    tls:
      - secretName: grafana-tls
        hosts:
          - grafana.example.com
```

Make sure to replace `grafana.example.com` with your domain name, and ensure your Ingress controller is correctly set up to handle the traffic.

## Uninstalling the Chart

To uninstall/delete the `kube-prometheus-stack` deployment:

```bash
helm uninstall [RELEASE_NAME] --namespace monitoring
```

## Customizing Your Deployment

You can customize the chart settings by modifying the `values.yaml` file in the Helm chart directory. To use a custom values file, use the `-f` flag during the install:

```bash
helm install [RELEASE_NAME] ./kube-prometheus-stack --namespace monitoring -f custom-values.yaml
```

### Key Configuration Parameters

- **Prometheus**: Adjust resources, service monitors, and scrape intervals.
- **Grafana**: Configure dashboards, datasources, and user authentication.
- **Alertmanager**: Manage alerting configurations and integrations.
- **Node Exporter**: Adjust resources for system-level metric collection.
- **kube-state-metrics**: Fine-tune settings for collecting Kubernetes object metrics.

## Example: Basic Installation with Ingress

To install with default settings and Ingress enabled for Grafana:

```bash
helm install my-monitoring ./kube-prometheus-stack --namespace monitoring -f custom-values.yaml
```

This will set up Prometheus, Grafana (with Ingress), Alertmanager, Node Exporter, and `kube-state-metrics` with your custom configurations.

## Accessing Grafana

Once the chart is deployed and Ingress is enabled, you can access Grafana through your configured domain. For example:

- **Grafana URL**: `http://monitoring.com`

Ensure that DNS is properly configured to point to your Ingress controller, and the necessary TLS certificates are provided if using HTTPS.

- **Default Username**: `admin`
- **Default Password**: `prom-operator` (you can change this in `values.yaml`)

## Monitoring Kubernetes

This chart automatically monitors the following components of your Kubernetes cluster:

- Nodes
- Pods
- Deployments
- Services
- Persistent Volumes
- Ingress Controllers

Additional custom application monitoring can be set up using **ServiceMonitor** or **PodMonitor** resources, which Prometheus uses to scrape custom metrics.

## Upgrading the Chart

To upgrade your Helm chart with any new changes made to `values.yaml` or other resources:

```bash
helm upgrade [RELEASE_NAME] ./kube-prometheus-stack --namespace monitoring -f custom-values.yaml
```

