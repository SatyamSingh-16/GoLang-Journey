# 🚀 GoLang Journey - Day 40

---

# 🏗️ Theme

Today marked another major milestone in my Go backend journey by moving beyond Docker fundamentals into actually **containerizing, running, debugging, and optimizing a Go backend application**.

Instead of treating Docker as just a collection of commands, I learned how Dockerfiles, images, containers, port mapping, container lifecycle, logs, debugging, and multi-stage builds fit together.

The focus shifted from simply running applications inside containers to understanding **how production-ready Docker images are built**, why Go applications are well suited for multi-stage builds, and why the build environment should be separated from the runtime environment.

---

# 🎯 Goal of the Day

Today's goal was to understand the complete Docker workflow for a Go backend application.

Instead of simply learning Docker commands, I learned how Docker builds images from Dockerfiles, how containers are created from images, how ports connect the host machine to containers, how containers are inspected and debugged, and why multi-stage builds are important for production applications.

By the end of the day, I understood the complete Docker application lifecycle:

```text
Dockerfile
    ↓
docker build
    ↓
Docker Image
    ↓
docker run
    ↓
Docker Container
    ↓
Go Application
I also understood why a Go compiler is required during the build stage but does not necessarily need to exist inside the final runtime image.
📚 Topics Covered
Dockerfile Fundamentals
Learned how a Dockerfile defines the instructions Docker follows to build an application image.
Covered:
Dockerfile
Base Image
WORKDIR
COPY
RUN
EXPOSE
CMD
Build Instructions
Runtime Instructions
Understood that a Dockerfile acts as a blueprint for creating a Docker image.
Docker Images
Studied Docker images as reusable application artifacts.
Covered:
Docker Images
Image Tags
Image Layers
Image References
Image Reusability
Build Context
Example:
docker build -t student-api:1.0 .
Learned that an image is not the running application itself. It is the packaged artifact from which containers are created.
Docker Containers
Learned the relationship between images and containers.
Covered:
Containers
Container Instances
Image → Container
Container Lifecycle
Running Containers
Stopped Containers
Understood the difference between:
Image
and:
Container
An image is the blueprint, while a container is an instance created from that blueprint.
Docker Build
Learned how Docker builds an image from a Dockerfile.
Used:
docker build -t student-api:1.0 .
Covered:
docker build
Image Naming
Image Tags
Build Context
Dockerfile Instructions
Image Creation
Understood that the . represents the current directory as the build context.
Docker Run
Learned how to create and start a container from an image.
Used:
docker run -p 8080:8080 student-api:1.0
Covered:
docker run
Container Creation
Container Startup
Port Mapping
Image Selection
Learned that docker run creates a new container from an image and starts it.
Port Mapping
Learned how Docker connects host machine ports to container ports.
Covered:
Host Port
Container Port
Port Mapping
-p
Backend Accessibility
Example:
-p 8080:8080
The format is:
HOST_PORT:CONTAINER_PORT
Therefore:
localhost:8080
      ↓
container:8080
      ↓
Go Server
This allows requests from the host machine to reach the Go application running inside the container.
Container Names
Learned how to assign meaningful names to containers.
Example:
docker run \
  --name student-api \
  -p 8080:8080 \
  student-api:1.0
Instead of relying on automatically generated names, the container can be referenced directly using:
student-api
This makes commands such as:
docker logs student-api
much easier to use.
Container Lifecycle
Studied the lifecycle of a Docker container.
Covered:
Created
Running
Stopped
Exited
Started Containers
Removed Containers
The basic lifecycle is:
Created
   ↓
Running
   ↓
Exited
   ↓
Removed
Learned that stopping a container does not automatically remove it.
Docker ps
Used:
docker ps
to inspect currently running containers.
Learned that this command primarily shows active containers.
Docker ps -a
Used:
docker ps -a
to inspect all containers.
Covered:
Running Containers
Stopped Containers
Exited Containers
Container Status
This became especially useful when debugging containers that had already stopped.
Docker Logs
Learned how to inspect application output from containers.
Used:
docker logs student-api
Also learned:
docker logs -f student-api
where -f continuously follows the container logs.
Covered:
Application Logs
Error Logs
Runtime Debugging
Log Following
Learned that container logs are one of the first places to look when a backend container fails.
docker stop
Used:
docker stop student-api
to stop a running container.
Learned that stopping a container does not delete it.
The stopped container can still be inspected using:
docker ps -a
docker start
Used:
docker start student-api
to start an existing stopped container.
Understood the difference between:
docker run
→ creates + starts a new container
and:
docker start
→ starts an existing container
docker rm
Used:
docker rm student-api
to remove a container.
Learned that removing a container and removing its image are two separate operations.
docker rmi
Used:
docker rmi student-api:1.0
to remove an image.
Understood the difference between:
docker rm
→ removes a container
and:
docker rmi
→ removes an image
docker exec
Learned how to execute commands inside a running container.
Example:
docker exec -it student-api sh
Covered:
docker exec
Interactive Mode
Terminal Access
Shell Access
Container Inspection
Understood that this allows us to enter a running container and inspect its runtime environment.
sh vs bash
Learned that Docker images do not necessarily contain Bash.
Some lightweight images provide:
sh
while other minimal images may not contain a shell at all.
This becomes particularly important when working with minimal runtime images such as:
scratch
🏗️ Multi-Stage Docker Builds
Introduced multi-stage Docker builds to solve an important production problem.
A normal Go build environment contains:
Go Compiler
Go Toolchain
Source Code
Dependencies
Build Tools
Compiled Binary
However, once the Go application has been compiled, the final application may only require the compiled binary and its runtime requirements.
Therefore, keeping the complete Go development environment inside the production image can be unnecessary.
Why Multi-Stage Builds Exist
Learned the fundamental purpose of multi-stage builds:
Build the application in one environment and run it in another environment.

The architecture becomes:
Builder Stage
      ↓
Go Compiler
      ↓
Go Source Code
      ↓
Compiled Binary
      ↓
Runtime Stage
      ↓
Run Application
This separates:
Build Environment
from:
Runtime Environment
and allows unnecessary build dependencies to stay outside the final image.
Builder Stage
Created a dedicated builder stage.
Example:
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o server .
The important artifact produced by this stage is:
server
Why Copy go.mod and go.sum Separately?
Learned why the Dockerfile uses:
COPY go.mod go.sum ./
before:
COPY . .
The dependency files are copied first so that Docker can cache the dependency installation layer.
Then:
RUN go mod download
downloads the Go dependencies.
After that:
COPY . .
copies the application source code.
This creates a more efficient build process because changing application source code does not necessarily require Docker to download all dependencies again.
The structure is:
go.mod + go.sum
       ↓
go mod download
       ↓
Dependencies Cached
       ↓
COPY Application
       ↓
go build
Runtime Stage
Created a second stage using another FROM.
Example:
FROM debian:bookworm-slim
The runtime stage is responsible only for running the application.
The final architecture becomes:
Builder Image
     ↓
Compile Application
     ↓
server binary
     ↓
Runtime Image
     ↓
Run server
COPY --from
Learned how to copy an artifact from the builder stage into the runtime stage.
Example:
COPY --from=builder /app/server /app/server
This means:
Builder Stage
/app/server
     ↓
Runtime Stage
/app/server
The final image receives the compiled application without needing the entire Go toolchain.
Complete Multi-Stage Dockerfile
Created a multi-stage Dockerfile following this architecture:
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o server .


FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/server /app/server

EXPOSE 8080

CMD ["./server"]
This separates the build environment from the runtime environment.
🐹 Why Go Works Well With Multi-Stage Builds
Learned that Go is particularly well suited for multi-stage Docker builds because Go applications can be compiled into binaries.
The process becomes:
Go Source Code
      ↓
Go Compiler
      ↓
Compiled Binary
      ↓
Minimal Runtime Image
The Go compiler is required during the build stage but does not necessarily need to be present in the final runtime image.
This makes Go a strong fit for lightweight production containers.
🏔️ Alpine
Learned about Alpine as a lightweight Linux distribution that can be used as a runtime environment.
Example:
FROM alpine:3.22
Covered:
Lightweight Runtime
Small Base Image
Linux Userland
Shell Availability
Runtime Dependencies
Learned that Alpine provides more functionality than scratch while still being significantly smaller than many traditional Linux distributions.
🪶 Scratch
Learned about:
FROM scratch
scratch represents an extremely minimal base image.
It provides essentially no normal Linux userland.
Therefore, it generally does not contain:
Shell
Package Manager
Debugging Utilities
Standard Linux Tools
This can produce extremely small images, but debugging and runtime configuration become more difficult.
⚙️ CGO_ENABLED and GOOS
Learned about:
CGO_ENABLED=0 GOOS=linux go build -o server .
Covered:
CGO_ENABLED
GOOS
Linux Builds
CGO
Portable Go Binaries
GOOS=linux tells Go to build the application for Linux.
CGO_ENABLED=0 disables CGO during compilation.
This can make it easier to produce a self-contained Go binary suitable for minimal runtime environments.
🧠 Static and Dynamic Linking
Learned that the runtime image must be selected based on the requirements of the compiled Go application.
A highly minimal image such as:
scratch
can work well for suitable self-contained binaries.
However, the smallest possible image is not automatically the best production choice.
Runtime dependencies, debugging requirements, certificates, timezone data, CGO requirements, and operational needs must all be considered.
🔍 Docker Debugging
Learned how to systematically debug Docker applications.
Started with:
docker ps
Then:
docker ps -a
Then:
docker logs student-api
And when required:
docker inspect student-api
This provides different levels of information about the container.
Debugging Application Crashes
If:
docker ps
does not show the expected container, check:
docker ps -a
If the container shows:
Exited
inspect its logs:
docker logs student-api
This helps distinguish between:
Docker Problem
and:
Application Problem
Debugging Port Problems
If:
localhost:8080
does not respond, check:
Is the container running?
Is the port mapped correctly?
Is the Go server actually running?
What do the container logs show?
Is the application listening on the expected port/interface?
Learned that a containerized API failing to respond does not automatically mean Docker itself is broken.
💻 Concepts Learned
Dockerfile
Docker Image
Docker Container
Docker Build
Docker Run
Build Context
Port Mapping
Container Lifecycle
Container Logs
docker ps
docker ps -a
docker logs
docker exec
docker stop
docker start
docker rm
docker rmi
Multi-Stage Builds
Builder Stage
Runtime Stage
COPY --from
Docker Layer Caching
go.mod
go.sum
go mod download
Alpine
Scratch
Go Binary
CGO_ENABLED
GOOS
Runtime Environment
Production Image Optimization
🧠 Important Concepts Learned
Dockerfiles define how application images are built.
Images are reusable artifacts from which containers are created.
Containers are runtime instances of images.
docker run creates and starts a new container.
docker start starts an existing container.
docker rm removes containers.
docker rmi removes images.
Port mapping connects host ports to container ports.
docker logs is one of the first tools to use when debugging a container.
docker ps -a is useful for investigating stopped containers.
go.mod and go.sum can be copied separately to improve Docker layer caching.
Multi-stage builds separate the build environment from the runtime environment.
COPY --from allows artifacts to be transferred between build stages.
Go's compiled binary model makes it well suited for multi-stage Docker builds.
scratch provides an extremely minimal runtime environment.
Alpine provides a lightweight Linux runtime with more functionality than scratch.
The smallest Docker image is not automatically the best production image.
Production containers should contain what the application actually requires at runtime rather than unnecessary build tooling.
⚠️ Common Mistakes I Learned
Confusing Docker images with containers.
Assuming docker ps shows stopped containers.
Forgetting to map ports.
Confusing docker run with docker start.
Confusing docker rm with docker rmi.
Assuming every Docker image contains Bash.
Forgetting that a container can exist even after it stops.
Re-downloading dependencies unnecessarily during every Docker build.
Copying the entire Go development environment into the final production image.
Assuming multi-stage builds are only about reducing image size.
Assuming scratch is automatically the best runtime image.
Forgetting that minimal images can be harder to debug.
Assuming Docker is always the source of an application failure.
🎯 Interview Notes
What Is the Difference Between an Image and a Container?
An image is a packaged application artifact used to create containers, while a container is a runtime instance created from that image.
What Is the Difference Between docker run and docker start?
docker run
→ creates + starts a new container

docker start
→ starts an existing container
What Is the Difference Between docker rm and docker rmi?
docker rm
→ removes a container

docker rmi
→ removes an image
Why Use Multi-Stage Docker Builds?
Multi-stage builds separate the build environment from the runtime environment.
This allows build tools, compilers, source code, and other unnecessary dependencies to remain outside the final runtime image.
Why Are Multi-Stage Builds Particularly Useful for Go?
Go applications can be compiled into binaries.
Therefore:
Go Toolchain
     ↓
Compile
     ↓
Binary
The final runtime image can focus on running the binary instead of containing the complete Go development environment.
Why Do We Copy go.mod and go.sum Before the Source Code?
Because Docker caches image layers.
By copying:
COPY go.mod go.sum ./
and downloading dependencies before:
COPY . .
the dependency layer can remain cached when only application source code changes.
This makes subsequent builds more efficient.
What Does COPY --from=builder Do?
It copies files from a previous Docker build stage into the current stage.
Example:
COPY --from=builder /app/server /app/server
This allows the final image to receive only the compiled application artifact.
Why Can scratch Be Difficult to Debug?
Because scratch does not provide a normal shell or common Linux utilities.
Therefore, entering the container and debugging it interactively is much harder than with images containing a normal Linux userland.
Why Isn't the Smallest Image Always the Best Image?
Production image selection depends on more than size.
You also need to consider:
Runtime dependencies
Security
Debuggability
Certificates
Timezone data
CGO requirements
Operational requirements
The best image is the smallest image that still provides everything the application actually needs.
💡 Biggest Takeaways
Today I learned that Docker is much more than simply running an application inside a container.
I developed a complete mental model of how a Go application moves from source code to a production-style container:
Dockerfile
     ↓
Docker Build
     ↓
Image
     ↓
Container
     ↓
Go Application
The most important concept of the day was multi-stage builds.
I understood why we need two different environments:
Builder Environment
        ↓
Go Compiler
        ↓
Compiled Binary
        ↓
Runtime Environment
        ↓
Application
The Go compiler is necessary to build the application, but the compiled binary does not necessarily need the entire Go development environment to run it.
This helped me understand the difference between what an application needs to be built and what an application needs to run.
I also learned why go.mod and go.sum are copied separately before the application source code. This was not just Docker syntax—it was about understanding Docker's layer caching mechanism and designing Dockerfiles that make repeated builds more efficient.
📈 Progress
Completed:
✅ Dockerfile Fundamentals
✅ Docker Images
✅ Docker Containers
✅ Docker Build
✅ Docker Run
✅ Port Mapping
✅ Container Names
✅ Container Lifecycle
✅ docker ps
✅ docker ps -a
✅ docker logs
✅ docker exec
✅ docker stop
✅ docker start
✅ docker rm
✅ docker rmi
✅ Docker Debugging
✅ go.mod and go.sum
✅ Docker Layer Caching
✅ Multi-Stage Builds
✅ Builder Stage
✅ Runtime Stage
✅ COPY --from
✅ Alpine
✅ Scratch
✅ Go Binary Optimization
✅ CGO_ENABLED
✅ GOOS
✅ Production Runtime Concepts
🔥 Looking Ahead
Next Steps:
Docker Compose
Multi-Container Applications
Docker Networking
Environment Variables
Persistent Volumes
Go API + PostgreSQL
Go API + Redis
Service-to-Service Communication
Containerized Backend Architecture
Production Docker Workflows
💭 Reflection
Day 40 was an important milestone in my Go backend journey because I moved beyond simply writing backend applications and started understanding how those applications are packaged, executed, debugged, and optimized in real environments.
I learned that Docker is not just about creating containers. It involves understanding the complete relationship between Dockerfiles, images, containers, ports, logs, lifecycle management, and runtime environments.
The biggest lesson of the day was multi-stage builds. A Go application needs the Go toolchain to build the application, but the compiled binary does not necessarily need the entire Go development environment to run it.
Understanding this distinction helped me see why production Docker images should contain only what is actually required at runtime.
I also learned why go.mod and go.sum are copied separately before the application source code. This was not just Docker syntax—it was about understanding Docker's layer caching mechanism and designing Dockerfiles that make repeated builds more efficient.
With Day 56 complete, I now have a strong foundation in Docker and understand how to containerize, run, inspect, debug, and optimize a Go backend application.
This gives me the foundation required to move into Docker Compose and multi-container backend architectures, where multiple services such as Go, PostgreSQL, and Redis can work together as a complete backend system. 🚀💙
```
