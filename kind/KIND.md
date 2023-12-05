# KIND

## Instalação

### Osx

```bash
brew install kubectl
brew install kind
```

### Linux

```bash
curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.9.0/kind-linux-amd64
chmod +x ./kind
mv ./kind /some-dir-in-your-PATH/kind
```

```bash
curl -LO https://dl.k8s.io/release/v1.20.0/bin/linux/amd64/kubectl
chmod +x ./kubectl
sudo mv ./kubectl /usr/local/bin/kubectl
```

### Windows

```bash
curl.exe -Lo kind-windows-amd64.exe https://kind.sigs.k8s.io/dl/v0.9.0/kind-windows-amd64
Move-Item .\kind-windows-amd64.exe c:\some-dir-in-your-PATH\kind.exe
```

```bash
curl -LO https://dl.k8s.io/release/v1.20.0/bin/windows/amd64/kubectl.exe
Move-Item .\kubectl.exe C:\some-dir-in-your-PATH\kubectl.exe
```

```bash
choco install kubernetes-cli
choco install kind
```

### Docker

```bash
docker pull kindest/node:v1.20.0
```

## Cluster local

### Criando um cluster

```bash
kind create cluster --name kind-cluster
```

### Configurando o kubectl

```bash
kubectl cluster-info --context kind-kind-cluster
```

### Listando os containers do cluster

```bash
docker ps
```

### Listando os clusters

```bash
kind get clusters
```

### Listando os nodes

```bash
kubectl get nodes
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

## Criando um pod atraves do kubectl apply

### Criando um pod

```bash
kubectl apply -f k8s/pod.yaml

# Voce pode testar a aplicacao usando um port-forward
kubectl port-forward pod/goserver 8080:80

# Excluindo o pod
kubectl delete pod goserver
```
