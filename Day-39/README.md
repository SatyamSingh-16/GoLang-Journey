# 🚀 GoLang Journey - Day 55

---

# 🎯 Goal of the Day

Today's goal was to build a strong foundation in Docker and understand how Docker is used to package and run backend applications in consistent environments.

Instead of treating Docker as just a collection of commands, I learned how images, containers, Docker Engine, registries, container lifecycle, writable layers, and persistent storage fit together.

I also created and managed my first real Docker container using Redis, connecting my previous Redis knowledge with Docker and understanding how backend infrastructure can be containerized.

---

# 📚 Topics Covered

## Docker Fundamentals

Learned the fundamental problem Docker solves.

Docker provides a consistent environment for running applications by packaging the application and its required environment.

Covered:

- Docker
- Containerization
- Application Environment
- Docker Engine
- Containers
- Images
- Reproducible Environments

---

## Docker Images

Studied Docker images as reusable application artifacts.

Covered:

- Docker Images
- Image Layers
- Read-Only Layers
- Image Reusability
- Image Tags
- Image Repositories
- Image References

An image acts as the base from which containers are created.

Example:

```bash
redis:7
```

---

## Docker Containers

Studied how containers are created from Docker images.

Covered:

- Containers
- Container Instances
- Writable Container Layer
- Container Runtime
- Container Isolation
- Container Processes

Learned that multiple containers can be created from the same image while maintaining their own runtime state.

---

## Docker Architecture

Studied the major components involved in running containers.

Covered:

- Docker CLI
- Docker Engine
- Docker Daemon
- Docker Images
- Docker Containers
- Container Runtime
- Docker Client-Server Model

Learned how Docker commands communicate with the Docker Engine to create and manage containers.

---

## Docker Hub & Container Registries

Learned how Docker images are stored and distributed.

Covered:

- Docker Hub
- Container Registry
- Image Repository
- Image Tags
- `docker pull`
- `docker push`
- Private Registries
- Cloud Registries

Learned the difference between a container registry and the Docker containers running locally.

---

## Image vs Container

Studied the relationship between images and containers.

Covered:

- Immutable Images
- Writable Container Layer
- Copy-on-Write
- Image Reuse
- Container State
- Ephemeral Containers
- Persistent Data

Learned that a container does not modify the underlying image directly and instead has its own writable layer.

---

## Container Lifecycle

Studied the complete lifecycle of a Docker container.

Covered:

- `docker create`
- `docker start`
- `docker run`
- `docker stop`
- `docker restart`
- `docker rm`
- Created State
- Running State
- Stopped State
- Deleted State
- Container Exit Codes

Learned how a container moves between different states throughout its lifecycle.

```text
IMAGE
  |
  | docker create
  v
CREATED
  |
  | docker start
  v
RUNNING
  |
  | docker stop
  v
STOPPED
  |
  | docker rm
  v
DELETED
```

A stopped container can also be started again:

```text
STOPPED
   |
   | docker start
   v
RUNNING
```

---

## First Real Docker Container

Created and managed my first real Docker container using Redis.

Covered:

- Redis Docker Image
- `docker run`
- Container Naming
- `docker ps`
- `docker ps -a`
- `docker logs`
- `docker exec`
- `docker stop`
- `docker start`
- `docker rm`

Learned how to operate a real container from the terminal.

---

## Docker + Redis

Connected my previous Redis knowledge with Docker.

Created a Redis container using:

```bash
docker run --name my-redis redis
```

Then accessed Redis inside the container using:

```bash
docker exec -it my-redis redis-cli
```

Performed Redis operations:

```redis
SET name Satyam
GET name
```

This demonstrated how infrastructure such as Redis can be run inside containers without requiring Redis to be directly installed on the host system.

---

## Essential Docker Commands

Studied the most important Docker commands used during development and debugging.

### Images

```bash
docker images
docker pull
docker build
docker rmi
docker push
```

### Containers

```bash
docker run
docker ps
docker ps -a
docker start
docker stop
docker restart
docker rm
```

### Debugging

```bash
docker logs
docker exec
docker inspect
docker stats
```

### Cleanup

```bash
docker container prune
docker image prune
docker system prune
```

---

# 💻 Programs Written

- Docker Redis Container
- Redis Container Management
- Container Lifecycle Operations
- Docker Image Operations
- Container Inspection
- Container Logs
- Executing Commands Inside Containers
- Docker Port Mapping Examples

---

# 🧠 Important Concepts Learned

- Docker
- Containerization
- Docker Engine
- Docker CLI
- Docker Images
- Docker Containers
- Image Layers
- Writable Container Layer
- Copy-on-Write
- Container Registry
- Docker Hub
- Image Repository
- Image Tags
- `docker pull`
- `docker push`
- `docker build`
- `docker run`
- `docker create`
- `docker start`
- `docker stop`
- `docker restart`
- `docker rm`
- `docker rmi`
- `docker ps`
- `docker logs`
- `docker exec`
- `docker inspect`
- `docker stats`
- Container Lifecycle
- Container Exit Codes
- Port Mapping
- Container Processes
- Persistent Storage
- Ephemeral Containers

---

# ⚠️ Common Mistakes I Learned

- Confusing a Docker image with a container.
- Thinking Docker Hub and Docker are the same thing.
- Thinking `docker pull` starts a container.
- Thinking `docker run` starts an existing container.
- Confusing `docker run` with `docker start`.
- Confusing `docker stop` with `docker rm`.
- Confusing `docker rm` with `docker rmi`.
- Assuming deleting a container deletes the image.
- Storing important persistent data only inside the container writable layer.
- Assuming every container has Bash installed.
- Treating containers like virtual machines.
- Using destructive cleanup commands without understanding what they remove.

---

# 🎯 Interview Notes

## What Is a Docker Image?

A Docker image is a reusable, immutable application artifact containing the filesystem, application dependencies, configuration, and other information required to create a container.

---

## What Is a Container?

A container is a running or stopped instance created from a Docker image with its own runtime state and writable layer.

---

## Difference Between Image and Container

```text
Image
  |
  v
Reusable immutable artifact

Container
  |
  v
Runtime instance of the image
```

---

## What Is Docker Hub?

Docker Hub is a container registry used to store and distribute Docker images.

---

## Difference Between `docker run` and `docker start`

`docker run` creates a new container from an image and starts it.

`docker start` starts an existing container.

---

## Difference Between `docker stop` and `docker rm`

`docker stop` stops the container but keeps it.

`docker rm` removes the container.

---

## Difference Between `docker rm` and `docker rmi`

```text
docker rm
    |
    v
Removes container

docker rmi
    |
    v
Removes image
```

---

## What Happens When a Container's Main Process Exits?

The container normally transitions from the running state to a stopped/exited state.

---

## Why Are Containers Considered Ephemeral?

Containers are designed to be replaceable runtime units. Important persistent data should therefore be stored separately using appropriate persistent-storage mechanisms.

---

## What Is `docker exec`?

`docker exec` starts an additional process inside an already-running container.

Example:

```bash
docker exec -it my-redis redis-cli
```

---

## What Does `-p 8080:8080` Mean?

It publishes/maps:

```text
Host Port : Container Port
```

So:

```bash
-p 8080:8080
```

allows traffic reaching the host's port `8080` to be published to port `8080` of the container.

---

# 💡 Biggest Takeaways

Today fundamentally changed how I understand Docker.

Instead of seeing Docker as simply a tool that runs commands, I now understand the relationship between **images, containers, registries, Docker Engine, container processes, and container lifecycle**.

The biggest realization was that an image is a reusable artifact while a container is a runtime instance of that artifact. Multiple containers can share the same underlying image layers while maintaining their own runtime state.

Running Redis inside Docker also connected the concepts with real backend infrastructure. I was able to create a Redis container, interact with Redis using `redis-cli`, inspect its logs, stop it, start it again, and eventually remove the container.

Another important realization was understanding that containers should generally be treated as **disposable compute units**, while important persistent data should be handled separately.

---

# 📈 Progress

Completed:

- ✅ Docker Fundamentals
- ✅ Docker Images
- ✅ Docker Containers
- ✅ Docker Architecture
- ✅ Docker Hub & Container Registries
- ✅ Image vs Container
- ✅ Container Lifecycle
- ✅ First Real Docker Container
- ✅ Docker + Redis
- ✅ Essential Docker Commands

---

# 🔥 Looking Ahead

Tomorrow:

- Dockerfile
- Docker Build Context
- Dockerizing a Go Application
- Go Multi-Stage Docker Builds
- Docker Networking
- Docker Volumes
- Environment Variables
- Docker Compose
- Go API + PostgreSQL + Redis
- Production Docker Setup
- Docker + CI/CD

The next major milestone will be learning how to take a **Go backend and turn it into a production-quality Docker image**.

---

# 💭 Reflection

Day 55 was the beginning of an important shift from writing backend applications to understanding how backend applications are packaged, deployed, and operated.

Today I learned that Docker is not simply about running applications inside containers. It provides a complete model around **images, reproducible environments, container lifecycle, distribution, isolation, and deployment**.

The most valuable practical experience was running Redis inside Docker and interacting with it directly. It made the relationship between an image and a container much clearer than theory alone.

I now have a solid Docker foundation and understand the essential commands and concepts required to work with containers.

With the foundation complete, I am ready to move into the most important practical part of Docker for my Go backend journey: **writing Dockerfiles and containerizing my own Go applications.**
