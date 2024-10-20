# Golang-ObservabilityStack


## First golden signal: Latency
**Grafana variables:**

1. ingress =  label_values(nginx_ingress_controller_request_duration_seconds_bucket, ingress)  type of query
2. status = success : [2-3].*,error : [4-5].* type custom
3. method = GET,POST,PUT,DELETE  type custom, include all options

```
grafana legend: P99



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

grafana legend: P90



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

grafana legend: P50



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

## Second golden signal: Traffic
**Grafana variables:**

1. ingress =  label_values(nginx_ingress_controller_request_duration_seconds_bucket, ingress)  type of query
2. status = success : [2-3].*,error : [4-5].* type custom
3. method = GET,POST,PUT,DELETE  type custom, include all options

```
grafana legend: Traffic



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

## Third golden signal: Availability (success/total_requests)
**Grafana variables:**


1. ingress =  label_values(nginx_ingress_controller_request_duration_seconds_bucket, ingress)  type of query
2. status = success : [2-3].*,error : [4-5].* type custom
3. method = GET,POST,PUT,DELETE  type custom, include all options

```
grafana legend: ingress



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
            ingress=~"$ingress",
        }[1m]
    )
) by (ingress)
```

## Forth golden signal: Saturation(cpu)
**Grafana variables:**



1. namespace = label_values(container_cpu_usage_seconds_total,container_label_io_kubernetes_pod_namespace)  include all options
2. pod = label_values(container_cpu_usage_seconds_total{container_label_io_kubernetes_pod_namespace=~"$namespace"},container_label_io_kubernetes_pod_name)   include all options

```
grafana legend: {{ container_label_io_kubernetes_container_name }} in {{ container_label_io_kubernetes_pod_name }}



sum(
    rate(
        container_cpu_usage_seconds_total
        {
            container_label_io_kubernetes_pod_namespace=~"$namespace",
            image!=""
        }[1m]
    )
)
by (
    container_label_io_kubernetes_pod_name,
    container_label_io_kubernetes_container_name
) /
sum(
    container_spec_cpu_quota{
        container_label_io_kubernetes_pod_namespace=~"$namespace",
        image!=""
    }
    /
    container_spec_cpu_period{
        container_label_io_kubernetes_pod_namespace=~"$namespace",
        image!=""
    }
)
by(
    container_label_io_kubernetes_pod_name,
    container_label_io_kubernetes_container_name
)
```

## Forth golden signal: Saturation(memory)
**Grafana variables:**



1. namespace = label_values(container_cpu_usage_seconds_total,container_label_io_kubernetes_pod_namespace)  include all options
2. pod = label_values(container_cpu_usage_seconds_total{container_label_io_kubernetes_pod_namespace=~"$namespace"},container_label_io_kubernetes_pod_name)   include all options

```
grafana legend: {{ container_label_io_kubernetes_container_name }} in {{ container_label_io_kubernetes_pod_name  }}




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