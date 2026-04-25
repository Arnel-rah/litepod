# 🚀 Litepod

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![GitHub Stars](https://img.shields.io/github/stars/Arnel-rah/litepod?style=social)](https://github.com/Arnel-rah/litepod)
[![Docker](https://img.shields.io/badge/Docker-Powered-2496ED?logo=docker)](https://www.docker.com/)

A lightweight, open-source orchestrator to spin up ephemeral remote development environments in seconds using Docker. Work from anywhere on any device while maintaining full control of your infrastructure.

## 🎯 Why Litepod?

Modern development demands consistent, reproducible environments—yet cloud-based IDE services often come with hefty price tags and restrictive data policies. **Litepod** changes that.

Instead of expensive SaaS solutions, **Litepod** empowers developers to:

- 🏠 **Host their own** development pods on personal infrastructure (VPS, Raspberry Pi, local Linux machines)
- 🔐 **Keep full control** of their code, data, and resources
- 🌍 **Work from anywhere** with just a web browser
- 💰 **Reduce costs** dramatically compared to traditional cloud IDEs
- ⚡ **Set up in seconds** instead of hours

## ✨ Key Features

| Feature | Benefit |
|---------|---------|
| **One-Click Setup** | Spin up a complete dev environment with a single CLI command |
| **Persistent Storage** | Your files and configurations remain safe between sessions |
| **Custom Environments** | Pre-configured Docker images for Go, Rust, Node.js, Python, and more |
| **Browser-Ready IDE** | Seamless VS Code integration via `code-server` |
| **Lightweight** | Uses resource-efficient Docker containers, not heavy VMs |
| **Multi-User Ready** | Host multiple development pods simultaneously |
| **SSH Access** | Direct terminal access for power users |
| **Zero Setup Overhead** | No complex configuration required to get started |

## 🚀 Quick Start

### Prerequisites

- Docker (v20.10+)
- Docker Compose (v2.0+)
- Linux-based host (WSL2 on Windows, native on macOS with Lima)

### Installation

**Via Binary (Recommended)**

```bash
# Download the latest release
curl -sSL https://github.com/Arnel-rah/litepod/releases/latest/download/litepod-linux-x86_64 \
  -o litepod && chmod +x litepod && sudo mv litepod /usr/local/bin/
```

**Via Source**

```bash
git clone https://github.com/Arnel-rah/litepod.git
cd litepod
make build
sudo mv ./dist/litepod /usr/local/bin/
```

**Via Package Manager**

```bash
# Homebrew (macOS)
brew tap Arnel-rah/litepod
brew install litepod

# Arch Linux
yay -S litepod

# Ubuntu/Debian
sudo apt install litepod
```

### Spin Up Your First Environment

```bash
# Basic: Create a development pod from current directory
litepod up

# Advanced: Specify project path and custom image
litepod up ./my-project --image golang:latest

# List running pods
litepod list

# Connect via browser
open http://localhost:3000

# Stop a pod
litepod down pod-name

# Clean up all pods
litepod down --all
```

## 📋 Usage Examples

### Node.js Project

```bash
litepod up ./my-app --image node:18-bullseye
```

### Rust Development

```bash
litepod up ./rust-project --image rust:latest
```

### Python Data Science

```bash
litepod up ./data-analysis --image python:3.11-slim --resources cpu=2,memory=4G
```

### Full Configuration

```bash
litepod up ./project \
  --name my-dev-env \
  --image ubuntu:22.04 \
  --port 3000 \
  --volume /home/user/projects:/workspace \
  --env NODE_ENV=development \
  --resources cpu=4,memory=8G
```

## 🛠️ Configuration

Create a `.litepod.yml` in your project root for persistent settings:

```yaml
name: my-project
image: node:18-bullseye
port: 3000

volume:
  - /home/user/projects:/workspace

environment:
  NODE_ENV: development
  DEBUG: true

resources:
  cpu: 2
  memory: 4G

persistence:
  enabled: true
  path: /var/lib/litepod/volumes
```

## 📚 Documentation

- [Installation Guide](./docs/INSTALL.md)
- [Configuration Reference](./docs/CONFIG.md)
- [Custom Docker Images](./docs/CUSTOM_IMAGES.md)
- [Networking & Security](./docs/SECURITY.md)
- [Troubleshooting](./docs/TROUBLESHOOTING.md)
- [API Reference](./docs/API.md)
- [CLI Reference](./docs/CLI.md)

## 🌟 Comparison

| Feature | Litepod | Cloud IDE | Local Dev |
|---------|---------|-----------|-----------|
| **Cost** | 🟢 Free | 🔴 $$$/month | 🟢 Free |
| **Data Privacy** | 🟢 Full Control | 🔴 Third-party | 🟢 Local |
| **Setup Time** | 🟢 <30 seconds | 🟢 <1 minute | 🔴 Hours |
| **Device Flexibility** | 🟢 Browser-based | 🟢 Browser-based | 🔴 Fixed machine |
| **Resource Control** | 🟢 Full | 🔴 Limited | 🟢 Full |
| **Collaboration** | 🟡 Partial | 🟢 Built-in | 🔴 Complex |
| **Infrastructure** | 🟢 Self-hosted | 🔴 Vendor-locked | 🟢 Local |

## 🎓 Supported Languages & Frameworks

Litepod works with any Docker image, but comes with optimized templates for:

- **Go** - Go 1.21+
- **Rust** - Latest stable
- **Node.js** - 16, 18, 20
- **Python** - 3.9, 3.10, 3.11, 3.12
- **Java** - OpenJDK 17, 21
- **PHP** - 8.1, 8.2, 8.3
- **Ruby** - 3.1, 3.2, 3.3
- **C/C++** - GCC, Clang

[View all templates →](./docs/TEMPLATES.md)

## 🔒 Security & Privacy

- **Isolated Containers**: Each pod runs in its own isolated Docker container
- **Network Segmentation**: Pods are isolated by default
- **User Authentication**: Built-in authentication layer
- **Encrypted Communication**: TLS/SSL support
- **No Data Sharing**: Your code never leaves your infrastructure

[Security documentation →](./docs/SECURITY.md)

## 📊 Architecture

```
┌─────────────────────────────────────────────┐
│         Your Server/Machine                 │
├─────────────────────────────────────────────┤
│  Litepod Daemon                             │
│  ├── Pod Manager                            │
│  ├── Docker Integration                     │
│  ├── Volume Manager                         │
│  └── Network Manager                        │
├─────────────────────────────────────────────┤
│  Docker Containers (Dev Pods)               │
│  ├── Pod 1 (code-server + tools)           │
│  ├── Pod 2 (code-server + tools)           │
│  └── Pod N (code-server + tools)           │
└─────────────────────────────────────────────┘
         ↑
    Browser Access (HTTP/HTTPS)
         ↑
    Your Laptop, Tablet, Phone
```

## 🤝 Contributing

We love contributions! Here's how to get started:

1. **Fork** the repository
2. **Create a feature branch** (`git checkout -b feature/amazing-feature`)
3. **Make your changes** and add tests
4. **Commit** with clear messages (`git commit -m 'Add amazing feature'`)
5. **Push** to your fork (`git push origin feature/amazing-feature`)
6. **Open a Pull Request**

### Development Setup

```bash
git clone https://github.com/Arnel-rah/litepod.git
cd litepod
make dev          # Install dependencies and build
make test         # Run test suite
make fmt          # Format code
make lint         # Run linter
```

### Code Standards

- Follow existing code style
- Write tests for new features
- Update documentation accordingly
- Ensure all tests pass before submitting PR

See [CONTRIBUTING.md](./CONTRIBUTING.md) for detailed guidelines.

## 🐛 Issues & Roadmap

- **Report bugs**: [GitHub Issues](https://github.com/Arnel-rah/litepod/issues)
- **View roadmap**: [Roadmap](./ROADMAP.md)
- **Request features**: [Discussions](https://github.com/Arnel-rah/litepod/discussions)

### Planned Features

- [ ] Kubernetes integration
- [ ] Web-based file manager
- [ ] Built-in collaboration tools
- [ ] Automated backups
- [ ] Health monitoring dashboard
- [ ] Multi-region support
- [ ] GPU support for ML workloads

## 🧪 Testing

```bash
# Run all tests
make test

# Run specific test
make test TEST=TestPodCreation

# Run with coverage
make test-coverage
```

## 📝 License

Litepod is licensed under the **MIT License**. See [LICENSE](./LICENSE) for details.

This means you can freely use, modify, and distribute Litepod for any purpose, including commercial projects.

## 💬 Community

- **GitHub Issues**: [Report bugs or request features](https://github.com/Arnel-rah/litepod/issues)
- **GitHub Discussions**: [Ask questions & share ideas](https://github.com/Arnel-rah/litepod/discussions)
- **Discord**: [Join our community](https://discord.gg/example) *(Coming soon)*
- **Twitter**: [@litepod_dev](https://twitter.com/litepod_dev) *(Coming soon)*

## 📖 Getting Help

- 🎯 **First time?** Start with the [Quick Start](#quick-start)
- ❓ **Have questions?** Check [FAQ](./docs/FAQ.md)
- 🐛 **Found a bug?** [Open an issue](https://github.com/Arnel-rah/litepod/issues)
- 💡 **Have ideas?** [Start a discussion](https://github.com/Arnel-rah/litepod/discussions)

## 🙏 Acknowledgments

Litepod is inspired by the amazing work of:

- [code-server](https://github.com/coder/code-server) - VS Code in the browser
- [Docker](https://www.docker.com/) - Container platform
- [Devpod](https://devpod.sh/) - Developer environments as code
- [Podman](https://podman.io/) - Container engine

## 📊 Project Statistics

![GitHub forks](https://img.shields.io/github/forks/Arnel-rah/litepod?style=social)
![GitHub watchers](https://img.shields.io/github/watchers/Arnel-rah/litepod?style=social)
![GitHub issues](https://img.shields.io/github/issues/Arnel-rah/litepod)
![GitHub pull requests](https://img.shields.io/github/issues-pr/Arnel-rah/litepod)

## 🚀 Performance Benchmarks

- **Container startup time**: < 5 seconds
- **IDE responsiveness**: < 100ms latency
- **Memory footprint**: ~300MB per pod
- **Disk usage**: ~2GB per pod (base image)

[View detailed benchmarks →](./docs/BENCHMARKS.md)

---

**Made with ❤️ by the Litepod community**

*Start contributing today: [GitHub](https://github.com/Arnel-rah/litepod)*
