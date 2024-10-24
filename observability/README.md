

# Golang Observability Stack

This repository sets up a complete observability stack to monitor the **Four Golden Signals**—**Latency**, **Traffic**, **Availability**, and **Saturation**—in your Kubernetes cluster using **Prometheus**, **Loki**, **Promtail**, and **Grafana**. 

By following this guide, you will install the required Helm charts and configure Grafana dashboards to monitor these metrics.

---

## Table of Contents
- [Golang Observability Stack](#golang-observability-stack)
  - [Table of Contents](#table-of-contents)
  - [Installation Overview](#installation-overview)
  - [Grafana Variable Configuration](#grafana-variable-configuration)
    - [Variables to Create:](#variables-to-create)
  - [Golden Signals](#golden-signals)
    - [1. Latency](#1-latency)
      - [PromQL for P99 (99th Percentile Latency):](#promql-for-p99-99th-percentile-latency)
      - [PromQL for P90 (90th Percentile Latency):](#promql-for-p90-90th-percentile-latency)
      - [PromQL for P50 (50th Percentile Latency):](#promql-for-p50-50th-percentile-latency)
    - [2. Traffic](#2-traffic)
      - [PromQL for Traffic:](#promql-for-traffic)
    - [3. Availability](#3-availability)
      - [PromQL for Availability:](#promql-for-availability)
    - [4. Saturation (CPU)](#4-saturation-cpu)
      - [PromQL for CPU Saturation:](#promql-for-cpu-saturation)
    - [4. Saturation (Memory)](#4-saturation-memory)
      - [PromQL for Memory Saturation:](#promql-for-memory-saturation)
  - [Helm Chart Installations](#helm-chart-installations)
  - [Accessing Grafana](#accessing-grafana)
  - [Troubleshooting](#troubleshooting)

---

## Installation Overview

To monitor the four golden signals and collect logs from Kubernetes, the stack leverages **Loki** for logs and **Promtail** for log collection. You will install all the required Helm charts found in this directory, and each sub-directory contains specific instructions on how to install its components.

**Before starting:**

1. Ensure that Helm is installed on your system.
2. Follow the installation steps for each subdirectory within the `observability/` directory, ensuring all components are installed and running properly.

---

## Grafana Variable Configuration

Before diving into monitoring, you will need to configure **Grafana Variables** to help build your custom dashboards and PromQL expressions. 

1. **Create a new dashboard** in Grafana and name it `Four Golden Signals`.
2. Go to the **Variables** section within the dashboard and create the following variables:

### Variables to Create:

1. **ingress**
    - **Type**: Query
    - **Query**: `label_values(nginx_ingress_controller_request_duration_seconds_bucket, ingress)`
    - **Description**: Lists all available ingress controllers.
   
2. **status**
    - **Type**: Custom
    - **Values**: `success : [2-3].*, error : [4-5].*`
    - **Description**: Represents success (2xx-3xx) and error (4xx-5xx) status codes.
   
3. **method**
    - **Type**: Custom
    - **Values**: `GET, POST, PUT, DELETE`
    - **Description**: HTTP methods available for requests.
   
4. **namespace**
    - **Type**: Query
    - **Query**: `label_values(container_cpu_usage_seconds_total, container_label_io_kubernetes_pod_namespace)`
    - **Description**: Lists all namespaces in the cluster.
   
5. **pod**
    - **Type**: Query
    - **Query**: `label_values(container_cpu_usage_seconds_total{container_label_io_kubernetes_pod_namespace=~"$namespace"}, container_label_io_kubernetes_pod_name)`
    - **Description**: Lists all pods in the selected namespace.

---

## Golden Signals

### 1. Latency

Latency measures the time it takes for a request to be processed. Here, we use the **P99**, **P90**, and **P50** quantiles for latency measurement.

#### PromQL for P99 (99th Percentile Latency):

```promql
histogram_quantile(
    0.99, sum(
        rate(
            nginx_ingress_controller_request_duration_seconds_bucket{
                ingress=~"$ingress",
                status=~"$status",
                method=~"$method"
            }[1m]
        )
    ) by (le, ingress)
)
```

**Grafana Legend**: `P99`

#### PromQL for P90 (90th Percentile Latency):

```promql
histogram_quantile(
    0.90, sum(
        rate(
            nginx_ingress_controller_request_duration_seconds_bucket{
                ingress=~"$ingress",
                status=~"$status",
                method=~"$method"
            }[1m]
        )
    ) by (le, ingress)
)
```

**Grafana Legend**: `P90`

#### PromQL for P50 (50th Percentile Latency):

```promql
histogram_quantile(
    0.50, sum(
        rate(
            nginx_ingress_controller_request_duration_seconds_bucket{
                ingress=~"$ingress",
                status=~"$status",
                method=~"$method"
            }[1m]
        )
    ) by (le, ingress)
)
```

**Grafana Legend**: `P50`

---

### 2. Traffic

Traffic measures the number of requests received by the ingress over time.

#### PromQL for Traffic:

```promql
round(
    sum(
        irate(
            nginx_ingress_controller_requests{
                ingress=~"$ingress"
            }[1m]
        )
    ) by (ingress), 0.001
)
```

**Grafana Legend**: `Traffic`

---

### 3. Availability

Availability tracks the proportion of successful requests (status codes 2xx-3xx) compared to total requests.

#### PromQL for Availability:

```promql
sum(
    rate(
        nginx_ingress_controller_requests{
            ingress=~"$ingress",
            status!~"[4-5].*"
        }[1m]
    )
) by (ingress) / 
sum(
    rate(
        nginx_ingress_controller_requests{
            ingress=~"$ingress"
        }[1m]
    ) by (ingress)
```

**Grafana Legend**: `ingress`

---

### 4. Saturation (CPU)

Saturation refers to the resource usage in the system. Here, we measure CPU and Memory saturation.

#### PromQL for CPU Saturation:

```promql
sum(
    rate(
        container_cpu_usage_seconds_total{
            container_label_io_kubernetes_pod_namespace=~"$namespace",
            image!=""
        }[1m]
    )
) by (container_label_io_kubernetes_pod_name, container_label_io_kubernetes_container_name) /
sum(
    container_spec_cpu_quota{
        container_label_io_kubernetes_pod_namespace=~"$namespace",
        image!=""
    } / container_spec_cpu_period{
        container_label_io_kubernetes_pod_namespace=~"$namespace",
        image!=""
    }
) by (container_label_io_kubernetes_pod_name, container_label_io_kubernetes_container_name)
```

**Grafana Legend**: `{{ container_label_io_kubernetes_container_name }} in {{ container_label_io_kubernetes_pod_name }}`

---

### 4. Saturation (Memory)

#### PromQL for Memory Saturation:

```promql
container_memory_working_set_bytes{
    container_label_io_kubernetes_pod_namespace=~"$namespace",
    container_label_io_kubernetes_pod_name=~"$pod",
    container_label_io_cri_containerd_kind="container"
} / 
container_spec_memory_limit_bytes{
    container_label_io_kubernetes_pod_namespace=~"$namespace",
    container_label_io_kubernetes_pod_name=~"$pod",
    container_label_io_cri_containerd_kind="container"
}
```

**Grafana Legend**: `{{ container_label_io_kubernetes_container_name }} in {{ container_label_io_kubernetes_pod_name }}`

---

## Helm Chart Installations

The observability stack requires several components to be installed via Helm:

1. **Prometheus**
2. **Loki**
3. **Promtail**
4. **Grafana**
5. **Alertmanager**

Each component has its own directory in the `observability/` folder with specific installation instructions. Ensure that each Helm chart is installed and running properly by cross-checking the pods and services:

```bash
kubectl get pods --namespace monitoring
kubectl get svc --namespace monitoring
```

---

## Accessing Grafana

Once everything is deployed:

1. Grafana will be accessible through the Ingress (or LoadBalancer) configured.
2. Default login credentials for Grafana:
   - **Username**: `admin`
   - **Password**: `prom-operator`

---

## Troubleshooting

If you encounter any issues with the observability stack:

- **Grafana Dashboards**: Ensure that the PromQL queries and variables are correctly defined.
- **Helm Chart Failures**: Verify that all Helm releases are successfully installed using `helm list --namespace monitoring`.
- **No Data in Grafana**: Check Prometheus and Loki logs to ensure they are scraping metrics and logs correctly.
