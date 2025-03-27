# On premise deployment

Containers are not only for cloud. Some customers don't want to deploy their applications to a cloud, for example for security reasons, or for cost considerations. Your container can run anywhere.

## Introduction to Hashicorp Nomad

Hashicorp Nomad is a flexible, easy-to-use orchestrator that enables organizations to deploy and manage applications across on-premise and cloud environments. It supports containerized workloads, such as Docker, as well as non-containerized applications, making it a versatile choice for diverse deployment needs.

Nomad is designed to be lightweight and efficient, with a single binary that simplifies installation and reduces operational complexity. It integrates seamlessly with other Hashicorp tools like Consul for service discovery and Vault for secrets management, providing a comprehensive ecosystem for modern application deployment.

With its focus on simplicity, scalability, and multi-cloud support, Nomad is an excellent choice for organizations looking to manage their workloads in a consistent and reliable manner. It is a nice alternative to Kubernetes.

## Run the container on Hashicorp Nomad

### prerequisites for deploying a container to Hashicorp Nomad:
1. **Nomad server address**: The server address is needed to access the webinterface.
2. **Client secret**: A client secret is required to login to user interface of Nomad

### Step-by-Step Guide to Deploy a Container on Hashicorp Nomad

Application deployments are defined as Job files in Nomad. These files are written in the `hcl` language. HCL, or HashiCorp Configuration Language, is a domain-specific language created by HashiCorp. It is designed to be both human-readable and machine-friendly, making it ideal for defining infrastructure, configurations, and workflows in HashiCorp tools like Terraform, Nomad, and Consul.

1. **Create a Nomad Job File**  
    Define a job file to specify the container you want to deploy. Create a file named `canvas-client.hcl` with the following content:
    ```hcl
        job "canvas-client-yourname" {
            # Specifies the datacenter where this job should be run
            # This can be omitted and it will default to ["*"]
            datacenters = ["*"]


            # A group defines a series of tasks that should be co-located
            # on the same client (host). All tasks within a group will be
            # placed on the same host.
            group "canvas-client" {

                # Specifies the number of instances of this group that should be running.
                # Use this to scale or parallelize your job.
                # This can be omitted and it will default to 1.
                count = 1

                # Tasks are individual units of work that are run by Nomad.
                task "canvas-client" {
                    # This particular task starts a simple web server within a Docker container
                    driver = "docker"

                    config {
                        image   = "docker.io/yournamespace/yourimage:tag"
                        args    = ["100", "100", "banana.png"]
                    }

                    env {
                        MY_KEY = "my_value"
                    }

                    # Specify the maximum resources required to run the task
                    resources {
                        cpu    = 50
                        memory = 64
                    }
                }
            }
        }
    ```
    Ensure you add `your name` to the job name, so it doesn't conflict with other students.

    Replace `yournamespace`, `yourimage` and the `tag` with your container image that you pushed to Docker hub.

2. **Run the Job**  
    * Log in the webinterface.
    * In the left bar click `Jobs`
    * In the top right, click the `Run Job` button
    * Use the `Upload file` button to upload your `hcl` file
    * Click `Plan` to review your job
    * Then click `Run` to submit your job

3. **Verify the Deployment**  
    You will find out that your job will not start. The status keeps `Deploying`.

    Try to find the logs of the `Allocation` so find out what the issue is.

    Few hints:
    * Ensure that the picture you want to upload is added to the container
    * You might need to adjust the Python script or Go application so that it keeps running


