[![GitHub Release][release-img]][release]
[![Test][test-img]][test]
[![Go Report Card][go-report-img]][go-report]
[![License: Apache-2.0][license-img]][license]
[![GitHub Downloads][github-downloads-img]][release]
![Docker Pulls][docker-pulls]

[📖 Documentation][docs]
</div>

Tunnel ([pronunciation][pronunciation]) is a comprehensive and versatile security scanner.
Tunnel has *scanners* that look for security issues, and *targets* where it can find those issues.

Targets (what Tunnel can scan):

- Container Image
- Filesystem
- Git Repository (remote)
- Virtual Machine Image
- Kubernetes
- AWS

Scanners (what Tunnel can find there):

- OS packages and software dependencies in use (SBOM)
- Known vulnerabilities (CVEs)
- IaC issues and misconfigurations
- Sensitive information and secrets
- Software licenses

## Quick Start

### Get Tunnel

Tunnel is available in most common distribution methods. The full list of installation options is available in the [Installation] page, here are a few popular options:

- `apt-get install tunnel`
- `yum install tunnel`
- `brew install khulnasoft/tunnel/tunnel`
- `docker run khulnasoft/tunnel`
- Download binary from <https://github.com/khulnasoft/tunnel/releases/latest/>

Tunnel is integrated with many popular platforms and applications. The full list of integrations is available in the [Ecosystem] page. Here are a few popular options:

- [GitHub Actions](https://github.com/khulnasoft/tunnel-action)
- [CircleCI](https://circleci.com/developer/orbs/orb/fifteen5/tunnel-orb)
- [Kubernetes operator](https://github.com/khulnasoft/tunnel-operator)
- [VS Code plugin](https://github.com/khulnasoft/tunnel-vscode-extension)

### General usage

```bash
tunnel <target> [--security-checks <scanner1,scanner2>] <subject>
```

Examples:

```bash
tunnel image python:3.4-alpine
```

<details>
<summary>Result</summary>

https://user-images.githubusercontent.com/1161307/171013513-95f18734-233d-45d3-aaf5-d6aec687db0e.mov

</details>

```bash
tunnel fs --security-checks vuln,secret,config myproject/
```

<details>
<summary>Result</summary>

https://user-images.githubusercontent.com/1161307/171013917-b1f37810-f434-465c-b01a-22de036bd9b3.mov

</details>

```bash
tunnel k8s --report summary cluster
```

<details>
<summary>Result</summary>

![k8s summary](docs/imgs/tunnel-k8s.png)

</details>

## Highlights

- Comprehensive vulnerability detection
    - OS packages (Alpine Linux, Red Hat Universal Base Image, Red Hat Enterprise Linux, CentOS, AlmaLinux, Rocky Linux, CBL-Mariner, Oracle Linux, Debian, Ubuntu, Amazon Linux, openSUSE Leap, SUSE Enterprise Linux, Photon OS and Distroless)
    - **Language-specific packages** (Bundler, Composer, Pipenv, Poetry, npm, yarn, Cargo, NuGet, Maven, and Go)
    - High accuracy, especially [Alpine Linux][alpine] and RHEL/CentOS
- Supply chain security (SBOM support)
    - Support CycloneDX
    - Support SPDX
    - Generating and Scanning SBOM
    - Leveraging in-toto attestations
    - Integrated with [Sigstore]
- Misconfiguration detection (IaC scanning) 
    - Wide variety of security checks are provided **out of the box**
    - Kubernetes, Docker, Terraform, and more
    - User-defined policies using [OPA Rego][rego]
- Secret detection
    - A wide variety of built-in rules are provided **out of the box**
    - User-defined patterns
    - Efficient scanning of container images
- Simple
    - Available in apt, yum, brew, dockerhub
    - **No pre-requisites** such as a database, system libraries, or eny environmental requirements. The binary runs anywhere.
    - The first scan will finish within 10 seconds (depending on your network). Consequent scans will finish instantaneously.
- Fits your workflow
    - **Great for CI** such as GitHub Actions, Jenkins, GitLab CI, etc.
    - Available as extension for IDEs such as vscode, jetbrains, vim
    - Available as extension for Docker Desktop, Rancher Desktop
    - See [Ecosystem] section in the documentation.

## FAQ

### How to pronounce the name "Tunnel"?

`tri` is pronounced like **tri**gger, `vy` is pronounced like en**vy**.

---

Tunnel is an [Khulnasoft Security][khulnasoft] open source project.  
Learn about our open source work and portfolio [here][oss].  
Contact us about any matter by opening a GitHub Discussion [here][discussions]

[test]: https://github.com/khulnasoft/tunnel/actions/workflows/test.yaml
[test-img]: https://github.com/khulnasoft/tunnel/actions/workflows/test.yaml/badge.svg
[go-report]: https://goreportcard.com/report/github.com/khulnasoft/tunnel
[go-report-img]: https://goreportcard.com/badge/github.com/khulnasoft/tunnel
[release]: https://github.com/khulnasoft/tunnel/releases
[release-img]: https://img.shields.io/github/release/khulnasoft/tunnel.svg?logo=github
[github-downloads-img]: https://img.shields.io/github/downloads/khulnasoft/tunnel/total?logo=github
[docker-pulls]: https://img.shields.io/docker/pulls/khulnasoft/tunnel?logo=docker&label=docker%20pulls%20%2F%20tunnel
[license]: https://github.com/khulnasoft/tunnel/blob/main/LICENSE
[license-img]: https://img.shields.io/badge/License-Apache%202.0-blue.svg
[docs]: https://khulnasoft.github.io/tunnel
[pronunciation]: #how-to-pronounce-the-name-tunnel

[Installation]:https://khulnasoft.github.io/tunnel/latest/getting-started/installation/
[Ecosystem]: https://khulnasoft.github.io/tunnel/latestecosystem/tools

[alpine]: https://ariadne.space/2021/06/08/the-vulnerability-remediation-lifecycle-of-alpine-containers/
[rego]: https://www.openpolicyagent.org/docs/latest/#rego
[sigstore]: https://www.sigstore.dev/

[khulnasoft]: https://khulnasoft.com
[oss]: https://www.khulnasoft.com/products/open-source-projects/
[discussions]: https://github.com/khulnasoft/tunnel/discussions
