# Container registries

Container registries are centralized repositories used to store, manage, and distribute container images. They play a crucial role in containerized application development by providing a secure and efficient way to share images across teams and environments. Popular container registries include Docker Hub, Azure Container Registry, and Amazon Elastic Container Registry. By using a container registry, developers can ensure version control, scalability, and seamless integration into CI/CD pipelines.

## Create an account on Docker Hub

To create an account on Docker Hub, follow these steps:

1. Visit the Docker Hub website  
    Go to [https://hub.docker.com/](https://hub.docker.com/) in your web browser.

2. Click on `Sign Up`  
    On the homepage, click the `Sign Up` button to begin the registration process.

3. Ensure you select `Personal` and fill in the Registration Form  
    Provide the required details, including your username, email address, and password. Make sure to choose a strong password.

4. Verify Your Email
    After submitting the form, check your email inbox for a verification email from Docker Hub. Click the verification link to activate your account.

5. Log In to Docker Hub
    Once your account is verified, log in to Docker Hub using your credentials.

By completing these steps, you will have a Docker Hub account ready to use for managing and sharing container images.



## Tag the container
To tag a container image, follow these steps:

1. Identify the Image  
    Use the `docker images` command to list all available images on your local machine and identify the image you want to tag:
    ```bash
    docker images
    ```

2. Tag the Image  
    Use the `docker tag` command to assign a new tag to the image. The format is `source-image:tag target-image:tag`:
    ```bash
    docker tag <source-image>:<source-tag> <username>/<repository>:<new-tag>
    ```
    Example:
    ```bash
    docker tag my-app:latest myusername/python-canvas-client:1.0
    ```

By tagging your container image, you prepare it for uploading to a container registry like Docker Hub or for easier identification in your local environment.

### Key Notes

- Use clear and descriptive names for your repositories and tags to make it easier to identify and manage your images.
- The tag can be used as version for your image. So you can also create a tag `python-canvas-client:2.0`


## Pushing a Container to Docker Hub

To push a container image to Docker Hub, follow these steps:

1. Log in to Docker Hub  
    Use the `docker login` command to authenticate with your Docker Hub account:
    ```bash
    docker login
    ```
    Enter your Docker Hub username and password when prompted.

2. Push the Image  
    Use the `docker push` command to upload the image to Docker Hub:
    ```bash
    docker push <username>/<repository>:<tag>
    ```
    Example:
    ```bash
    docker push myusername/python-canvas-client:1.0
    ```

3. Verify the Image on Docker Hub  
    Log in to your Docker Hub account via the web interface and navigate to your repository to confirm the image has been successfully uploaded.

By following these steps, you can easily share your container images with others or use them in your deployment pipelines.

### Key Notes

- Docker Hub allows you to create both public and private repositories. Public repositories are accessible to everyone, while private repositories are restricted to authorized users.

## Pull a container from Docker Hub

To pull a container image from Docker Hub, follow these steps:

1. Pull the Image  
    Use the `docker pull` command to download the image to your local machine:
    ```bash
    docker pull <username>/<repository>:<tag>
    ```
    Example:
    ```bash
    docker pull myusername/python-canvas-client:1.0

2. Try to pull images from other students

3. Verify the Image  
    After pulling the image, use the `docker images` command to confirm it is available locally:
    ```bash
    docker images
    ```

### Key Notes

- If no tag is specified, Docker will pull the `latest` tag by default.
- Ensure you trust the source of the image before pulling it, especially for public repositories (or other students :) ).
- You can use the pulled image to create and run containers on your local machine or in your deployment environment.
- Always keep your images up to date by pulling the latest versions when necessary.
