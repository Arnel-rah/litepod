# Installation Guide

## System Requirements

### Minimum Requirements
- **OS**: Linux (Ubuntu 20.04+, Debian 11+, CentOS 8+), macOS 10.15+, Windows with WSL2
- **CPU**: 2 cores
- **RAM**: 2GB
- **Disk**: 10GB free space
- **Docker**: v20.10 or later
- **Docker Compose**: v2.0 or later (optional, for advanced setups)

### Recommended Requirements
- **OS**: Ubuntu 22.04 LTS or later
- **CPU**: 4+ cores
- **RAM**: 8GB+
- **Disk**: 50GB+ free space
- **Network**: 100Mbps+ connection

## Installation Methods

### 1. Binary Installation (Linux/macOS)

**Latest Release**

```bash
# Download the latest binary
curl -sSL https://github.com/Arnel-rah/litepod/releases/latest/download/litepod-linux-x86_64 \
  -o litepod

# Make it executable
chmod +x litepod

# Move to PATH
sudo mv litepod /usr/local/bin/
```

**Specific Version**

```bash
VERSION=v0.1.0
curl -sSL https://github.com/Arnel-rah/litepod/releases/download/$VERSION/litepod-linux-x86_64 \
  -o litepod && chmod +x litepod && sudo mv litepod /usr/local/bin/
```

**Verify Installation**

```bash
litepod version
litepod help
```

### 2. Building from Source

**Clone Repository**

```bash
git clone https://github.com/Arnel-rah/litepod.git
cd litepod
```

**Install Build Dependencies**

```bash
# Ubuntu/Debian
sudo apt-get install build-essential git python3 python3-pip

# macOS (with Homebrew)
brew install python3

# Fedora/CentOS
sudo dnf install gcc git python3 python3-devel
```

**Build**

```bash
# Install and build
make install

# Or just build
make build

# Binary will be in ./dist/
./dist/litepod version
```

**Install to System**

```bash
sudo make install-system
# or
sudo mv ./dist/litepod /usr/local/bin/
```

### 3. Package Manager Installation

#### Homebrew (macOS)

```bash
brew tap Arnel-rah/litepod
brew install litepod
```

To upgrade:
```bash
brew upgrade litepod
```

#### AUR (Arch Linux)

```bash
yay -S litepod
# or
yay -S litepod-bin  # Pre-compiled binary
```

#### Ubuntu/Debian PPA (Coming Soon)

```bash
sudo add-apt-repository ppa:arnel-rah/litepod
sudo apt update
sudo apt install litepod
```

### 4. Docker Installation

**Pull Docker Image**

```bash
docker pull arnelrah/litepod:latest
```

**Run as Docker Container**

```bash
# Basic
docker run -it \
  -v /var/run/docker.sock:/var/run/docker.sock \
  arnelrah/litepod:latest litepod help

# With persistent storage
docker run -d \
  --name litepod \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v /home/user/.litepod:/home/litepod/.litepod \
  -p 3000:3000 \
  arnelrah/litepod:latest
```

### 5. Kubernetes Installation (Coming Soon)

```bash
# Using Helm
helm repo add litepod https://charts.litepod.dev
helm repo update
helm install litepod litepod/litepod
```

## Verifying Installation

### Check Version

```bash
litepod version
```

Expected output:
```
Litepod version v0.1.0
Build date: 2024-04-25
Git commit: abc123def456
```

### Check Docker Integration

```bash
litepod doctor
```

This will verify:
- Docker installation
- Docker daemon is running
- Docker permissions
- Available disk space
- System resources

### Run First Pod

```bash
# Create a test directory
mkdir -p ~/litepod-test
cd ~/litepod-test

# Start a pod
litepod up

# List pods
litepod list

# Access the IDE
open http://localhost:3000
```

## Uninstallation

### Remove Binary

```bash
# If installed via package manager
sudo apt remove litepod          # Ubuntu/Debian
yay -R litepod                   # Arch
brew uninstall litepod           # macOS

# If installed manually
sudo rm /usr/local/bin/litepod
```

### Clean Up Data

```bash
# Remove Litepod configuration and data
rm -rf ~/.litepod

# Remove Docker volumes (caution: this deletes data)
docker volume prune -f

# Remove Litepod Docker images
docker rmi arnelrah/litepod
docker rmi arnelrah/litepod-base
```

## Troubleshooting

### Docker Permission Denied

If you get `permission denied while trying to connect to Docker daemon`:

```bash
# Add your user to docker group
sudo usermod -aG docker $USER
newgrp docker

# Or use sudo
sudo litepod up
```

### Insufficient Disk Space

```bash
# Check available space
df -h

# Clean up Docker
docker system prune -a

# Then try again
litepod up
```

### Port Already in Use

If port 3000 is already in use:

```bash
litepod up --port 3001
```

Or find and stop the process:

```bash
lsof -i :3000
kill -9 <PID>
```

### Network Issues

```bash
# Test Docker network
docker run --rm alpine ping google.com

# Check DNS
docker run --rm alpine nslookup github.com

# Restart Docker daemon
sudo systemctl restart docker
```

### Installation on WSL2 (Windows)

1. **Install WSL2**
   ```powershell
   wsl --install
   ```

2. **Install Ubuntu**
   ```powershell
   wsl --install -d Ubuntu-22.04
   ```

3. **Inside WSL2, install Docker**
   ```bash
   curl -fsSL https://get.docker.com -o get-docker.sh
   sh get-docker.sh
   ```

4. **Install Litepod**
   ```bash
   curl -sSL https://github.com/Arnel-rah/litepod/releases/latest/download/litepod-linux-x86_64 \
     -o litepod && chmod +x litepod && sudo mv litepod /usr/local/bin/
   ```

## Next Steps

1. Read the [Quick Start](../README.md#quick-start) guide
2. Check out [Configuration](./CONFIG.md) options
3. Explore [Examples](./EXAMPLES.md)
4. Join our [Community](https://github.com/Arnel-rah/litepod/discussions)

## Getting Help

- 📖 Check the [Troubleshooting](./TROUBLESHOOTING.md) guide
- 💬 Ask on [GitHub Discussions](https://github.com/Arnel-rah/litepod/discussions)
- 🐛 Report issues on [GitHub Issues](https://github.com/Arnel-rah/litepod/issues)
- 📧 Email: hello@litepod.dev

---

**Happy coding with Litepod!** 🚀
