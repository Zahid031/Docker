# Docker Mastery Guide - From Scratch to Pro Level 🐳

Welcome to the ultimate Docker tutorial! This guide will take you from a beginner to a pro, covering everything you need to know to master Docker. Let's dive in! 🚀

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
18. [Docker Commands Cheat Sheet](#docker-commands-cheat-sheet)

---

## Why Docker? Why Container-Based Software Delivery? {#why-docker}

### Problems Docker Solves
- **"It works on my machine"** - Ensures consistent environments across development, testing, and production.
- **Dependency Hell** - Packages applications with all their dependencies.
- **Resource Efficiency** - Lightweight compared to virtual machines.
- **Scalability** - Easy horizontal scaling and orchestration.
- **DevOps Integration** - Seamless CI/CD pipeline integration.
- **Microservices Architecture** - Perfect for breaking monoliths into smaller services.

### Benefits of Container-Based Delivery
- **Portability**: Run anywhere - your laptop, a test server, or in the cloud.
- **Consistency**: The same environment across all stages of the development lifecycle.
- **Speed**: Fast startup times compared to VMs.
- **Resource Utilization**: Containers share the host OS kernel, using fewer resources.
- **Isolation**: Applications run in isolated environments, preventing conflicts.
- **Version Control**: Treat your infrastructure as code.

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

**Key Difference**: VMs have their own kernel, while containers share the host OS kernel. This makes containers much more lightweight and faster.

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
- **Docker Client**: The command-line interface (CLI) you use to interact with Docker.
- **Docker Daemon (dockerd)**: The server that manages images, containers, networks, and volumes.
- **containerd**: A high-level container runtime.
- **runc**: A low-level container runtime (OCI compliant).
- **Docker Registry**: A storage system for Docker images (like Docker Hub or a private registry).

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
Docker uses Linux namespaces to provide isolation for containers:
- **PID**: Process isolation.
- **NET**: Network isolation.
- **MNT**: Filesystem mount points.
- **UTS**: Hostname and domain name.
- **IPC**: Inter-process communication.
- **USER**: User and group IDs.

### Control Groups (cgroups)
cgroups are used for resource control and monitoring:
- **CPU**: Limit CPU usage.
- **Memory**: Control memory allocation.
- **Disk I/O**: Manage disk access.
- **Network**: Control network bandwidth.

### Union & Overlay Filesystem
- **Layered Architecture**: Images are built in layers.
- **Copy-on-Write (COW)**: Containers share image layers, and only the changes are written to a new layer.
- **Overlay2**: The default storage driver for efficient space usage.
- **Image Reuse**: When you pull an image, only the new or changed layers are downloaded.

---

## Getting Started with Docker {#getting-started}

### Essential Commands
```bash
# Pull an image from a registry (like Docker Hub)
docker pull nginx:latest

# Create a container (without starting it)
docker create --name my-nginx nginx

# Run a container
docker run nginx

# Run a container in interactive mode
docker run -it ubuntu bash

# Run a container in detached mode (in the background)
docker run -d nginx

# Run a container and automatically remove it when it exits
docker run -d --rm nginx

# Stop a running container
docker stop container_name

# Start a stopped container
docker start container_name

# Create an image from a running container
docker commit container_id new_image:tag
```

### Container Inspection & Logs
```bash
# List running containers
docker ps

# List all containers (including stopped ones)
docker ps -a

# Show the last created container
docker ps -l

# Show the last n created containers
docker ps -n 2

# View the logs of a container
docker logs container_name

# Follow the logs in real-time
docker logs -f container_name

# Inspect the details of a container
docker inspect container_name

# Execute a command in a running container
docker exec -it container_name bash

# See the file changes in a container
docker diff container_name
```

---

## Working with Images {#working-with-images}

### Image Management
```bash
# List all images on your system
docker images

# View the history and layers of an image
docker history image_name

# Remove an image
docker rmi image_name

# Remove all dangling images (images without a tag)
docker image prune

# Build an image from a Dockerfile
docker build -t image_name:tag .

# Tag an image
docker tag source_image:tag target_image:tag

# Push an image to a registry
docker push image_name:tag
```

### Dockerfile Best Practices
```dockerfile
# Use a specific base image version
FROM node:18-alpine

# Set the working directory
WORKDIR /app

# Copy package files first (for better caching)
COPY package*.json ./

# Install dependencies
RUN npm ci --only=production

# Copy the application code
COPY . .

# Expose a port (for documentation purposes)
EXPOSE 3000

# Use a non-root user for security
USER node

# Set the default command to run when the container starts
CMD ["npm", "start"]
```

---

## Container Lifecycle Management {#container-lifecycle}

### Container States
1. **Created**: The container has been created but not started.
2. **Running**: The container is currently executing.
3. **Paused**: The container's processes are paused.
4. **Stopped**: The container has exited.
5. **Deleted**: The container has been removed.

### Lifecycle Commands
```bash
# Create a container
docker create --name web nginx

# Start a container
docker start web

# Stop a container (sends SIGTERM, then SIGKILL after a grace period)
docker stop web

# Kill a container immediately (sends SIGKILL)
docker kill web

# Pause and unpause a container
docker pause web
docker unpause web

# Remove a container
docker rm web

# Force remove a running container
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

# Map a random host port to container port 80
docker run -P nginx

# Expose container port 80 to a random host port
docker run -p 80 nginx

# Map to a specific network interface
docker run -p 127.0.0.1:8000:80 nginx

# Map multiple ports
docker run -p 8000:80 -p 8443:443 nginx
```

### Dockerfile Port Exposure
```dockerfile
# Expose ports (this is for documentation and doesn't publish the ports)
EXPOSE 80 443

# When you expose ports in a Dockerfile, Docker can choose random host ports.
# Use the -P flag to publish all exposed ports to random host ports.
```

---

## Docker Networking Deep Dive {#docker-networking}

### Network Types
1. **bridge**: The default network for containers.
2. **host**: The container uses the host's network stack.
3. **none**: The container has no networking.
4. **overlay**: Used for multi-host networking (in Docker Swarm).
5. **macvlan**: Assigns a MAC address to a container.

### Network Management
```bash
# List all networks
docker network ls

# Create a custom network
docker network create -d bridge my-network

# Create an overlay network
docker network create -d overlay my-overlay

# Inspect a network
docker network inspect my-network

# Connect a container to a network
docker network connect my-network container_name

# Disconnect a container from a network
docker network disconnect my-network container_name

# Remove a network
docker network rm my-network

# Remove all unused networks
docker network prune
```

### Bridge vs User-Defined Networks
```bash
# Default bridge network
docker run -d --name web1 nginx
docker run -d --name web2 nginx
# Communication is only possible via IP addresses.

# User-defined bridge network
docker network create my-net
docker run -d --name web1 --network my-net nginx
docker run -d --name web2 --network my-net nginx
# Communication is possible via container names (thanks to DNS resolution).
```

---

## Data Persistence with Volumes {#docker-volumes}

### Volume Types
1. **Named Volumes**: Managed by Docker.
2. **Bind Mounts**: A host directory is mounted into a container.
3. **tmpfs Mounts**: A temporary filesystem in memory.

### Volume Management
```bash
# Create a named volume
docker volume create my-volume

# List all volumes
docker volume ls

# Inspect a volume
docker volume inspect my-volume

# Remove a volume
docker volume rm my-volume

# Remove all unused volumes
docker volume prune

# The location of volumes on the host
# /var/lib/docker/volumes/
```

### Using Volumes
```bash
# Use a named volume
docker run -v my-volume:/data nginx

# Use a bind mount (requires an absolute path)
docker run -v /host/path:/container/path nginx
docker run -v $(pwd):/app nginx

# Use a read-only bind mount
docker run -v /host/path:/container/path:ro nginx

# Use the --mount flag (more explicit)
docker run --mount source=my-volume,destination=/data nginx
docker run --mount type=bind,source=/host/path,target=/container/path nginx
```

### Volume vs Mount Comparison
| Feature | -v/--volume | --mount |
|---------|-------------|---------|
| Syntax | Compact | Verbose and explicit |
| Bind mounts | Limited options | Full options |
| Named volumes | Simple | Full control |
| tmpfs | Not supported | Supported |

---

## Resource Management & Monitoring {#resource-management}

### Container Stats
```bash
# Get real-time stats for all containers
docker stats

# Get stats for specific containers
docker stats container1 container2

# Get a one-time stats report (no streaming)
docker stats --no-stream

# Get detailed container stats
docker container stats
```

### Resource Limits
```bash
# Set a memory limit
docker run -m 512M nginx
docker run --memory=1G nginx

# Update the memory limit for a running container
docker update -m 256M container_name

# Set CPU limits
docker run --cpus="1.5" nginx          # Use 1.5 CPU cores
docker run --cpu-shares=512 nginx      # Set a relative CPU weight (default is 1024)

# Set multiple resource limits
docker run -m 512M --cpus="1" nginx
```

### File Operations
```bash
# Copy files to and from a container
docker cp file.txt container_id:/app/
docker cp container_id:/app/file.txt ./

# Copy an entire directory
docker cp ./src container_id:/app/
```

---

## Docker Compose {#docker-compose}

### What is Docker Compose?
Docker Compose is a tool for defining and running multi-container Docker applications using YAML files.

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
# Start the services
docker-compose up

# Start the services in detached mode
docker-compose up -d

# Build and start the services
docker-compose up --build

# Stop the services
docker-compose down

# Stop the services and remove volumes
docker-compose down -v

# View the logs
docker-compose logs

# Scale services
docker-compose up --scale web=3

# Execute a command in a service
docker-compose exec web bash

# View the running services
docker-compose ps
```

---

## Multi-Stage Builds {#multi-stage-builds}

### Why Multi-Stage Builds?
- Reduce the final image size.
- Separate the build and runtime environments.
- Improve security by not including build tools in the production image.

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
# Build the production stage
docker build --target production -t myapp:prod .

# Build the development stage
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
# Run with a read-only filesystem
docker run --read-only nginx

# Drop unnecessary capabilities
docker run --cap-drop=ALL --cap-add=NET_BIND_SERVICE nginx

# Use security profiles
docker run --security-opt seccomp=default.json nginx

# Limit resources
docker run -m 512M --cpus="0.5" nginx
```

### Secrets Management
```bash
# Use Docker secrets (in Swarm mode)
echo "my-secret" | docker secret create my-secret -

# Use environment files
docker run --env-file .env nginx

# Avoid secrets in the Dockerfile
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
# Configure the log driver
docker run --log-driver=json-file --log-opt max-size=10m --log-opt max-file=10 nginx

# Use centralized logging
docker run --log-driver=syslog --log-opt syslog-address=udp://logserver:514 nginx
```

### Restart Policies
```bash
# Always restart the container
docker run --restart=always nginx

# Restart on failure (up to 3 times)
docker run --restart=on-failure:3 nginx

# Restart unless the container is stopped
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
# Fix: Adjust file permissions or the user context

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
# Enter a container for debugging
docker exec -it container_name bash

# Check the processes running in a container
docker top container_name

# Monitor the resource usage of a container
docker stats container_name

# Inspect the configuration of a container
docker inspect container_name

# View the filesystem changes in a container
docker diff container_name

# Export the filesystem of a container
docker export container_name > container.tar
```

---

## Advanced Topics {#advanced-topics}

### Docker Registry
```bash
# Login to a registry
docker login registry.example.com

# Push to a private registry
docker tag myapp:latest registry.example.com/myapp:latest
docker push registry.example.com/myapp:latest

# Run a local registry
docker run -d -p 5000:5000 --name registry registry:2
```

### Docker Swarm (Orchestration)
```bash
# Initialize a swarm
docker swarm init

# Join a swarm as a worker
docker swarm join --token TOKEN manager-ip:2377

# Deploy a stack
docker stack deploy -c docker-compose.yml mystack

# Scale a service
docker service scale mystack_web=3
```

### Performance Optimization
```bash
# Use BuildKit for faster builds
export DOCKER_BUILDKIT=1
docker build .

# Create multi-platform builds
docker buildx build --platform linux/amd64,linux/arm64 -t myapp .

# Use cache mounts
RUN --mount=type=cache,target=/var/cache/apt apt-get update
```

### Container Orchestration Alternatives
- **Kubernetes**: The production-grade standard for container orchestration.
- **Docker Swarm**: Docker's native orchestration solution.
- **Amazon ECS**: AWS's container service.
- **Google Cloud Run**: A serverless container platform.

---

## Docker Commands Cheat Sheet {#docker-commands-cheat-sheet}

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
FROM          # Specifies the base image
WORKDIR       # Sets the working directory
COPY/ADD      # Copies files into the image
RUN           # Executes commands
EXPOSE        # Documents the ports
ENV           # Sets environment variables
USER          # Sets the user context
CMD           # Provides the default command
ENTRYPOINT    # Configures the entry point
VOLUME        # Creates a mount point
HEALTHCHECK   # Defines a health check
```

---

## Learning Path Recommendations

### Beginner (Weeks 1-2)
- Understand the difference between containers and VMs.
- Learn the basic Docker commands.
- Create simple Dockerfiles.
- Work with volumes and networking.

### Intermediate (Weeks 3-4)
- Master Docker Compose.
- Implement multi-stage builds.
- Learn security best practices.
- Practice troubleshooting common issues.

### Advanced (Weeks 5-8)
- Explore container orchestration with Kubernetes.
- Integrate Docker into a CI/CD pipeline.
- Learn performance optimization techniques.
- Practice production deployment strategies.

### Expert Level
- Develop custom network and storage drivers.
- Create Docker plugins.
- Contribute to the Docker ecosystem.

---

**Remember**: Practice is key! Set up a local development environment and experiment with these concepts. Start with simple applications and gradually work your way up to complex multi-service architectures.

**Next Steps**: 
1. Set up a multi-tier application with Docker Compose.
2. Implement a CI/CD pipeline with Docker.
3. Learn Kubernetes for production-grade orchestration.
4. Explore cloud-native technologies like service mesh and observability.

Happy Dockerizing! 🐳
