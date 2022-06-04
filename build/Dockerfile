FROM golang:1.18-bullseye as builder

ARG FLUX_VERSION=0.30.2
ARG K0SCTL_VERSION=v0.13.0
ARG K0S_VERSION=v1.23.6+k0s.2

WORKDIR /app

COPY . .

RUN go build -o /tmp/hc -ldflags "-s -w" ./cmd && \
    wget -q -O /tmp/k0sctl https://github.com/k0sproject/k0sctl/releases/download/${K0SCTL_VERSION}/k0sctl-linux-x64 && \
    wget -q -O flux_0.30.2_linux_amd64.tar.gz https://github.com/fluxcd/flux2/releases/download/v${FLUX_VERSION}/flux_${FLUX_VERSION}_linux_amd64.tar.gz && \
    tar zxvf flux_${FLUX_VERSION}_linux_amd64.tar.gz && \
    mv flux /tmp/flux && \
    rm -rf flux_${FLUX_VERSION}_linux_amd64.tar.gz && \
    wget -q -O /tmp/k0s-airgap-bundle-${K0S_VERSION}-amd64 https://github.com/k0sproject/k0s/releases/download/${K0S_VERSION}/k0s-airgap-bundle-${K0S_VERSION}-amd64 && \
    wget -q -O /tmp/k0s-${K0S_VERSION}-amd64 https://github.com/k0sproject/k0s/releases/download/${K0S_VERSION}/k0s-${K0S_VERSION}-amd64

FROM debian:bullseye-slim

COPY --from=builder /tmp/hc /usr/local/bin/hc
COPY --from=builder /tmp/k0sctl /usr/local/bin/k0sctl
COPY --from=builder /tmp/flux /usr/local/bin/flux

ARG K0S_VERSION=v1.23.6+k0s.2

COPY --from=builder /tmp/k0s-airgap-bundle-${K0S_VERSION}-amd64 /resource/
COPY --from=builder /tmp/k0s-${K0S_VERSION}-amd64 /resource/

WORKDIR /app

COPY ./env.yaml.sample ./env.yaml
COPY ./deploy.yaml ./deploy.yaml

CMD "/bin/bash"
