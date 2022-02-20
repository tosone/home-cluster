# home-cluster

## Installation

``` bash
export K0SCTL_VERSION=v0.13.0-beta.2
wget -O k0sctl https://github.com/k0sproject/k0sctl/releases/download/$K0SCTL_VERSION/k0sctl-darwin-x64

k0sctl apply

k0sctl kubeconfig > kubeconfig
```
