Here’s the `README.md` file for your main directory. It provides a comprehensive guide for setting up MetalLB using Helm, along with the steps for configuring it with the provided YAML files.

---

## MetalLB Helm Chart Installation

This repository contains the necessary files to deploy and configure MetalLB using Helm in the `metallb-system` namespace.

### Prerequisites

Before installing MetalLB, ensure that:

- You have a running Kubernetes cluster.
- Helm is installed and configured to interact with your cluster.
- MetalLB is required for providing external IPs for services, particularly useful in environments without cloud-based load balancers.

### Step 1: Create the `metallb-system` Namespace

To ensure that MetalLB operates in the correct namespace, first create the `metallb-system` namespace.

```bash
kubectl create namespace metallb-system
```

### Step 2: Install MetalLB Using Helm

You can install MetalLB via Helm by running the following command:

```bash
helm install metallb metallb/ \
    --namespace metallb-system \
    -f my-values.yaml
```

- `metallb/` refers to the MetalLB Helm chart. Make sure to either add the correct repository or reference your local chart.
- `-f my-values.yaml` is used to provide custom values for the MetalLB installation.

### Step 3: Configure MetalLB with IP Address Pool

Once MetalLB is installed, you need to apply the IP address pool and Layer 2 advertisement configuration from the `metallb-config.yaml` file.

```bash
kubectl apply -f metallb-config.yaml
```

This configuration defines a range of IP addresses that MetalLB will assign to services of type `LoadBalancer`.

**Example `metallb-config.yaml`:**

```yaml
# metallb-config.yaml
apiVersion: metallb.io/v1beta1
kind: IPAddressPool
metadata:
  namespace: metallb-system
  name: my-ip-pool
spec:
  addresses:
    - 172.18.18.100-172.18.18.110 # Range of IP addresses

---
apiVersion: metallb.io/v1beta1
kind: L2Advertisement
metadata:
  namespace: metallb-system
  name: advert
spec:
  ipAddressPools:
    - my-ip-pool
```

This YAML creates:



- An `IPAddressPool` named `my-ip-pool` that reserves a range of IP addresses (`172.18.18.100-172.18.18.110`). Ensure that this IP address range is outside of any existing IP ranges used by your Docker gateway or other network components. You can verify your Docker network gateway by running:

    ```bash
    docker network inspect bridge
    ```
- Make sure the range specified in the `IPAddressPool` does not overlap with Docker's network or other services.

- A Layer 2 (L2) advertisement allowing the services to announce their availability using the pool.

### Step 4: Verify Installation

To verify that MetalLB has been installed and configured properly, you can check the status of the pods and resources:

```bash
kubectl get pods -n metallb-system
kubectl get ipaddresspool -n metallb-system
kubectl get l2advertisement -n metallb-system
```

### Customization

You can further customize the installation by editing the `my-values.yaml` file, which contains the Helm values specific to your deployment. The default values are in `metallb/values.yaml`.

To learn more about the Helm chart's values and options, review the documentation in the chart's `README.md` located in the `metallb/` folder.

### Cleanup

If you need to uninstall MetalLB and clean up the resources, use the following command:

```bash
helm uninstall metallb -n metallb-system
kubectl delete namespace metallb-system
```

---

### Troubleshooting

- Ensure that the IP range specified in `metallb-config.yaml` does not overlap with other resources in your network.
- Check the logs of the MetalLB controller and speaker for any errors:

  ```bash
  kubectl logs -n metallb-system deploy/metallb-controller
  kubectl logs -n metallb-system daemonset/metallb-speaker
  ```
