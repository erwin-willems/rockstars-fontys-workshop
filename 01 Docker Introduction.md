# Docker introduction

In this introduction you will learn the basic steps to create a Docker container. A container is described in a Dockerfile. 

We've prepared a few Dockerfiles already in the Python folder. A container needs to be build before you can run it. 

Go in to the Python folder and open the Dockerfile. We will inspect this file line by line:

```dockerfile
FROM python:3.13-alpine

WORKDIR /app

COPY requirements.txt .
COPY main.py .

RUN pip install -r requirements.txt

ENTRYPOINT ["python", "main.py"]
```

## Explanation of Each Line:

1. **`FROM python:3.13-alpine`**  
   - This specifies the base image for the container. Every container starts with a base image.
   - The very first container is `scratch`, which is a special container. It's the smallest container possible (we will learn more about that later).
   - `python:3.13-alpine` is a lightweight Python image based on the Alpine Linux distribution.
   - Using Alpine reduces the image size, making it faster to download and deploy.

2. **`WORKDIR /app`**  
   - Sets the working directory inside the container to `/app`.  
   - All subsequent commands (e.g., `COPY`, `RUN`) will be executed relative to this directory.  
   - This helps keep the container organized.

3. **`COPY requirements.txt .`**  
   - Copies the `requirements.txt` file from your local machine to the container's `/app` directory.  
   - This file typically lists the Python dependencies for your application.
   
4. **`RUN pip install -r requirements.txt`**  
   - Installs the Python dependencies listed in `requirements.txt` using `pip`.
   - This ensures the container has all the necessary libraries to run your application.

5. **`COPY main.py .`**
   - Copies the Python script from your local project directory (where the `Dockerfile` resides) to the container's `/app` directory.  

6. **`ENTRYPOINT ["python", "main.py"]`**  
    - Specifies the default executable to run when the container starts.  
    - In this case, it ensures the Python interpreter runs the `main.py` script.  
    - `ENTRYPOINT` is used to define the main process of the container, allowing additional arguments to be passed during runtime.

### Key Notes:
- **Order Matters**: Docker caches each instruction as a layer. Placing frequently changing instructions (like `COPY main.py .`) later in the file can optimize build times.
- **Alpine Base**: While Alpine is lightweight, it may require additional libraries for some Python packages. If you encounter issues, consider using a non-Alpine base image like `python:3.13-slim`.

## Build the container

To build the container, use the `docker build` command. This command reads the `Dockerfile` in the current directory and creates a Docker image based on its instructions.

### Steps to Build the Container:

1. Open a terminal and navigate to the `Python` folder where the `Dockerfile` is located.
2. Run the following command to build the Docker image:

    ```bash
    docker build --tag python-canvas-client .
    ```

    - `--tag python-canvas-client`: Assigns a name (`python-canvas-client`) to the image.
    - `.`: Specifies the current directory as the build context.

3. Once the build is complete, verify that the image was created by running:

    ```bash
    docker images
    ```

    You should see `python-canvas-client` listed in the output.

### Key Notes:
- Ensure that the `Dockerfile`, `requirements.txt`, and `main.py` are in the same directory before running the build command.
- If you encounter errors during the build, check the `Dockerfile` for typos or missing files.
- Use descriptive tags (e.g., `python-canvas-client:v1`) to version your images for better management.

### Update the container:

## Update the container:

To observe how Docker optimizes the build process, let's make a small update to the `main.py` file.

### Steps to Update and Rebuild:

1. Open the `main.py` file in the `Python` folder and make a small change. For example, just add a newline in the end.

2. Save the file and rebuild the Docker image using the same command:
    ```bash
    docker build --tag python-canvas-client .
    ```

3. Notice that Docker reuses the cached layers for all instructions up to the `COPY main.py .` step. Only the layers after this step are rebuilt.

4. Verify the updated image by running:
    ```bash
    docker run --rm python-canvas-client
    ```

    You should see the updated output from the `main.py` script.

### Key Notes:
- Docker's layer caching mechanism speeds up rebuilds by only processing the layers affected by changes.
- To optimize build times, place frequently changing instructions (like `COPY main.py .`) towards the end of the `Dockerfile`.
- Use the `docker history python-canvas-client` command to inspect the image layers and understand which ones were rebuilt.

## Run the container

### Steps to Run the Container:

1. Use the `docker run` command to start a container from the image you built:

    ```bash
    docker run --name python-container --rm python-canvas-client
    ```

    - `--name python-container`: Assigns a name (`python-container`) to the running container.
    - `--rm`: Automatically removes the container once it stops.
    - `python-canvas-client`: Specifies the image to use for the container. It uses the container that we build earlier.

2. The container will execute the `CMD` instruction from the `Dockerfile` and run the `main.py` script.

3. You should see the output of the `main.py` script in your terminal.

### Key Notes:
- Use `docker ps` to list running containers and `docker ps -a` to list all containers (including stopped ones).
- To stop a running container manually, use `docker stop <container_name>`.

### Pass the API_KEY variable to the container

The container does not work because we forgot to pass the `API_KEY` environment variable. To pass the `API_KEY` environment variable to the container, you can use the `--env` flag with the `docker run` command.

### Steps to Pass the API_KEY:

1. Run the container with the `--env` flag to specify the `API_KEY` environment variable (you will find the API key in the introduction):

    ```bash
    docker run --name python-container --rm --env API_KEY=your_api_key_here python-canvas-client
    ```

    - Replace `your_api_key_here` with the actual API key value.
    - The `--env API_KEY=your_api_key_here` flag sets the `API_KEY` environment variable inside the container.

2. The application inside the container will now have access to the `API_KEY` variable and should function as expected.

3. Verify the output of the `main.py` script to ensure the application is using the provided API key.

### Key Notes:
- Environment variables are a secure way to pass sensitive information like API keys to containers.
- Avoid hardcoding sensitive data in your `Dockerfile` or source code. That's why we use environment variables instead.
- If you need to pass multiple environment variables, you can use multiple `--env` flags or an `.env` file with Docker Compose.
- Ensure that you pass `--env` before the image name. Every flag after the image name will be interpreted as arguments / flags of the application

### Pass arguments to the application

### Steps to Pass Arguments to the Application:

1. Use the `docker run` command to pass arguments to the application. Arguments are specified after the image name and are passed directly to the `CMD` instruction in the `Dockerfile`.

    ```bash
    docker run --name python-container --rm --env API_KEY=your_api_key_here python-canvas-client 100 200 banana.png
    ```

    - Replace `100` (x), `200` (y), and `image_path` with the actual values you want to pass to the application.
    - These arguments will be available to the `main.py` script as command-line arguments.

You will discover that the image name can not be found. You will get this error:

```
FileNotFoundError: [Errno 2] No such file or directory: 'banana.png'
```

### Key Notes:
- Ensure that the `main.py` script is designed to handle the arguments properly.
- Arguments passed after the image name are treated as input to the application and are appended to the `ENTRYPOINT` instruction in the `Dockerfile`.


### Add the File to the Container Using `COPY`

The application fails to locate the image file specified in the `image_path` argument because the file is not included in the Docker image. To include the image file in the container, you can use the `COPY` instruction in the `Dockerfile`. This approach ensures the file is bundled into the container during the build process.

### Steps to Add the File:

1. Place the image file (e.g., `banana.png`) in the same directory as the `Dockerfile`.

2. Update the `Dockerfile` to include the `COPY` instruction for the image file. Ensure it's added before the `ENTRYPOINT`:

    ```dockerfile
    COPY banana.png .
    ```

    - This copies the `banana.png` file from your local directory to the `/app` directory inside the container.

3. Rebuild the Docker image to include the updated `Dockerfile`:

    ```bash
    docker build --tag python-canvas-client .
    ```

4. Run the container and pass the updated path to the image file as an argument:

    ```bash
    docker run --name python-container --rm --env API_KEY=your_api_key_here python-canvas-client 100 200 /app/banana.png
    ```

    - The `/app/banana.png` path corresponds to the location of the image file inside the container.

5. Verify that the application processes the image file correctly.

### Key Notes:
- Adding files to the container using `COPY` is useful for static assets that do not change frequently.
- Ensure that the file is in the same directory as the `Dockerfile` or adjust the path in the `COPY` instruction accordingly.
- Avoid adding large or unnecessary files to the container to keep the image size small.
- Use descriptive paths inside the container to keep the file structure organized.


### Mounting Volumes to Pass Files

Although the `COPY` solution works, it's not very flexible. If you want to upload another image, you need to modify the Dockerfile and rebuild the container again. To resolve this, you can mount a volume to make the file accessible to the container.

### Steps to Mount a Volume:

1. First remove the `COPY` statement we added to the `Dockerfile` and rebuild the container
2. Use the `--volume` flag with the `docker run` command to mount a directory or file from your host machine into the container.

    ```bash
    docker run --name python-container --rm --env API_KEY=your_api_key_here --volume /path/to/your/image:/app/image.jpg python-canvas-client 100 200 /app/image.jpg
    ```

    - Replace `/path/to/your/image` with the full path to the image file on your host machine.
    - The `-v /path/to/your/image:/app/image.jpg` flag maps the file from your host machine to the `/app/image.jpg` path inside the container.
    - Update the `image_path` argument to match the path inside the container (`/app/image.jpg` in this example).

3. The container will now have access to the image file, and the application should function as expected.

4. Verify the output of the `main.py` script to ensure the application processes the image correctly.

### Key Notes:
- Use absolute paths for the host file or directory to ensure proper mounting.
- The `-v` flag can also be used to mount entire directories if your application requires access to multiple files.
- Avoid hardcoding file paths in your application. Instead, use arguments or environment variables to specify file locations.
- Ensure that the mounted file or directory has the necessary permissions for the container to access it.
- If you encounter issues, check the container logs using `docker logs <container_name>` for debugging.

By mounting the required file as a volume, you can keep your Docker image lightweight and avoid bundling unnecessary files into the image.

## Create a container from the Go project

Go source code needs to be compiled. Once compiled you can run the binary. So you don't need Go to run, like you need Python to run a python script.

1. Using the following information, create a Dockerfile and build the container using the source code in the Go folder:

   - Use the base image: `docker.io/golang:1.23-alpine`
   - Add the files `go.mod` and `main.go` to the container
   - Run the build proces. The full command for that is: `go build -o shared-canvas-client main.go`
   - Then use the output binary `shared-canvas-client` as entrypoint

Of course, you can use the Dockerfile from Python as a reference.

## Run the container

Running the container should work pretty much the same as the Python container. And that's nice about containers. The end-users should not see much difference
in which programming language is used to create the container.

1. Try to run the container you've created from the Go directory.

## See the difference in size between the Python and Go Container

One of the key advantages of using Go is the ability to produce small, self-contained binaries. This often results in smaller container images compared to Python-based containers, which require a runtime and additional dependencies.

### Steps to Compare Container Sizes:

1. List all the Docker images on your system using the following command:

    ```bash
    docker images
    ```

    This will display a table with the repository name, tag, image ID, creation date, and size.

2. Look for the sizes of the Python container (`python-canvas-client`) and the Go container (`go-canvas-client`) in the output.

3. Compare the sizes of the two images. You should notice that the Go container is significantly smaller than the Python container.

Ok, that's weird... The Go container should be much smaller than the Python container, but did you see it isn't? No worries, in Chapter 06 we will investigate why
and how we can reduce the Go container. 

### Key Notes:
- The Python container includes the Python runtime and any additional dependencies specified in `requirements.txt`, which can increase the image size.
- The Go container, on the other hand, only includes the compiled binary and its minimal dependencies, resulting in a smaller image.
- Using lightweight base images like `alpine` further reduces the size of both containers.

By comparing the sizes, you can see how the choice of programming language and runtime affects the overall container footprint.
