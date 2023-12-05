# MINIKUBE

## Instalação

### Osx

```bash
brew install kubectl
brew install minikube
```

### Linux

```bash
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
```

```bash
curl -LO https://dl.k8s.io/release/v1.20.0/bin/linux/amd64/kubectl
chmod +x ./kubectl
sudo mv ./kubectl /usr/local/bin/kubectl
```

### Windows

```bash
curl -LO https://dl.k8s.io/release/v1.20.0/bin/windows/amd64/kubectl.exe
Move-Item .\kubectl.exe C:\some-dir-in-your-PATH\kubectl.exe
```

```bash
choco install kubernetes-cli
choco install minikube
```

## Cluster local

### Iniciando o cluster

```bash
# criando o cluster com docker, versao do kubernetes e nome do cluster
minikube start --driver=docker --kubernetes-version=latest --vm=true

# removendo o cluster
minikube delete
```

### Adicionando alguns addons ao cluster

```bash
minikube addons enable metrics-server # kubectl top nodes
minikube addons enable dashboard # dashboard
minikube addons enable ingress # nginx ingress controller
```

### Configurando o kubectl

```bash
kubectl cluster-info --context minikube
```

### Listando os containers do cluster

```bash
docker ps
```

### Verificando o status do cluster

```bash
minikube status
```

### Abrindo o dashboard do cluster

```bash
minikube dashboard --url
```

### Excluindo um cluster

```bash
kind delete cluster --name kind-cluster
```

## Recriando o cluster com o arquivo kind.yaml

### Crie o arquivo kind.yaml

```yaml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
  - role: control-plane
  - role: worker
  - role: worker
  - role: worker
```

### Crie o cluster

```bash
kind create cluster --name kind-cluster --config k8s/kind.yaml
```

### Repetir os passos anteriores

```bash
kubectl cluster-info --context kind-kind-cluster
docker ps
kind get clusters
kubectl get nodes
```

## Trabalhando em diferentes contextos

### Listando os contextos

```bash
kubectl config get-contexts
```

### Trocando de contexto

```bash
kubectl config use-context kind-kind-cluster
```

### Listando os nodes do contexto

```bash
kubectl get nodes
```

### Dica com vscode

```bash
code --install-extension ms-kubernetes-tools.vscode-kubernetes-tools
```
