# Complete Docker Mastery Guide - Zero to Hero 🐳

Welcome to the most comprehensive Docker guide you'll find! This guide covers everything from basic concepts to advanced techniques, making you a Docker expert ready for any real-world scenario.

## Table of Contents
1. [Why Docker? The Container Revolution](#why-docker)
2. [Understanding the Container Landscape](#container-landscape)
3. [Docker Architecture Deep Dive](#docker-architecture)
4. [Docker Under the Hood](#docker-under-the-hood)
5. [Essential Docker Commands](#essential-commands)
6. [Docker Images Mastery](#docker-images)
7. [Container Lifecycle Management](#container-lifecycle)
8. [Docker Networking Complete Guide](#docker-networking)
9. [Data Persistence with Volumes](#docker-volumes)
10. [Resource Management & Monitoring](#resource-management)
11. [Docker Registries & Distribution](#docker-registries)
12. [Multi-Container Apps with Docker Compose](#docker-compose)
13. [Advanced Dockerfile Techniques](#advanced-dockerfile)
14. [Production Best Practices](#production-best-practices)
15. [Troubleshooting & Debugging](#troubleshooting)
16. [Complete Command Reference](#command-reference)

---

## Why Docker? The Container Revolution {#why-docker}

Docker revolutionized software development and deployment by solving critical problems that plagued developers for decades:

### The Problems Docker Solves

**"It Works on My Machine!" Syndrome**
- Eliminates environment inconsistencies between development, testing, and production
- Ensures your application runs identically across all environments
- Packages the entire runtime environment with your application

**Dependency Hell**
- Isolates applications and their dependencies in separate containers
- No more conflicting library versions or missing dependencies
- Each container has its own isolated environment

**Resource Efficiency**
- Containers share the host OS kernel, making them incredibly lightweight
- Start in seconds, not minutes like virtual machines
- Better resource utilization compared to traditional VMs

**Scalability & DevOps**
- Perfect for microservices architecture
- Enables continuous integration and deployment (CI/CD)
- Makes horizontal scaling simple and efficient
- Supports modern DevOps practices

### Why Container-Based Software Delivery?

Container-based delivery represents the evolution of application deployment:

1. **Consistency**: Same container runs everywhere - development, testing, production
2. **Portability**: Run on any system that supports containers
3. **Efficiency**: Better resource utilization than VMs
4. **Speed**: Rapid deployment and scaling
5. **Isolation**: Applications don't interfere with each other
6. **Microservices**: Perfect for distributed architectures

---

## Understanding the Container Landscape {#container-landscape}

### Bare Metal vs VMs vs Containers

| Aspect | Bare Metal | Virtual Machines | Containers |
|--------|------------|------------------|------------|
| **OS** | Direct on hardware | Each VM has full OS | Share host OS kernel |
| **Size** | N/A | Heavy (GBs) | Lightweight (MBs) |
| **Startup** | Boot time | Slow (minutes) | Fast (seconds) |
| **Isolation** | Hardware level | Complete OS isolation | Process-level isolation |
| **Resources** | Full hardware | High overhead | Minimal overhead |
| **Density** | 1 per server | ~10-50 per server | 100s-1000s per server |
| **Security** | Hardware isolation | Strong isolation | Process isolation |

### Key Differences Explained

**Virtual Machines:**
- Include a complete operating system
- Require hypervisor layer
- Heavy resource consumption
- Strong isolation but at a cost

**Containers:**
- Share the host kernel
- No hypervisor needed
- Lightweight and efficient
- Process-level isolation
- **Note**: Containers don't have their own kernel - they share the host's kernel

---

## Docker Architecture Deep Dive {#docker-architecture}

Docker follows a client-server architecture with several key components:

### Core Components

**Docker Engine**
- The heart of Docker, consisting of:
  - **dockerd (Docker Daemon)**: Background service managing containers, images, networks, volumes
  - **Docker API**: RESTful API for communication
  - **Docker CLI**: Command-line interface for user interaction

**Docker Client**
- Command-line tool (`docker` command)
- Communicates with Docker daemon via REST API
- Can connect to remote Docker daemons

**Docker Registry**
- Stores and distributes Docker images
- Docker Hub is the default public registry
- Can be private (Docker Registry, Harbor, etc.)

**containerd**
- Container runtime that actually runs containers
- Manages container lifecycle
- Interface between Docker daemon and kernel

### Architecture Commands

```bash
# Get comprehensive system information
docker system info

# Monitor real-time events from Docker daemon
docker system events

# Check Docker system resource usage
docker system df

# View detailed breakdown of space usage
docker system df -v
```

---

## Docker Under the Hood {#docker-under-the-hood}

Understanding how Docker works internally helps you use it more effectively:

### Linux Namespaces

Namespaces provide isolation for containers. Each container gets its own namespace for:

**Mount Namespace (mnt)**
- Isolates filesystem mount points
- Each container sees its own filesystem tree

**Network Namespace (net)**
- Provides separate network stack
- Own IP addresses, routing tables, network interfaces

**UTS Namespace**
- Isolates hostname and domain name
- Each container can have unique hostname

**Process ID Namespace (pid)**
- Isolates process IDs
- Process ID 1 in container is different from host

**Inter-Process Communication Namespace (ipc)**
- Isolates IPC resources
- Shared memory, semaphores, message queues

**User Namespace**
- Maps user IDs between container and host
- Root in container ≠ root on host

### Control Groups (cgroups)

Control groups manage and limit resource usage:

**CPU Control**
- Limit CPU usage per container
- CPU shares for relative priority

**Memory Control**
- Set memory limits
- Monitor memory usage
- OOM (Out of Memory) killer

**Disk I/O Control**
- Limit read/write operations
- Bandwidth throttling

**Network Control**
- Bandwidth limiting
- Traffic shaping

### Union + Overlay Filesystem

Docker uses layered filesystems for efficiency:

**Image Layers**
- Each instruction in Dockerfile creates a layer
- Layers are cached and reusable
- Only changed layers need to be pulled/pushed

**Copy-on-Write (CoW)**
- Containers share image layers
- Changes create new writable layer
- Efficient storage utilization

**Storage Driver**
- overlay2 is the current default
- Manages how layers are stored and accessed

```bash
# View image layers and history
docker image history <image_name>

# See filesystem changes in container
docker diff <container_name>

# Create image from container changes
docker commit <container_id> <new_name>:<tag>
```

---

## Essential Docker Commands {#essential-commands}

### Image Operations

```bash
# Pull image from registry
docker pull nginx:latest
docker pull ubuntu:20.04

# List all local images
docker images
docker image ls

# Remove images
docker rmi nginx:latest
docker image rm nginx:latest

# Remove dangling images (untagged)
docker image prune

# Remove all unused images
docker image prune -a

# Build image from Dockerfile
docker build -t myapp:1.0 .
docker build -t myapp:1.0 -f custom.dockerfile .
```

### Container Basic Operations

```bash
# Create container without starting
docker create --name my-container nginx

# Run container (create + start)
docker run nginx
docker run -d nginx                    # Detached mode
docker run -it ubuntu bash             # Interactive mode
docker run -idt ubuntu                 # Interactive + detached + tty

# Start/stop containers
docker start my-container
docker stop my-container
docker restart my-container

# Remove containers
docker rm my-container
docker rm -f my-container              # Force remove running container
docker rm container1 container2        # Remove multiple

# Auto-remove after exit
docker run -d --rm nginx
```

### Container Inspection & Monitoring

```bash
# List containers
docker ps                              # Running containers
docker ps -a                          # All containers
docker ps -l                          # Latest container
docker ps -n 2                        # Last 2 containers

# Container logs
docker logs my-container
docker logs -f my-container            # Follow logs
docker logs --tail 100 my-container   # Last 100 lines

# Execute commands in running container
docker exec -it my-container bash
docker exec my-container ls /app

# Inspect container details
docker inspect my-container

# Copy files to/from container
docker cp ./file.txt container:/app/
docker cp container:/app/file.txt ./
```

### Container Lifecycle Control

```bash
# Container states: created → running → stopped → removed

# Pause/unpause containers
docker pause my-container
docker unpause my-container

# Kill container (SIGKILL)
docker kill my-container

# Stop container (SIGTERM → SIGKILL after 10s)
docker stop my-container

# Container cleanup
docker container prune                 # Remove all stopped containers
```

---

## Docker Images Mastery {#docker-images}

### Understanding Docker Images

Images are read-only templates used to create containers. They consist of layers that are cached and reusable.

### Image Operations

```bash
# Search for images on Docker Hub
docker search nginx

# Pull specific version
docker pull nginx:alpine
docker pull nginx:1.21.6

# Tag images
docker tag nginx:latest mynginx:v1.0
docker tag mynginx:v1.0 myregistry.com/mynginx:v1.0

# Push to registry
docker push myregistry.com/mynginx:v1.0

# Save/load images (for backup/transfer)
docker save nginx > nginx.tar
docker load < nginx.tar

# Export/import (creates new image from container)
docker export container > container.tar
docker import container.tar myapp:latest

# Image history and layers
docker image history nginx
docker image inspect nginx
```

### Creating Images from Containers

```bash
# Make changes to running container
docker run -it ubuntu bash
# (make changes inside container)

# Commit changes to new image
docker commit <container_id> myubuntu:v1.0
docker commit -m "Added custom config" <container_id> myubuntu:v1.0

# View changes made to container
docker diff <container_name>
```


---

## Container Lifecycle Management {#container-lifecycle}

### Container States

Containers go through several states in their lifecycle:

1. **Created**: Container exists but not started
2. **Running**: Container is actively running
3. **Paused**: Container processes are paused
4. **Stopped**: Container has exited
5. **Removed**: Container is deleted

### Lifecycle Commands

```bash
# Create container
docker create --name web nginx

# Start existing container
docker start web

# Stop running container (graceful)
docker stop web                       # Sends SIGTERM, waits 10s, then SIGKILL

# Kill container (immediate)
docker kill web                       # Sends SIGKILL immediately

# Restart container
docker restart web

# Pause/unpause
docker pause web
docker unpause web

# Remove container
docker rm web
docker rm -f web                      # Force remove running container

# Update container resources
docker update --memory 512m web
docker update --cpus 2 web
```

### Container Resource Management

```bash
# Monitor resource usage
docker stats
docker stats --no-stream
docker stats web db cache

# Container processes
docker top web

# Container filesystem usage
docker exec web df -h
```

---

## Docker Networking Complete Guide {#docker-networking}

### Network Drivers

Docker provides several network drivers for different use cases:

**Bridge Network (Default)**
- Default network for containers
- Provides internal connectivity
- NAT for external access

**Host Network**
- Container uses host's network directly
- No network isolation
- Better performance

**None Network**
- No networking
- Complete network isolation

**Overlay Network**
- Multi-host networking
- Used in Docker Swarm
- Encrypted by default

**Macvlan Network**
- Assigns MAC address to container
- Appears as physical device on network

### Network Management

```bash
# List networks
docker network ls

# Create custom networks
docker network create mynetwork
docker network create -d bridge mybridge
docker network create -d overlay myoverlay

# Inspect network
docker network inspect mynetwork
docker network inspect bridge

# Connect/disconnect containers
docker network connect mynetwork container1
docker network disconnect mynetwork container1

# Remove networks
docker network rm mynetwork
docker network prune                  # Remove unused networks
```

### Port Mapping & Exposure

```bash
# Port mapping syntax: -p host_port:container_port

# Map specific port
docker run -p 8080:80 nginx

# Map to random host port
docker run -P nginx
docker run -p 80 nginx                # Container port 80 to random host port

# Multiple port mappings
docker run -p 8080:80 -p 8443:443 nginx

# Bind to specific interface
docker run -p 127.0.0.1:8080:80 nginx

# UDP ports
docker run -p 8080:80/udp myapp

# Check port mappings
docker port container_name
```

### Dockerfile Port Exposure

```dockerfile
# EXPOSE in Dockerfile is for documentation
# It doesn't actually publish ports
EXPOSE 80 443

# Still need -p or -P when running
docker run -P myapp                   # -P publishes all EXPOSED ports to random host ports
```

### Network Communication

**Default Bridge Network**
- Containers communicate via IP addresses only
- No automatic DNS resolution between containers

**User-Defined Bridge Networks**
- Containers can communicate using container names as DNS
- Better isolation and security
- Automatic service discovery

```bash
# Create user-defined network
docker network create myapp-net

# Run containers on custom network
docker run -d --name web --network myapp-net nginx
docker run -d --name db --network myapp-net postgres

# Containers can now communicate using names:
# web can reach db using hostname "db"
```

### Advanced Networking

```bash
# Create network with custom subnet
docker network create --subnet=172.20.0.0/16 mynet

# Assign static IP
docker run --network mynet --ip 172.20.0.10 nginx

# Network aliases
docker run --network mynet --network-alias web nginx

# Multiple networks
docker run --network net1 --network net2 myapp
```

---

## Data Persistence with Volumes {#docker-volumes}

Containers are ephemeral - data is lost when containers are removed. Volumes provide persistent storage.

### Volume Types

**Named Volumes (Recommended)**
- Managed by Docker
- Stored in `/var/lib/docker/volumes/`
- Easy to backup and manage

**Bind Mounts**
- Mount host directory/file into container
- Direct access from host
- Good for development

**tmpfs Mounts**
- Store in host memory
- Never written to disk
- Good for sensitive data

### Volume Operations

```bash
# Create named volume
docker volume create mydata
docker volume create --driver local mydata

# List volumes
docker volume ls

# Inspect volume
docker volume inspect mydata

# Remove volumes
docker volume rm mydata
docker volume prune                   # Remove unused volumes

# Volume with custom options
docker volume create --driver local \
  --opt type=nfs \
  --opt o=addr=192.168.1.1,rw \
  --opt device=:/path/to/dir \
  myvolume
```

### Using Volumes with Containers

```bash
# Named volume
docker run -v mydata:/data nginx
docker run -v mydata:/var/lib/mysql mysql

# Bind mount (host path:container path)
docker run -v /host/path:/container/path nginx
docker run -v $(pwd):/app node:alpine

# Read-only mount
docker run -v mydata:/data:ro nginx
docker run -v /host/path:/app:ro nginx

# tmpfs mount
docker run --tmpfs /tmp nginx
```

### Mount Syntax (Alternative to -v)

```bash
# Named volume mount
docker run --mount source=mydata,target=/data nginx

# Bind mount
docker run --mount type=bind,source=/host/path,target=/app nginx

# tmpfs mount
docker run --mount type=tmpfs,target=/tmp nginx

# Read-only mount
docker run --mount source=mydata,target=/data,readonly nginx

# Mount with specific options
docker run --mount type=volume,source=mydata,target=/data,volume-driver=local nginx
```

### Volume Best Practices

```bash
# Backup volume
docker run --rm -v mydata:/data -v $(pwd):/backup alpine tar czf /backup/backup.tar.gz -C /data .

# Restore volume
docker run --rm -v mydata:/data -v $(pwd):/backup alpine tar xzf /backup/backup.tar.gz -C /data

# Share volume between containers
docker run -d --name web -v shared:/data nginx
docker run -d --name worker -v shared:/data worker-image

# Volume from another container
docker run --volumes-from web worker-image
```

---

## Resource Management & Monitoring {#resource-management}

### Monitoring Containers

```bash
# Real-time resource usage
docker stats
docker stats --no-stream=true         # One-time snapshot
docker stats web db                   # Specific containers
docker stats --format "table {{.Container}}\t{{.CPUPerc}}\t{{.MemUsage}}"

# Container processes
docker top container_name

# System-wide Docker resource usage
docker system df
docker system df -v                   # Verbose output
```

### Memory Management

```bash
# Set memory limit
docker run -m 512m nginx
docker run --memory=1g nginx
docker run --memory=1g --memory-swap=2g nginx

# Update running container
docker update -m 1g container_name

# OOM killer disable (not recommended)
docker run --oom-kill-disable -m 512m nginx
```

### CPU Management

```bash
# CPU shares (relative weight)
docker run --cpu-shares=512 nginx     # 512/1024 = 50% relative weight

# CPU limit (absolute)
docker run --cpus="1.5" nginx         # 1.5 CPU cores max
docker run --cpus="0.5" nginx         # Half CPU core max

# CPU set (specific cores)
docker run --cpuset-cpus="0,1" nginx  # Use only cores 0 and 1
docker run --cpuset-cpus="0-3" nginx  # Use cores 0 through 3

# Update CPU resources
docker update --cpus 2 container_name
docker update --cpu-shares 1024 container_name
```

### Other Resource Limits

```bash
# Disk I/O
docker run --device-read-bps /dev/sda:1mb nginx
docker run --device-write-bps /dev/sda:1mb nginx

# Process limits
docker run --pids-limit=100 nginx

# File descriptor limits
docker run --ulimit nofile=1024:1024 nginx

# Restart policies
docker run --restart=always nginx
docker run --restart=unless-stopped nginx
docker run --restart=on-failure:3 nginx
```

### System Cleanup

```bash
# Comprehensive cleanup
docker system prune                   # Remove stopped containers, unused networks, dangling images

# More aggressive cleanup
docker system prune -a               # Also remove unused images
docker system prune --volumes        # Also remove unused volumes

# Specific cleanup
docker container prune                # Only stopped containers
docker image prune                    # Only dangling images
docker image prune -a                # All unused images
docker network prune                 # Unused networks
docker volume prune                   # Unused volumes
```

---

## Docker Registries & Distribution {#docker-registries}

### Working with Docker Hub

```bash
# Login to Docker Hub
docker login
docker login -u username -p password

# Tag image for hub
docker tag myapp:latest username/myapp:latest
docker tag myapp:latest username/myapp:v1.0

# Push to Docker Hub
docker push username/myapp:latest
docker push username/myapp:v1.0

# Pull from Docker Hub
docker pull username/myapp:latest

# Logout
docker logout
```

### Private Registries

```bash
# Login to private registry
docker login myregistry.com
docker login myregistry.com:5000

# Tag for private registry
docker tag myapp:latest myregistry.com/myapp:latest

# Push to private registry
docker push myregistry.com/myapp:latest

# Pull from private registry
docker pull myregistry.com/myapp:latest
```

### Running Local Registry

```bash
# Start local registry
docker run -d -p 5000:5000 --name registry registry:2

# Tag for local registry
docker tag myapp:latest localhost:5000/myapp:latest

# Push to local registry
docker push localhost:5000/myapp:latest

# Pull from local registry
docker pull localhost:5000/myapp:latest
```

### Registry Best Practices

```bash
# Multi-platform images
docker buildx build --platform linux/amd64,linux/arm64 -t myapp:latest --push .

# Image signing (Docker Content Trust)
export DOCKER_CONTENT_TRUST=1
docker push myapp:latest

# Vulnerability scanning
docker scan myapp:latest

# Registry cleanup
docker exec registry bin/registry garbage-collect /etc/docker/registry/config.yml
```

---

## Multi-Container Apps with Docker Compose {#docker-compose}

Docker Compose simplifies multi-container application management through YAML configuration.

### Basic Compose File

```yaml
version: '3.8'

services:
  web:
    build: .
    ports:
      - "8000:8000"
    volumes:
      - .:/app
    environment:
      - DEBUG=1
    depends_on:
      - db
      - redis

  db:
    image: postgres:13
    environment:
      POSTGRES_DB: myapp
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
    volumes:
      - postgres_data:/var/lib/postgresql/data

  redis:
    image: redis:alpine
    command: redis-server --appendonly yes
    volumes:
      - redis_data:/data

volumes:
  postgres_data:
  redis_data:

networks:
  default:
    driver: bridge
```

### Compose Commands

```bash
# Start all services
docker-compose up
docker-compose up -d                  # Detached mode
docker-compose up web db              # Specific services

# Build and start
docker-compose up --build

# Stop services
docker-compose down
docker-compose down -v               # Also remove volumes
docker-compose down --rmi all        # Also remove images

# View logs
docker-compose logs
docker-compose logs -f web           # Follow logs for web service

# Scale services
docker-compose up --scale web=3

# Execute commands
docker-compose exec web bash
docker-compose run web python manage.py migrate
```

### Advanced Compose Features

```yaml
version: '3.8'

services:
  web:
    build:
      context: .
      dockerfile: Dockerfile.prod
      args:
        - BUILD_ARG=value
    ports:
      - "80:8000"
    volumes:
      - static_volume:/app/static
    environment:
      - DATABASE_URL=postgresql://user:pass@db:5432/dbname
    depends_on:
      db:
        condition: service_healthy
    restart: unless-stopped
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8000/health"]
      interval: 30s
      timeout: 10s
      retries: 3

  db:
    image: postgres:13
    environment:
      POSTGRES_DB: myapp
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
    volumes:
      - postgres_data:/var/lib/postgresql/data
      - ./init.sql:/docker-entrypoint-initdb.d/init.sql
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U user -d myapp"]
      interval: 10s
      timeout: 5s
      retries: 5

  nginx:
    image: nginx:alpine
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - ./nginx.conf:/etc/nginx/nginx.conf
      - static_volume:/app/static
      - ./ssl:/etc/nginx/ssl
    depends_on:
      - web

volumes:
  postgres_data:
    driver: local
  static_volume:

networks:
  frontend:
    driver: bridge
  backend:
    driver: bridge
    internal: true
```

---

## Advanced Dockerfile Techniques {#advanced-dockerfile}

### Dockerfile Best Practices

```dockerfile
# Use specific, minimal base images
FROM node:18-alpine AS base

# Set working directory early
WORKDIR /app

# Copy package files first (for layer caching)
COPY package.json package-lock.json ./

# Install dependencies
RUN npm ci --only=production && npm cache clean --force

# Copy application code
COPY . .

# Create non-root user
RUN addgroup -g 1001 -S nodejs && \
    adduser -S nextjs -u 1001

# Change ownership of app directory
RUN chown -R nextjs:nodejs /app
USER nextjs

# Expose port (documentation)
EXPOSE 3000

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD curl -f http://localhost:3000/health || exit 1

# Default command
CMD ["npm", "start"]
```

### Multi-Stage Builds

```dockerfile
# Build stage
FROM node:18-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
RUN npm run build

# Production stage
FROM node:18-alpine AS runner
WORKDIR /app

# Create user
RUN addgroup --system --gid 1001 nodejs
RUN adduser --system --uid 1001 nextjs

# Copy built application
COPY --from=builder /app/public ./public
COPY --from=builder --chown=nextjs:nodejs /app/.next/standalone ./
COPY --from=builder --chown=nextjs:nodejs /app/.next/static ./.next/static

USER nextjs
EXPOSE 3000
ENV PORT 3000

CMD ["node", "server.js"]
```

### Dockerfile Optimization

```dockerfile
# Use .dockerignore to exclude unnecessary files
# .dockerignore content:
# node_modules
# npm-debug.log
# .git
# README.md
# .env

FROM python:3.9-slim

# Combine RUN instructions to reduce layers
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        gcc \
        libc6-dev && \
    rm -rf /var/lib/apt/lists/*

# Use COPY instead of ADD when possible
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Set environment variables
ENV PYTHONUNBUFFERED=1
ENV PYTHONDONTWRITEBYTECODE=1

# Copy application code last
COPY . .

# Use exec form for CMD/ENTRYPOINT
CMD ["python", "app.py"]
```

### Advanced Dockerfile Features

```dockerfile
# Build arguments
ARG BUILD_VERSION=latest
ARG BUILD_DATE
ARG VCS_REF

# Labels for metadata
LABEL maintainer="your-email@example.com"
LABEL version=$BUILD_VERSION
LABEL build-date=$BUILD_DATE
LABEL vcs-ref=$VCS_REF

# Multi-architecture builds
FROM --platform=$BUILDPLATFORM golang:1.19 AS builder
ARG TARGETPLATFORM
ARG BUILDPLATFORM
RUN echo "Building on $BUILDPLATFORM, targeting $TARGETPLATFORM"

# Conditional instructions
FROM alpine:latest
RUN if [ "$TARGETPLATFORM" = "linux/arm64" ]; then \
      echo "ARM64 specific setup"; \
    else \
      echo "AMD64 specific setup"; \
    fi

# Volume declarations
VOLUME ["/data", "/logs"]

# Stop signal
STOPSIGNAL SIGTERM

# Shell form vs exec form
RUN echo "Shell form - variables expanded: $HOME"
RUN ["echo", "Exec form - variables not expanded: $HOME"]
```

---

## Production Best Practices {#production-best-practices}

### Security Best Practices

```dockerfile
# Use non-root user
FROM node:18-alpine
RUN addgroup -g 1001 -S appgroup && \
    adduser -S appuser -u 1001 -G appgroup
USER appuser

# Use specific versions
FROM nginx:1.21.6-alpine

# Scan for vulnerabilities
docker scan myapp:latest

# Use secrets for sensitive data
docker run -d --name app \
  -e DATABASE_PASSWORD_FILE=/run/secrets/db_password \
  myapp:latest
```

### Health Checks

```dockerfile
# Dockerfile health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD curl -f http://localhost:8080/health || exit 1

# Docker Compose health check
version: '3.8'
services:
  web:
    image: myapp:latest
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8080/health"]
      interval: 30s
      timeout: 10s
      retries: 3
      start_period: 40s
```

### Logging Best Practices

```bash
# Configure logging driver
docker run -d --log-driver=json-file --log-opt max-size=10m --log-opt max-file=3 nginx

# Use centralized logging
docker run -d --log-driver=syslog --log-opt syslog-address=tcp://logserver:514 nginx

# Application logs to stdout/stderr
# Let Docker handle log rotation and collection
```

### Resource Limits in Production

```bash
# Always set resource limits
docker run -d \
  --memory=512m \
  --memory-swap=1g \
  --cpus="1.0" \
  --restart=unless-stopped \
  myapp:latest

# Monitor resource usage
docker stats --no-stream
```

### Secrets Management

```bash
# Docker secrets (Swarm mode)
echo "mysecretpassword" | docker secret create db_password -
docker service create --secret db_password myapp:latest

# External secrets management
docker run -d \
  -e VAULT_ADDR=https://vault.example.com \
  -e VAULT_TOKEN_FILE=/run/secrets/vault_token \
  myapp:latest
```

---

## Troubleshooting & Debugging {#troubleshooting}

### Common Issues and Solutions

**Container Won't Start**
```bash
# Check container logs
docker logs container_name
docker logs -f container_name

# Check container configuration
docker inspect container_name

# Check system events
docker system events

# Run in interactive mode for debugging
docker run -it --entrypoint bash image_name
```

**Network Issues**
```bash
# Check container networking
docker network inspect bridge
docker network inspect container_network

# Test connectivity
docker exec container ping google.com
docker exec container nslookup google.com

# Check port bindings
docker port container_name
netstat -tulpn | grep :8080
```

**Storage Issues**
```bash
# Check disk usage
docker system df
df -h

# Check volume mounts
docker inspect container_name | grep -A 20 Mounts

# Test volume access
docker exec container ls -la /data
docker exec container touch /data/test.txt
```

**Performance Issues**
```bash
# Monitor resource usage
docker stats
top -p $(docker inspect -f '{{.State.Pid}}' container_name)

# Check container processes
docker top container_name

# Analyze image layers
docker history image_name
```

### Debugging Commands

```bash
# Enter running container
docker exec -it container_name bash
docker exec -it container_name sh

# Debug network connectivity
docker run --rm -it nicolaka/netshoot
docker run --rm --net container:target_container nicolaka/netshoot

# Check container filesystem changes
docker diff container_name

# Copy files for debugging
docker cp container_name:/app/logs ./debug-logs
docker cp debug-file.txt container_name:/tmp/

# Run temporary debug container
docker run --rm -it --pid container:target_container alpine ps aux
docker run --rm -it --network container:target_container alpine netstat -tulpn

# Check container resource constraints
docker inspect container_name | grep -A 10 "Memory\|Cpu"

# Debug Docker daemon
sudo journalctl -u docker.service
sudo dockerd --debug
```

### Advanced Debugging

```bash
# Attach to running container (see stdout/stderr)
docker attach container_name

# Stream events from Docker daemon
docker events --filter container=container_name

# Debug image build
docker build --no-cache --progress=plain -t myapp .

# Check layer sizes
docker images --format "table {{.Repository}}\t{{.Tag}}\t{{.Size}}"

# Analyze image efficiency
docker run --rm -it wagoodman/dive myapp:latest

# Debug networking with tcpdump
docker run --rm --net container:target_container nicolaka/netshoot tcpdump -i eth0

# Check DNS resolution
docker exec container_name nslookup service_name
docker exec container_name cat /etc/resolv.conf
```

### Performance Optimization

```bash
# Use multi-stage builds to reduce image size
# Use .dockerignore to exclude unnecessary files
# Leverage build cache effectively
# Use Alpine-based images when possible

# Monitor and optimize resource usage
docker stats --format "table {{.Container}}\t{{.CPUPerc}}\t{{.MemUsage}}\t{{.NetIO}}\t{{.BlockIO}}"

# Use healthchecks for better container management
# Implement proper logging strategies
# Use secrets management for sensitive data
```

---

## Complete Command Reference {#command-reference}

### Image Commands
| Command | Description | Example |
|---------|-------------|---------|
| `docker pull <image>` | Download image from registry | `docker pull nginx:alpine` |
| `docker build -t <name> .` | Build image from Dockerfile | `docker build -t myapp:1.0 .` |
| `docker images` | List all local images | `docker images` |
| `docker rmi <image>` | Remove image | `docker rmi nginx:alpine` |
| `docker tag <src> <dest>` | Tag an image | `docker tag myapp:1.0 myapp:latest` |
| `docker push <image>` | Push image to registry | `docker push myapp:1.0` |
| `docker history <image>` | Show image layer history | `docker history nginx` |
| `docker inspect <image>` | Detailed image information | `docker inspect nginx` |
| `docker save <image>` | Save image to tar file | `docker save nginx > nginx.tar` |
| `docker load` | Load image from tar file | `docker load < nginx.tar` |
| `docker image prune` | Remove dangling images | `docker image prune -a` |

### Container Commands
| Command | Description | Example |
|---------|-------------|---------|
| `docker run <image>` | Create and start container | `docker run -d nginx` |
| `docker create <image>` | Create container (don't start) | `docker create --name web nginx` |
| `docker start <container>` | Start stopped container | `docker start web` |
| `docker stop <container>` | Stop running container | `docker stop web` |
| `docker restart <container>` | Restart container | `docker restart web` |
| `docker pause <container>` | Pause container processes | `docker pause web` |
| `docker unpause <container>` | Unpause container | `docker unpause web` |
| `docker kill <container>` | Kill container (SIGKILL) | `docker kill web` |
| `docker rm <container>` | Remove container | `docker rm web` |
| `docker ps` | List running containers | `docker ps -a` |
| `docker logs <container>` | View container logs | `docker logs -f web` |
| `docker exec <container> <cmd>` | Execute command in container | `docker exec -it web bash` |
| `docker attach <container>` | Attach to running container | `docker attach web` |
| `docker cp <src> <dest>` | Copy files to/from container | `docker cp file.txt web:/app/` |
| `docker commit <container>` | Create image from container | `docker commit web myapp:1.0` |
| `docker diff <container>` | Show filesystem changes | `docker diff web` |
| `docker top <container>` | Show running processes | `docker top web` |
| `docker stats` | Show resource usage | `docker stats --no-stream` |
| `docker update <container>` | Update container resources | `docker update -m 512m web` |
| `docker container prune` | Remove stopped containers | `docker container prune` |

### Network Commands
| Command | Description | Example |
|---------|-------------|---------|
| `docker network ls` | List networks | `docker network ls` |
| `docker network create <name>` | Create network | `docker network create mynet` |
| `docker network rm <name>` | Remove network | `docker network rm mynet` |
| `docker network inspect <name>` | Inspect network | `docker network inspect bridge` |
| `docker network connect <net> <container>` | Connect container to network | `docker network connect mynet web` |
| `docker network disconnect <net> <container>` | Disconnect from network | `docker network disconnect mynet web` |
| `docker network prune` | Remove unused networks | `docker network prune` |

### Volume Commands
| Command | Description | Example |
|---------|-------------|---------|
| `docker volume ls` | List volumes | `docker volume ls` |
| `docker volume create <name>` | Create volume | `docker volume create mydata` |
| `docker volume rm <name>` | Remove volume | `docker volume rm mydata` |
| `docker volume inspect <name>` | Inspect volume | `docker volume inspect mydata` |
| `docker volume prune` | Remove unused volumes | `docker volume prune` |

### System Commands
| Command | Description | Example |
|---------|-------------|---------|
| `docker system info` | Show system information | `docker system info` |
| `docker system df` | Show disk usage | `docker system df -v` |
| `docker system prune` | Clean up unused resources | `docker system prune -a` |
| `docker system events` | Show real-time events | `docker system events` |
| `docker version` | Show Docker version | `docker version` |

### Docker Compose Commands
| Command | Description | Example |
|---------|-------------|---------|
| `docker-compose up` | Start services | `docker-compose up -d` |
| `docker-compose down` | Stop and remove services | `docker-compose down -v` |
| `docker-compose build` | Build services | `docker-compose build web` |
| `docker-compose logs` | View service logs | `docker-compose logs -f web` |
| `docker-compose exec <service> <cmd>` | Execute command in service | `docker-compose exec web bash` |
| `docker-compose ps` | List services | `docker-compose ps` |
| `docker-compose pull` | Pull service images | `docker-compose pull` |
| `docker-compose restart` | Restart services | `docker-compose restart web` |
| `docker-compose scale <service>=<n>` | Scale service | `docker-compose scale web=3` |

### Registry Commands
| Command | Description | Example |
|---------|-------------|---------|
| `docker login` | Login to registry | `docker login myregistry.com` |
| `docker logout` | Logout from registry | `docker logout` |
| `docker search <term>` | Search Docker Hub | `docker search nginx` |

### Common Command Patterns

**Running Containers**
```bash
# Run with port mapping
docker run -p 8080:80 -d nginx

# Run with volume mount
docker run -v mydata:/data -d nginx

# Run with environment variables
docker run -e DATABASE_URL=postgres://user:pass@db:5432/mydb -d myapp

# Run with resource limits
docker run -m 512m --cpus="1.0" -d nginx

# Run with custom network
docker run --network mynet -d nginx

# Run with restart policy
docker run --restart unless-stopped -d nginx

# Run with health check
docker run --health-cmd="curl -f http://localhost:8080/health" -d myapp
```

**Development Workflows**
```bash
# Build and run
docker build -t myapp . && docker run -d myapp

# Development with bind mount
docker run -v $(pwd):/app -p 3000:3000 -d node:alpine

# Quick debugging
docker run --rm -it alpine sh
docker run --rm -it --entrypoint bash myapp

# One-time command execution
docker run --rm -v $(pwd):/app -w /app node:alpine npm install
```

**Production Patterns**
```bash
# Production deployment
docker run -d \
  --name myapp \
  --restart unless-stopped \
  -p 8080:8080 \
  -v myapp-data:/data \
  -e DATABASE_URL=$DATABASE_URL \
  --memory 512m \
  --cpus 1.0 \
  myapp:v1.0

# Health monitoring
docker run -d \
  --name web \
  --health-cmd="curl -f http://localhost:8080/health || exit 1" \
  --health-interval=30s \
  --health-timeout=10s \
  --health-retries=3 \
  myapp:latest
```

---

## Key Takeaways & Next Steps 🎯

Congratulations! You've now covered the complete Docker ecosystem. Here are the key concepts to remember:

### Essential Concepts
- **Containers vs VMs**: Containers share the host kernel, making them lightweight and fast
- **Image Layers**: Docker uses layered filesystem for efficiency and caching
- **Immutable Infrastructure**: Containers should be stateless and replaceable
- **12-Factor App Principles**: Design applications for containerization

### Best Practices Summary
1. **Security**: Use non-root users, scan for vulnerabilities, manage secrets properly
2. **Efficiency**: Use multi-stage builds, .dockerignore, and minimal base images
3. **Reliability**: Implement health checks, set resource limits, use proper restart policies
4. **Monitoring**: Use structured logging, metrics collection, and observability tools

### Next Steps
- **Container Orchestration**: Learn Kubernetes or Docker Swarm for production deployments
- **CI/CD Integration**: Integrate Docker into your development pipeline
- **Advanced Networking**: Explore service mesh and advanced networking patterns
- **Security**: Deep dive into container security best practices
- **Monitoring**: Set up comprehensive monitoring and logging solutions

### Useful Resources
- [Docker Official Documentation](https://docs.docker.com/)
- [Docker Hub](https://hub.docker.com/) - Public registry
- [Docker Best Practices](https://docs.docker.com/develop/dev-best-practices/)
- [Dockerfile Best Practices](https://docs.docker.com/develop/dockerfile_best_practices/)

---

## Final Thoughts

Docker has revolutionized how we build, ship, and run applications. With the knowledge from this comprehensive guide, you're now equipped to:

- Containerize any application effectively
- Design efficient and secure Docker images
- Manage complex multi-container applications
- Troubleshoot and optimize Docker deployments
- Follow production-ready best practices

Remember: Docker is not just a tool—it's a mindset shift toward immutable infrastructure, microservices, and DevOps practices. Keep practicing, stay updated with the latest features, and most importantly, focus on solving real-world problems with containers.

Happy containerizing! 🐳

---

*This guide covers Docker fundamentals through advanced concepts. For the latest updates and features, always refer to the official Docker documentation.*