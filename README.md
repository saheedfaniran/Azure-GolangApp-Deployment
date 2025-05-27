# Azure Golang App Deployment with Azure DevOps

This guide explains how to build, containerize, and deploy a Golang web application to Azure Kubernetes Service (AKS) using Azure DevOps for CI/CD, Azure Container Registry (ACR), and Terraform for infrastructure provisioning.

---

## 📦 Prerequisites

Ensure you have the following:

- Azure Subscription
- Azure DevOps organization and project
- Azure CLI
- Docker
- Terraform
- kubectl
- Go (v1.18+)

---

## 🛠️ Step 1: Clone the Repository

```bash
git clone https://dev.azure.com/<your_org>/<your_project>/_git/Azure-GolangApp-Deployment
cd Azure-GolangApp-Deployment
```

---

## 🐳 Step 2: Build and Push Docker Image to ACR

### Authenticate to ACR

```bash
az acr login --name saheedacr110525
```

### Build the Image

```bash
docker build -t saheedacr110525.azurecr.io/saheedfaniran/azure-golangapp-deployment:24 .
```

### Push the Image

```bash
docker push saheedacr110525.azurecr.io/saheedfaniran/azure-golangapp-deployment:24
```

---

## ☁️ Step 3: Provision Infrastructure with Terraform

### Initialize and Apply Terraform

```bash
cd terraform
terraform init
terraform apply
```

This sets up AKS, ACR, and networking resources.

---

## 🔐 Step 4: Grant ACR Pull Access to AKS

```bash
az role assignment create   --assignee <AKS_MANAGED_IDENTITY_ID>   --role "AcrPull"   --scope $(az acr show --name saheedacr110525 --resource-group aks_tf_rg --query id --output tsv)
```

Replace `<AKS_MANAGED_IDENTITY_ID>` with the client or object ID of the AKS managed identity.

---

## 🚀 Step 5: Set Up Azure DevOps Pipeline

### Create a Pipeline

1. Navigate to your Azure DevOps project.
2. Go to **Pipelines > Create Pipeline**.
3. Select your repository (e.g., GitHub or Azure Repos).
4. Use the classic editor or YAML and add the following steps:

```yaml
trigger:
- main

pool:
  vmImage: 'ubuntu-latest'

variables:
  imageName: 'saheedacr110525.azurecr.io/saheedfaniran/azure-golangapp-deployment'

steps:
- task: GoTool@0
  inputs:
    version: '1.18'

- script: |
    go build -o app .
  displayName: 'Build Go app'

- task: Docker@2
  inputs:
    command: 'buildAndPush'
    repository: '$(imageName)'
    dockerfile: '**/Dockerfile'
    containerRegistry: '<AzureContainerRegistryServiceConnection>'
    tags: |
      24

- task: Kubernetes@1
  inputs:
    connectionType: 'Azure Resource Manager'
    azureSubscription: '<AzureSubscriptionServiceConnection>'
    azureResourceGroup: 'aks_tf_rg'
    kubernetesCluster: 'saheed-aks-cluster'
    namespace: 'default'
    command: apply
    useConfigurationFile: true
    configuration: 'golang-k8s-deployment.yaml'
```

Replace placeholders like `<AzureContainerRegistryServiceConnection>` and `<AzureSubscriptionServiceConnection>` with actual service connections defined in Azure DevOps.

---

## 📡 Step 6: Access the Application

After deployment, fetch the external IP:

```bash
kubectl get svc golang-app-svc
```

Access the app at `http://<external-ip>:1337`

---

## 🧪 API Endpoints

| Method | Endpoint     | Description         |
|--------|--------------|---------------------|
| GET    | `/hello`     | Hello world message |
| GET    | `/greetings` | Custom greetings    |

---

## 📄 License

MIT

---

## 👨🏽‍💻 Author

Saheed Faniran
