# 06 Container Image Optimization

In this section, we will explore the importance of optimizing container images. A well-optimized container image is smaller, faster to build, and more secure. We will discuss strategies to reduce image sizes, the differences between base images, the use of build stages, and how to manage data outside of the container image. This section also includes practical examples and exercises to help you apply these optimizations.

---

## Why Optimize Container Images?

Optimizing container images has several benefits:
- **Smaller Images**: Smaller images are faster to pull, push, and deploy.
- **Improved Security**: Smaller images reduce the attack surface by including only the necessary components.
- **Faster Builds**: Optimized images take less time to build and test.

---

## Choosing the Right Base Image

The base image is the foundation of your container. Choosing the right one can significantly impact the size and performance of your image.

### Common Base Images
1. **Ubuntu**:
    - A full-featured Linux distribution.
    - Larger in size (~29MB for minimal images).
    - Suitable for applications requiring a wide range of tools and libraries.

2. **Alpine**:
    - A minimal Linux distribution designed for containers.
    - Very small (~5MB).
    - Ideal for lightweight applications but may require additional setup for compatibility.

3. **Scratch**:
    - An empty base image.
    - Contains nothing—no shell, no utilities.
    - Best for statically compiled binaries or extremely minimal setups.

### Example: Using Alpine
```dockerfile
# Use Alpine as the base image
FROM alpine:3.18

# Install necessary dependencies
RUN apk add --no-cache curl

# Add your application
COPY app /app

# Run the application
CMD ["/app"]
```

### Exercise: Base Image Comparison
1. Create a Dockerfile using `Ubuntu` as the base image and build it.
2. Create another Dockerfile using `Alpine` as the base image and build it.
3. Compare the sizes of the resulting images using `docker images`.

---

## Using Build Stages

Build stages (multi-stage builds) allow you to separate the build process into multiple steps, reducing the final image size.

### Example: Multi-Stage Build
```dockerfile
# Stage 1: Build the application
FROM golang:1.20 AS builder
WORKDIR /app
COPY . .
RUN go build -o myapp

# Stage 2: Create a minimal runtime image
FROM scratch
COPY --from=builder /app/myapp /myapp
CMD ["/myapp"]
```

### Exercise: Multi-Stage Build
1. Write a multi-stage Dockerfile for the Go canvas client.
2. Use the first stage to install dependencies and build the app.
3. Use the second stage to copy only the necessary files to a minimal runtime image.

---

## Managing Data Outside the Container

To keep container images small and secure, avoid embedding data like configuration files, secrets, or large datasets directly in the image. Instead, use environment variables, secrets, and volumes.

### Example: Using Environment Variables
```dockerfile
FROM alpine:3.18
ENV APP_ENV=production
CMD ["sh", "-c", "echo Running in $APP_ENV mode"]
```

### Example: Using Volumes
```dockerfile
FROM nginx:1.25
# Use a volume to serve static files
VOLUME /usr/share/nginx/html
```

### Exercise: Externalizing Data
1. Create a Dockerfile that uses environment variables to configure an application.
2. Use Docker secrets to securely pass sensitive data to a container.
3. Mount a volume to share data between the host and the container.

---

## Key Takeaways
- Choose the smallest base image that meets your application's needs.
- Use multi-stage builds to reduce the size of your final image.
- Keep data outside the container image using environment variables, secrets, and volumes.

By following these practices, you can create efficient, secure, and maintainable container images.


## Investigate the Layers of the Go Container

Docker images are built in layers, with each instruction in the `Dockerfile` creating a new layer. Investigating these layers can help you understand how the image is constructed and identify opportunities to optimize its size.

### Steps to Investigate the Layers:

1. Use the `docker history` command to inspect the layers of the Go container:

    ```bash
    docker history go-canvas-client
    ```

    - Replace `go-canvas-client` with the name of your Go container image.
    - This command will display a list of layers, including their size, creation date, and the command that created each layer.

2. Analyze the output to identify the largest layers and understand their purpose. For example:
    - The base image layer (`golang:1.23-alpine`) will typically be the largest.
    - Layers created by `COPY` or `RUN` instructions in the `Dockerfile` will also be listed.

3. Compare the layers of the Go container with those of the Python container to see how the build process differs.

    ```bash
    docker history python-canvas-client
    ```

    - Notice how the Python container includes additional layers for installing dependencies and copying files.

### Key Notes:
- Each layer is cached by Docker, which speeds up subsequent builds if the instructions haven't changed.
- Minimizing the number of layers and the size of each layer can help reduce the overall image size.
- Use multi-stage builds to separate the build process from the final image, ensuring that only the compiled binary is included in the Go container.

By investigating the layers, you can gain insights into how your container is constructed and identify ways to optimize it for performance and size. 
