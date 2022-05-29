# home-cluster

## Installation

``` bash
export K0SCTL_VERSION=v0.13.0-rc.3
wget -O k0sctl https://github.com/k0sproject/k0sctl/releases/download/$K0SCTL_VERSION/k0sctl-darwin-x64
chmod +x k0sctl

k0sctl apply

k0sctl kubeconfig > kubeconfig

kubectl taint nodes node-xxx node-role.kubernetes.io/master- --kubeconfig kubeconfig

KUBECONFIG=./kubeconfig helmfile sync --skip-deps
```
