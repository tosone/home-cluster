# home-cluster

## Installation

### Download prepared package

``` bash
export K0S_VERSION=v1.23.6+k0s.2
wget -c -O k0s-airgap-bundle-$K0S_VERSION-amd64  https://github.com/k0sproject/k0s/releases/download/$K0S_VERSION/k0s-airgap-bundle-$K0S_VERSION-amd64
wget -c -O k0s-$K0S_VERSION-amd64 https://github.com/k0sproject/k0s/releases/download/$K0S_VERSION/k0s-$K0S_VERSION-amd64
```

### Download k0sctl

``` bash
export K0SCTL_VERSION=v0.13.0-rc.3
wget -O k0sctl https://github.com/k0sproject/k0sctl/releases/download/$K0SCTL_VERSION/k0sctl-darwin-x64
chmod +x k0sctl

k0sctl apply

k0sctl kubeconfig > kubeconfig
```

### Create helm release with flux

``` bash
# create redis
flux create source helm bitnami --namespace=dev --url=https://charts.bitnami.com/bitnami
flux create helmrelease redis --namespace=dev --source=HelmRepository/bitnami --chart=redis

# create local-path-provisioner
flux create source git local-path-provisioner --namespace=dev --url=https://github.com/rancher/local-path-provisioner.git --tag=v0.0.22
flux create helmrelease local-path-provisioner --namespace=dev --source=GitRepository/local-path-provisioner --chart=./deploy/chart/
```
