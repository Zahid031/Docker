# Docker Mastery Guide - From Scratch to Pro Level

## Table of Contents
1. [Why Docker? Why Container-Based Software Delivery?](#why-docker)
2. [Understanding the Container Landscape](#container-landscape)
3. [Docker Architecture & Components](#docker-architecture)
4. [Docker Under the Hood](#docker-under-the-hood)
5. [Getting Started with Docker](#getting-started)
6. [Working with Images](#working-with-images)
7. [Container Lifecycle Management](#container-lifecycle)
8. [Port Mapping & Networking](#port-mapping-networking)
9. [Docker Networking Deep Dive](#docker-networking)
10. [Data Persistence with Volumes](#docker-volumes)
11. [Resource Management & Monitoring](#resource-management)
12. [Docker Compose](#docker-compose)
13. [Multi-Stage Builds](#multi-stage-builds)
14. [Docker Security Best Practices](#security-best-practices)
15. [Production Deployment Strategies](#production-deployment)
16. [Troubleshooting & Debugging](#troubleshooting)
17. [Advanced Topics](#advanced-topics)

---

## Why Docker? Why Container-Based Software Delivery? {#why-docker}

### Problems Docker Solves
- **"It works on my machine"** - Ensures consistent environments across development, testing, and production
- **Dependency Hell** - Packages applications with all dependencies
- **Resource Efficiency** - Lightweight compared to virtual machines
- **Scalability** - Easy horizontal scaling and orchestration
- **DevOps Integration** - Seamless CI/CD pipeline integration
- **Microservices Architecture** - Perfect for breaking monoliths into services

### Benefits of Container-Based Delivery
- **Portability**: Run anywhere - dev laptop, test server, cloud
- **Consistency**: Same environment across all stages
- **Speed**: Fast startup times compared to VMs
- **Resource Utilization**: Share OS kernel, use fewer resources
- **Isolation**: Applications run in isolated environments
- **Version Control**: Infrastructure as code

---

## Understanding the Container Landscape {#container-landscape}

### Bare Metal vs VMs vs Containers

| Aspect | Bare Metal | Virtual Machines | Containers |
|--------|------------|------------------|------------|
| **OS Overhead** | None | Full OS per VM | Shared OS kernel |
| **Resource Usage** | Maximum | High | Minimal |
| **Startup Time** | Boot time | Minutes | Seconds |
| **Isolation** | Hardware | Complete | Process-level |
| **Portability** | Limited | Good | Excellent |
| **Density** | 1 app | Few VMs | Many containers |

**Key Difference**: VMs have their own kernel, containers share the host OS kernel.

---

## Docker Architecture & Components {#docker-architecture}

### Client-Server Architecture
```
Docker Client  <--REST API-->  Docker Daemon (dockerd)
     |                              |
     |                              |
  Commands                    Container Runtime
(docker run,                      |
docker build,                containerd
docker push)                      |
                                 runc
```

### Core Components
- **Docker Client**: Command-line interface (CLI)
- **Docker Daemon (dockerd)**: Server that manages images, containers, networks, volumes
- **containerd**: High-level container runtime
- **runc**: Low-level container runtime (OCI compliant)
- **Docker Registry**: Storage for Docker images (Docker Hub, private registries)

### System Information Commands
```bash
# Get system information
docker system info

# Monitor system events (server-side logs)
docker system events

# Check resource usage
docker system df

# Clean up system resources
docker system prune
```

---

## Docker Under the Hood {#docker-under-the-hood}

### Linux Namespaces
Docker uses Linux namespaces for isolation:
- **PID**: Process isolation
- **NET**: Network isolation
- **MNT**: Filesystem mount points
- **UTS**: Hostname and domain name
- **IPC**: Inter-process communication
- **USER**: User and group IDs

### Control Groups (cgroups)
Resource control and monitoring:
- **CPU**: Limit CPU usage
- **Memory**: Control memory allocation
- **Disk I/O**: Manage disk access
- **Network**: Control network bandwidth

### Union & Overlay Filesystem
- **Layered Architecture**: Images are built in layers
- **Copy-on-Write (COW)**: Containers share image layers, only changes are written
- **Overlay2**: Default storage driver for efficient space usage
- **Image Reuse**: When pulling images, only new/changed layers are downloaded

---

## Getting Started with Docker {#getting-started}

### Essential Commands
```bash
# Pull an image from registry
docker pull nginx:latest

# Create a container (without starting)
docker create --name my-nginx nginx

# Run a container
docker run nginx

# Run in interactive mode
docker run -it ubuntu bash

# Run in detached mode
docker run -d nginx

# Run with auto-removal after exit
docker run -d --rm nginx

# Stop a container
docker stop container_name

# Start a stopped container
docker start container_name

# Create image from running container
docker commit container_id new_image:tag
```

### Container Inspection & Logs
```bash
# List running containers
docker ps

# List all containers (including stopped)
docker ps -a

# Show last created container
docker ps -l

# Show last n containers
docker ps -n 2

# View container logs
docker logs container_name

# Follow logs in real-time
docker logs -f container_name

# Inspect container details
docker inspect container_name

# Execute command in running container
docker exec -it container_name bash

# See file changes in container
docker diff container_name
```

---

## Working with Images {#working-with-images}

### Image Management
```bash
# List images
docker images

# View image history/layers
docker history image_name

# Remove image
docker rmi image_name

# Remove dangling images
docker image prune

# Build image from Dockerfile
docker build -t image_name:tag .

# Tag an image
docker tag source_image:tag target_image:tag

# Push image to registry
docker push image_name:tag
```

### Dockerfile Best Practices
```dockerfile
# Use specific base image versions
FROM node:18-alpine

# Set working directory
WORKDIR /app

# Copy package files first (for better caching)
COPY package*.json ./

# Install dependencies
RUN npm ci --only=production

# Copy application code
COPY . .

# Expose port (documentation purposes)
EXPOSE 3000

# Use non-root user
USER node

# Set entrypoint
CMD ["npm", "start"]
```

---

## Container Lifecycle Management {#container-lifecycle}

### Container States
1. **Created**: Container exists but not started
2. **Running**: Container is executing
3. **Paused**: Container processes are paused
4. **Stopped**: Container has exited
5. **Deleted**: Container is removed

### Lifecycle Commands
```bash
# Create container
docker create --name web nginx

# Start container
docker start web

# Stop container (SIGTERM -> 10s -> SIGKILL)
docker stop web

# Kill container immediately (SIGKILL)
docker kill web

# Pause/Unpause container
docker pause web
docker unpause web

# Remove container
docker rm web

# Force remove running container
docker rm -f web

# Remove multiple containers
docker rm container1 container2

# Remove all stopped containers
docker container prune
```

---

## Port Mapping & Networking {#port-mapping-networking}

### Port Mapping Options
```bash
# Map host port 8000 to container port 80
docker run -p 8000:80 nginx

# Map random host port to container port 80
docker run -P nginx

# Expose container port 80 to random host port
docker run -p 80 nginx

# Map to specific interface
docker run -p 127.0.0.1:8000:80 nginx

# Map multiple ports
docker run -p 8000:80 -p 8443:443 nginx
```

### Dockerfile Port Exposure
```dockerfile
# Expose ports (documentation - doesn't publish)
EXPOSE 80 443

# When you expose ports in Dockerfile, Docker chooses random host ports
# Use -P flag to publish all exposed ports to random host ports
```

---

## Docker Networking Deep Dive {#docker-networking}

### Network Types
1. **bridge**: Default network for containers
2. **host**: Container uses host's network stack
3. **none**: No networking
4. **overlay**: Multi-host networking (Docker Swarm)
5. **macvlan**: Assign MAC address to container

### Network Management
```bash
# List networks
docker network ls

# Create custom network
docker network create -d bridge my-network

# Create overlay network
docker network create -d overlay my-overlay

# Inspect network
docker network inspect my-network

# Connect container to network
docker network connect my-network container_name

# Disconnect container from network
docker network disconnect my-network container_name

# Remove network
docker network rm my-network

# Remove unused networks
docker network prune
```

### Bridge vs User-Defined Networks
```bash
# Default bridge network
docker run -d --name web1 nginx
docker run -d --name web2 nginx
# Communication: Only via IP addresses

# User-defined bridge network
docker network create my-net
docker run -d --name web1 --network my-net nginx
docker run -d --name web2 --network my-net nginx
# Communication: Via container names (DNS resolution)
```

---

## Data Persistence with Volumes {#docker-volumes}

### Volume Types
1. **Named Volumes**: Managed by Docker
2. **Bind Mounts**: Host directory mounted into container
3. **tmpfs Mounts**: Temporary filesystem in memory

### Volume Management
```bash
# Create named volume
docker volume create my-volume

# List volumes
docker volume ls

# Inspect volume
docker volume inspect my-volume

# Remove volume
docker volume rm my-volume

# Remove unused volumes
docker volume prune

# Volume location on host
# /var/lib/docker/volumes/
```

### Using Volumes
```bash
# Named volume
docker run -v my-volume:/data nginx

# Bind mount (absolute path required)
docker run -v /host/path:/container/path nginx
docker run -v $(pwd):/app nginx

# Read-only bind mount
docker run -v /host/path:/container/path:ro nginx

# Using --mount (more explicit)
docker run --mount source=my-volume,destination=/data nginx
docker run --mount type=bind,source=/host/path,target=/container/path nginx
```

### Volume vs Mount Comparison
| Feature | -v/--volume | --mount |
|---------|-------------|---------|
| Syntax | Compact | Verbose, explicit |
| Bind mounts | Limited options | Full options |
| Named volumes | Simple | Full control |
| tmpfs | Not supported | Supported |

---

## Resource Management & Monitoring {#resource-management}

### Container Stats
```bash
# Real-time stats for all containers
docker stats

# Stats for specific containers
docker stats container1 container2

# One-time stats (no streaming)
docker stats --no-stream

# Detailed container stats
docker container stats
```

### Resource Limits
```bash
# Memory limit
docker run -m 512M nginx
docker run --memory=1G nginx

# Update memory limit for running container
docker update -m 256M container_name

# CPU limits
docker run --cpus="1.5" nginx          # 1.5 CPU cores
docker run --cpu-shares=512 nginx      # Relative CPU weight (default 1024)

# Multiple resource limits
docker run -m 512M --cpus="1" nginx
```

### File Operations
```bash
# Copy files to/from container
docker cp file.txt container_id:/app/
docker cp container_id:/app/file.txt ./

# Copy entire directory
docker cp ./src container_id:/app/
```

---

## Docker Compose {#docker-compose}

### What is Docker Compose?
Tool for defining and running multi-container Docker applications using YAML files.

### Basic docker-compose.yml
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
    networks:
      - app-network

  db:
    image: postgres:13
    environment:
      - POSTGRES_DB=myapp
      - POSTGRES_USER=user
      - POSTGRES_PASSWORD=password
    volumes:
      - postgres_data:/var/lib/postgresql/data
    networks:
      - app-network

volumes:
  postgres_data:

networks:
  app-network:
    driver: bridge
```

### Compose Commands
```bash
# Start services
docker-compose up

# Start in detached mode
docker-compose up -d

# Build and start
docker-compose up --build

# Stop services
docker-compose down

# Stop and remove volumes
docker-compose down -v

# View logs
docker-compose logs

# Scale services
docker-compose up --scale web=3

# Execute command in service
docker-compose exec web bash

# View running services
docker-compose ps
```

---

## Multi-Stage Builds {#multi-stage-builds}

### Why Multi-Stage Builds?
- Reduce final image size
- Separate build and runtime environments
- Security: Don't include build tools in production

### Example Multi-Stage Dockerfile
```dockerfile
# Build stage
FROM node:18-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
RUN npm run build

# Production stage
FROM nginx:alpine AS production
COPY --from=builder /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]

# Development stage
FROM node:18-alpine AS development
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
EXPOSE 3000
CMD ["npm", "run", "dev"]
```

### Building Specific Stages
```bash
# Build production stage
docker build --target production -t myapp:prod .

# Build development stage
docker build --target development -t myapp:dev .
```

---

## Docker Security Best Practices {#security-best-practices}

### Image Security
```dockerfile
# Use official, minimal base images
FROM node:18-alpine

# Don't run as root
RUN addgroup -g 1001 -S nodejs
RUN adduser -S nextjs -u 1001
USER nextjs

# Use specific versions
FROM node:18.17.0-alpine

# Scan for vulnerabilities
# docker scan image_name
```

### Runtime Security
```bash
# Run with read-only filesystem
docker run --read-only nginx

# Drop capabilities
docker run --cap-drop=ALL --cap-add=NET_BIND_SERVICE nginx

# Use security profiles
docker run --security-opt seccomp=default.json nginx

# Limit resources
docker run -m 512M --cpus="0.5" nginx
```

### Secrets Management
```bash
# Use Docker secrets (Swarm mode)
echo "my-secret" | docker secret create my-secret -

# Use environment files
docker run --env-file .env nginx

# Avoid secrets in Dockerfile
# DON'T: ENV PASSWORD=secret
# DO: Mount secrets at runtime
```

---

## Production Deployment Strategies {#production-deployment}

### Health Checks
```dockerfile
# Dockerfile health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD curl -f http://localhost/ || exit 1
```

```bash
# Runtime health check
docker run -d --health-cmd="curl -f http://localhost/" \
  --health-interval=30s --health-timeout=3s --health-retries=3 nginx
```

### Logging
```bash
# Configure log driver
docker run --log-driver=json-file --log-opt max-size=10m --log-opt max-file=3 nginx

# Use centralized logging
docker run --log-driver=syslog --log-opt syslog-address=udp://logserver:514 nginx
```

### Restart Policies
```bash
# Always restart
docker run --restart=always nginx

# Restart on failure
docker run --restart=on-failure:3 nginx

# Restart unless stopped
docker run --restart=unless-stopped nginx
```

---

## Troubleshooting & Debugging {#troubleshooting}

### Common Issues & Solutions
```bash
# Container exits immediately
docker logs container_name
docker run -it image_name bash  # Debug interactively

# Permission issues
docker exec -it container_name ls -la /path
# Fix: Adjust file permissions or user context

# Network connectivity issues
docker network ls
docker network inspect bridge
docker exec -it container_name ping other_container

# Resource issues
docker stats
docker system df
docker system events

# Port conflicts
netstat -tulpn | grep :port
docker port container_name
```

### Debugging Commands
```bash
# Enter container for debugging
docker exec -it container_name bash

# Check container processes
docker top container_name

# Monitor resource usage
docker stats container_name

# Inspect container configuration
docker inspect container_name

# View container filesystem changes
docker diff container_name

# Export container filesystem
docker export container_name > container.tar
```

---

## Advanced Topics {#advanced-topics}

### Docker Registry
```bash
# Login to registry
docker login registry.example.com

# Push to private registry
docker tag myapp:latest registry.example.com/myapp:latest
docker push registry.example.com/myapp:latest

# Run local registry
docker run -d -p 5000:5000 --name registry registry:2
```

### Docker Swarm (Orchestration)
```bash
# Initialize swarm
docker swarm init

# Join swarm as worker
docker swarm join --token TOKEN manager-ip:2377

# Deploy stack
docker stack deploy -c docker-compose.yml mystack

# Scale service
docker service scale mystack_web=3
```

### Performance Optimization
```bash
# Use BuildKit for faster builds
export DOCKER_BUILDKIT=1
docker build .

# Multi-platform builds
docker buildx build --platform linux/amd64,linux/arm64 -t myapp .

# Use cache mounts
RUN --mount=type=cache,target=/var/cache/apt apt-get update
```

### Container Orchestration Alternatives
- **Kubernetes**: Production-grade orchestration
- **Docker Swarm**: Docker's native orchestration
- **Amazon ECS**: AWS container service
- **Google Cloud Run**: Serverless containers

---

## Quick Reference Commands

### Most Used Commands
```bash
# Images
docker pull image
docker build -t name .
docker images
docker rmi image

# Containers
docker run -d image
docker ps
docker stop container
docker rm container

# Logs & Debug
docker logs container
docker exec -it container bash
docker inspect container

# Cleanup
docker system prune
docker container prune
docker image prune
docker volume prune
```

### Dockerfile Instructions
```dockerfile
FROM          # Base image
WORKDIR       # Set working directory
COPY/ADD      # Copy files
RUN           # Execute commands
EXPOSE        # Document ports
ENV           # Environment variables
USER          # Set user context
CMD           # Default command
ENTRYPOINT    # Entry point
VOLUME        # Mount point
HEALTHCHECK   # Health check
```

---

## Learning Path Recommendations

### Beginner (Weeks 1-2)
- Understand containers vs VMs
- Learn basic Docker commands
- Create simple Dockerfiles
- Work with volumes and networking

### Intermediate (Weeks 3-4)
- Master Docker Compose
- Implement multi-stage builds
- Learn security best practices
- Practice troubleshooting

### Advanced (Weeks 5-8)
- Container orchestration (Kubernetes)
- CI/CD integration
- Performance optimization
- Production deployment strategies

### Expert Level
- Custom network drivers
- Storage drivers
- Docker plugin development
- Contributing to Docker ecosystem

---

**Remember**: Practice is key! Set up a local development environment and experiment with these concepts. Start with simple applications and gradually work your way up to complex multi-service architectures.

**Next Steps**: 
1. Set up a multi-tier application with Docker Compose
2. Implement CI/CD pipeline with Docker
3. Learn Kubernetes for production orchestration
4. Explore cloud-native technologies (service mesh, observability)

Happy Dockerizing! 🐳