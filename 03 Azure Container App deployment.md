# Azure Container App Deployment

Deploying a container in Azure requires a few key components and steps to ensure a smooth and secure deployment process. This document will guide you through the process of deploying a container image to an Azure Container Instance (ACI). 

### Requirements for Deploying a Container in Azure:
1. **Azure Subscription**: You need an active Azure subscription to create and manage resources.
2. **Service Principal**: A service principal is required for authentication and authorization to interact with Azure resources programmatically.
3. **Azure Container Registry (ACR)**: ACR is used to store and manage container images securely.
4. **Container Image**: A pre-built container image that contains your application and its dependencies.
5. **Azure CLI**: The Azure Command-Line Interface (CLI) is used to interact with Azure resources.
6. **Docker CLI**: The Docker CLI is used to build, tag, and push container images to the ACR.

---

## Step 1: Log in to Azure Using a Service Principal

Use the Azure CLI to log in with the service principal:

```bash
# Azure configuration
AZURE_SUBSCRIPTION_ID="0dc77b19-1724-4a4c-8e4e-73625f11af59"
AZURE_TENANT_ID="99b9369b-ccb7-4eb1-ada0-f8efa1b0044a"
AZURE_CLIENT_ID="<your-service-principal-client-id>"
AZURE_CLIENT_SECRET="<your-service-principal-client-secret>"

az login --service-principal \
    --username $AZURE_CLIENT_ID \
    --password $AZURE_CLIENT_SECRET \
    --tenant $AZURE_TENANT_ID

az account set --subscription $AZURE_SUBSCRIPTION_ID
```

### Explanation:
1. **`az login --service-principal`**: Logs in to Azure using a service principal.
2. **`--username`**: Specifies the service principal client ID.
3. **`--password`**: Specifies the service principal client secret.
4. **`--tenant`**: Specifies the Azure tenant ID.
5. **`az account set`**: Sets the active Azure subscription.

---

## Step 2: Create the Resource Group and Set Up Azure Container Registry

Before deploying your container, you need to create a resource group and set up an Azure Container Registry (ACR) to store your container images.

### Create a Resource Group

```bash
# Resource group configuration, make it unique, like adding your name
RESOURCE_GROUP="fontys-container-workshop-yourname"
LOCATION="westeurope" # Best suited Azure region for the Netherlands

az group create --name $RESOURCE_GROUP --location $LOCATION
```

### Explanation:
1. **`az group create`**: Creates a new resource group.
2. **`--name`**: Specifies the name of the resource group.
3. **`--location`**: Specifies the Azure region where the resource group will be created.

---

### Set Up Azure Container Registry (ACR)

```bash
# ACR configuration, make it unique, like adding your name
ACR_NAME="fontysacr-your-name"

az acr create \
  --resource-group $RESOURCE_GROUP \
  --name $ACR_NAME \
  --sku Basic \
  --location $LOCATION \
  --admin-enabled true
```

### Explanation:
1. **`az acr create`**: Creates a new Azure Container Registry.
2. **`--resource-group`**: Specifies the resource group where the ACR will be created.
3. **`--name`**: Specifies the name of the ACR.
4. **`--sku Basic`**: Specifies the pricing tier for the ACR.
5. **`--location`**: Specifies the Azure region where the ACR will be created.
6. **`--admin-enabled true`**: Enables authentication credentials

---

## Step 3: Push an Existing Image to ACR

If you have already built a container image earlier in the workshop, you can push it to the ACR by tagging it with the ACR's URL and then pushing it.

```bash
# ACR configuration
ACR_URL="${ACR_NAME}.azurecr.io"
CONTAINER_IMAGE="python-canvas-client"

docker tag $CONTAINER_IMAGE $ACR_URL/$CONTAINER_IMAGE

az acr login --name $ACR_NAME

docker push $ACR_URL/$CONTAINER_IMAGE
```

### Explanation:
1. **`docker tag`**: Tags the existing image with the ACR's URL.
2. **`az acr login`**: Logs in to the ACR to allow pushing images.
3. **`docker push`**: Pushes the tagged image to the ACR.

---

## Step 4: Deploy the Container to Azure Container Instance

Deploy the container instance:

```bash
# Deployment configuration
CONTAINER_NAME="fontys-demo-student01"
CONTAINER_PORT=80 # Replace with your container's exposed port
ACR_PASSWORD=$(az acr credential show --name $ACR_NAME --query "passwords[0].value" -o tsv)

az container create \
  --resource-group $RESOURCE_GROUP \
  --name $CONTAINER_NAME \
  --image $ACR_URL/$CONTAINER_IMAGE \
  --registry-username $ACR_NAME \
  --registry-password $ACR_PASSWORD \
  --dns-name-label $CONTAINER_NAME \
  --os-type Linux \
  --cpu 1 \
  --memory 1.5 \
  --restart-policy Always
```

### Explanation:
1. **`az container create`**: Creates a new Azure Container Instance.
2. **`--resource-group`**: Specifies the resource group for the container.
3. **`--name`**: Specifies the name of the container.
4. **`--image`**: Specifies the container image from the Azure Container Registry.
5. **`--registry-username`**: Specifies the username for the Azure Container Registry.
6. **`--registry-password`**: Specifies the password for the Azure Container Registry.
7. **`--dns-name-label`**: Sets a DNS name label for the container.
8. **`--os-type`**: Specifies the operating system type (Linux).
9. **`--command-line`**: Runs a command to keep the container running.
10. **`--cpu`**: Allocates 1 CPU core to the container.
11. **`--memory`**: Allocates 1.5 GB of memory to the container.
12. **`--restart-policy`**: Sets the restart policy to always restart the container.

---

## Step 5: Verify the Container is Running

Check the status of the container:

```bash
az container show \
    --resource-group $RESOURCE_GROUP \
    --name $CONTAINER_NAME \
    --query "instanceView.state"
```

Retrieve the public IP address of the container:

```bash
az container show \
    --resource-group $RESOURCE_GROUP \
    --name $CONTAINER_NAME \
    --query "ipAddress.ip" \
    --output tsv
```

### Explanation:
1. **`az container show`**: Displays details of the container instance.
2. **`--query`**: Queries the state or public IP address of the container.
3. **`--output tsv`**: Outputs the result in plain text format.

---

## Step 6: Track Logs for the Running Container

View the logs of the container:

```bash
az container logs \
    --resource-group $RESOURCE_GROUP \
    --name $CONTAINER_NAME
```

Stream the logs in real-time:

```bash
az container attach \
    --resource-group $RESOURCE_GROUP \
    --name $CONTAINER_NAME
```

### Explanation:
1. **`az container logs`**: Fetches the logs of the container instance.
2. **`az container attach`**: Attaches to the container instance to stream logs in real-time.

---

By following these steps, you can deploy and manage your containerized application in Azure Container Instances using both Azure CLI and Docker CLI.