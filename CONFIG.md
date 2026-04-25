# Configuration Reference

## Overview

Litepod can be configured through:
1. **CLI flags** - For immediate overrides
2. **.litepod.yml** - Project-level configuration
3. **Environment variables** - For system-level settings
4. **Global config** - User-level defaults (~/.litepod/config.yml)

Priority order (highest to lowest): CLI flags > .litepod.yml > Environment variables > Global config

## CLI Flags

### `litepod up` Command

```bash
litepod up [PATH] [OPTIONS]
```

**Options**

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--name` | string | auto-generated | Pod name |
| `--image` | string | ubuntu:22.04 | Docker image to use |
| `--port` | int | 3000 | Port for code-server |
| `--volume` | string | | Volume mount (can be repeated) |
| `--env` | string | | Environment variable (can be repeated) |
| `--cpu` | float | 2 | CPU cores to allocate |
| `--memory` | string | 2G | Memory to allocate |
| `--gpu` | bool | false | Enable GPU support |
| `--build` | bool | false | Build image from Dockerfile |
| `--persist` | bool | true | Enable persistent storage |
| `--background` | bool | false | Run in background |
| `--no-interactive` | bool | false | Non-interactive mode |
| `--help` | | | Show help |

**Examples**

```bash
# Basic
litepod up ./my-project

# With custom image
litepod up ./my-project --image node:18

# Multiple volumes
litepod up ./project \
  --volume /home/user/projects:/workspace \
  --volume /tmp/cache:/cache

# Set environment variables
litepod up ./project \
  --env NODE_ENV=development \
  --env DEBUG=true

# Specify resources
litepod up ./project \
  --cpu 4 \
  --memory 8G \
  --gpu true

# Custom port and name
litepod up ./project \
  --name my-dev-env \
  --port 4000
```

### Other Commands

```bash
# List pods
litepod list [OPTIONS]
  --running  | Show only running pods
  --all      | Show all pods including stopped

# Stop pod
litepod down <POD_NAME> [OPTIONS]
  --force    | Force stop without graceful shutdown
  --all      | Stop all pods

# Restart pod
litepod restart <POD_NAME>

# View pod logs
litepod logs <POD_NAME> [OPTIONS]
  --follow   | Follow log output (like tail -f)
  --lines N  | Show last N lines

# Execute command in pod
litepod exec <POD_NAME> <COMMAND>

# System information
litepod doctor
litepod version
```

## .litepod.yml Configuration File

Place this file in your project root for persistent configuration.

### Basic Structure

```yaml
# Project metadata
name: my-project
description: My development environment
version: 1.0.0

# Docker configuration
image: node:18-bullseye
dockerfile: ./Dockerfile  # Alternative: build from custom Dockerfile

# Port configuration
port: 3000
expose:
  - 3000
  - 8000  # For additional services

# Volume mounts
volume:
  - /home/user/projects:/workspace
  - ./data:/app/data
  - /tmp/cache:/cache:ro  # Read-only

# Environment variables
environment:
  NODE_ENV: development
  DEBUG: "true"
  API_URL: http://localhost:8000

# Resource limits
resources:
  cpu: 2
  memory: 4G
  disk: 10G
  gpu:
    enabled: false
    type: nvidia
    count: 1

# Persistence settings
persistence:
  enabled: true
  path: /var/lib/litepod/volumes
  backup:
    enabled: false
    interval: daily
    retention: 7  # days

# Network configuration
network:
  host_mode: false
  dns:
    - 8.8.8.8
    - 1.1.1.1

# Security
security:
  privileged: false
  capabilities:
    - NET_BIND_SERVICE

# Startup commands
startup:
  - npm install
  - npm run build

# Health check
healthcheck:
  command: curl -f http://localhost:3000
  interval: 30s
  timeout: 10s
  retries: 3
  start_period: 10s

# Lifecycle hooks
hooks:
  on_create: echo "Pod created"
  on_start: echo "Pod started"
  on_stop: echo "Pod stopping"
  on_destroy: echo "Pod destroyed"
```

### Complete Example

```yaml
name: full-stack-dev
description: Full-stack development environment

image: ubuntu:22.04

port: 3000

expose:
  - 3000  # Frontend
  - 5432 # PostgreSQL
  - 6379 # Redis

volume:
  - ./src:/app/src
  - ./public:/app/public
  - /app/node_modules  # Anonymous volume

environment:
  NODE_ENV: development
  DATABASE_URL: postgresql://user:pass@localhost/db
  REDIS_URL: redis://localhost:6379
  DEBUG: "true"

resources:
  cpu: 4
  memory: 8G
  disk: 20G

persistence:
  enabled: true
  path: /data

network:
  host_mode: false

startup:
  - apt-get update
  - apt-get install -y npm postgresql-client redis-tools
  - npm install
  - npm run setup

healthcheck:
  command: curl -f http://localhost:3000
  interval: 30s
```

## Environment Variables

### System Environment Variables

| Variable | Type | Example | Description |
|----------|------|---------|-------------|
| `LITEPOD_HOME` | path | ~/.litepod | Litepod config directory |
| `LITEPOD_DOCKER_HOST` | string | unix:///var/run/docker.sock | Docker daemon socket |
| `LITEPOD_DATA_DIR` | path | /var/lib/litepod | Pod data directory |
| `LITEPOD_LOG_LEVEL` | string | info | Log level (debug/info/warn/error) |
| `LITEPOD_LOG_FORMAT` | string | json | Log format (json/text) |
| `LITEPOD_PORT` | int | 3000 | Default port |
| `LITEPOD_BIND_ADDRESS` | string | 0.0.0.0 | Bind address |

### Pod Environment Variables

In `.litepod.yml` or via `--env`:

```yaml
environment:
  # Your variables here
  MY_VAR: value
```

## Global Configuration

Located at `~/.litepod/config.yml`:

```yaml
# Default settings
defaults:
  image: ubuntu:22.04
  port: 3000
  cpu: 2
  memory: 2G

# Behavior
behavior:
  auto_start: false
  auto_cleanup: false
  confirm_delete: true

# Docker
docker:
  registry: docker.io
  insecure_registries: []
  auth:
    username: ""
    password: ""

# Logging
logging:
  level: info
  format: json
  file: ~/.litepod/litepod.log

# Networking
networking:
  nameservers:
    - 8.8.8.8
    - 1.1.1.1
```

## Configuration Precedence

When the same setting exists in multiple places:

```
CLI Flag
    ↓
.litepod.yml (project)
    ↓
Environment Variable
    ↓
~/.litepod/config.yml (user)
    ↓
Default Value
```

## Advanced Configuration

### Custom Docker Images

Create a `Dockerfile`:

```dockerfile
FROM node:18-bullseye

# Install additional tools
RUN apt-get update && apt-get install -y \
    postgresql-client \
    redis-tools \
    git

# Set working directory
WORKDIR /app

# Copy files
COPY . .

# Install dependencies
RUN npm install

EXPOSE 3000

CMD ["npm", "start"]
```

Then use in `.litepod.yml`:

```yaml
dockerfile: ./Dockerfile
build:
  context: .
  args:
    NODE_ENV: development
```

### Volume Permissions

```yaml
volume:
  - ./src:/app/src:rw      # Read-write
  - ./config:/etc/conf:ro  # Read-only
  - ./cache:/tmp/cache:Z   # SELinux context
```

### Network Isolation

```yaml
network:
  host_mode: false
  dns:
    - 8.8.8.8
  expose:
    - 3000
```

### Resource Limits

```yaml
resources:
  cpu: 4           # CPU cores
  memory: 8G       # Memory
  disk: 50G        # Disk space
  swap: 2G         # Swap memory
  io_weight: 500   # I/O priority (10-1000)

  # Rate limiting
  network:
    ingress: 1Gbps
    egress: 1Gbps
```

## Common Configurations

### Node.js Development

```yaml
image: node:18-bullseye
port: 3000
environment:
  NODE_ENV: development
startup:
  - npm install
  - npm run dev
```

### Python Development

```yaml
image: python:3.11-slim
port: 8000
startup:
  - pip install -r requirements.txt
  - python manage.py migrate
```

### Go Development

```yaml
image: golang:1.21
port: 8080
startup:
  - go mod download
  - go build -o app .
```

### Database + App

```yaml
image: ubuntu:22.04
port: 3000
volume:
  - ./app:/app
  - ./db:/data/db
startup:
  - apt-get update
  - apt-get install -y postgres redis-server
  - service postgresql start
  - cd /app && npm install
```

## Validation

Validate your configuration:

```bash
litepod config validate
litepod config validate .litepod.yml
```

## More Information

- [CLI Reference](./CLI.md)
- [Examples](./EXAMPLES.md)
- [Troubleshooting](./TROUBLESHOOTING.md)

---

*Need help? Check [GitHub Discussions](https://github.com/Arnel-rah/litepod/discussions)*
