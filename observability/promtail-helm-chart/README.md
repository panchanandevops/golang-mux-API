# Promtail Helm Chart

This repository contains the Helm chart for deploying Promtail in your Kubernetes cluster. Promtail is an agent that reads log files and forwards them to Loki, a log aggregation system.

## Prerequisites

- Kubernetes 1.16+
- Helm 3.0+
- Loki must be deployed in your cluster, and Promtail should be able to send logs to Loki via the Loki Gateway.

## Installation

To install the Promtail Helm chart, follow these steps:

1. **Create a custom `my-values.yaml`**:

   You can configure Promtail to send logs to your Loki instance by creating a custom `my-values.yaml` file.

   Example `my-values.yaml`:

   ```yaml
   config:
     clients:
       - url: http://my-loki-gateway.monitoring.svc.cluster.local/loki/api/v1/push
         tenant_id: 1
   ```

   Replace `my-loki-gateway.monitoring.svc.cluster.local` with the fully qualified domain name (FQDN) of your Loki Gateway service.

2. **Deploy the Promtail chart**:

   ```bash
   kubectl create namespace monitoring
   helm install promtail ./promtail -f my-values.yaml --namespace monitoring
   ```

This command installs Promtail using the custom values file `my-values.yaml`.

## Managing the Chart

### Upgrade the chart

To upgrade Promtail with modified values, use the following command:

```bash
helm upgrade promtail ./promtail -f my-values.yaml --namespace monitoring
```

### Uninstall the chart

To uninstall Promtail:

```bash
helm uninstall promtail --namespace monitoring
```

This removes the Promtail resources from your Kubernetes cluster.

## Notes

- If your Loki Gateway URL changes, you must update the `my-values.yaml` configuration and redeploy the chart.
- Ensure that your Promtail has the necessary permissions and network access to send logs to the Loki Gateway.

---

This README covers the basic usage of the Promtail Helm chart and the process to configure it for sending logs to a Loki instance. Make sure to adjust the configurations based on your specific environment.
