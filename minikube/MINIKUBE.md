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

## Trabalhando com pods

### Criando nosso primeiro pod

```bash
# criando o pod
kubectl run nginx-pod --image=nginx:latest

# acompanhando o pod
kubectl get pods --watch

# verificando o pod
kubectl get pods

# verificando os detalhes do pod
kubectl describe pod nginx-pod
```

### Alterar a imagem de um pod (para uma imagem com a versão 1.0) [VAI DAR ERROR]

```bash
# alterando a imagem do pod
kubectl set image pod nginx-pod nginx-pod=nginx:1.0

# verificando os detalhes do pod
kubectl describe pod nginx-pod

# verificando o pod
kubectl get pods --watch
```

### Alterar a imagem de um pod (para uma imagem com a versão 1.9.1) [VAI DAR CERTO]

```bash
# alterando a imagem do pod
kubectl set image pod nginx-pod nginx-pod=nginx:1.9.1

# verificando os detalhes do pod
kubectl describe pod nginx-pod

# verificando o pod
kubectl get pods --watch
```

### Outra forma de editar a imagem de um pod

```bash
# definindo o editor padrao para o kubectl
export KUBE_EDITOR="nano"

# editando o pod com vscode
kubectl edit pod nginx-pod
```

### Criando o primeiro pod de maneira

- Crie o arquivo primeiro-pod.yaml
- Edite o arquivo primeiro-pod.yaml

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: primeiro-pod
  labels:
    app: primeiro-pod
spec:
  containers:
    - name: primeiro-pod
      image: nginx:latest
      ports:
        - containerPort: 80
```

- Criando o pod usando o arquivo primeiro-pod.yaml

```bash
kubectl apply -f ./minikube/k8s/primeiro-pod.yaml
```

### Apagando pods

```bash
# apagando o pod criado manualmente
kubectl delete pod nginx-pod

# apagando o pod criado de forma declarativa
kubectl delete -f ./minikube/k8s/primeiro-pod.yaml
```

### Criando um pod com uma aplicacao web

- Crie o arquivo portal-noticias-pod.yaml
- Crie o pod usando o arquivo portal-noticias-pod.yaml

```bash
# criando o pod
kubectl apply -f ./minikube/k8s/portal-noticias-pod.yaml

# verificando a porta do pod
kubectl get pods -o wide

# acessando o pod
kubectl exec -it pod/portal-noticias -- /bin/bash

# verificando o pod
curl localhost
```

### Criando outros pods e um servico para acesso ao pod

- Crie o arquivo pod-1.yaml
- Crie o arquivo pod-2.yaml
- Crie o arquivo svc-pod-2.yaml

```bash
# criando os pods
kubectl apply -f ./minikube/k8s/pod-1.yaml
kubectl apply -f ./minikube/k8s/pod-2.yaml\

# criando o servico
kubectl apply -f ./minikube/k8s/svc-pod-2.yaml

# verificando os pods
kubectl get pods -o wide

# verificando os servicos
kubectl get svc -o wide

# pegando o ip do servico
# acessando o pod
kubectl exec -it pod/pod-1 -- /bin/bash

# verificando o pod
curl <ip-servico>

# entrar no pod portal-noticias
kubectl exec -it pod/portal-noticias -- /bin/bash
```

- Testando o serviço com o pod apagado

```bash
# apagando o pod
kubectl delete pod pod-2

# verificando o dns dos pods
kubectl get pods -o wide --show-labels

# verificando o ip e o dns do servico
kubectl get svc -o wide --show-labels

# acessando o pod portal-noticias
kubectl exec -it pod/portal-noticias -- /bin/bash

# verificando o pod [OBS: o servico nao vai funcionar]
curl <ip-servico>

# recriando o pod
kubectl apply -f ./minikube/k8s/pod-2.yaml

# verificando o pod
kubectl get pods -o wide

# verificando o pod [OBS: o servico vai funcionar]
curl <ip-servico>
```

### Criando um servico NodePort

- Crie o arquivo svc-pod-1.yaml

```bash
# criando o servico
kubectl apply -f ./minikube/k8s/svc-pod-1.yaml

# pegando o ip do servico
kubectl get svc -o wide

# acessando o pod
kubectl exec -it pod/portal-noticias -- /bin/bash

# verificando o pod
curl <ip-servico>:<porta-servico>
```

- Acessando o servico pelo browser

```bash
# pegando o ip do node
kubectl get nodes -o wide
```

### Criando um serviço de load balancer

- Crie o arquivo svc-pod-1-lb.yaml

```bash
# criando o servico
kubectl apply -f ./minikube/k8s/svc-pod-1-loadbalancer.yaml
```

- Acessando o servico pelo browser

```bash
# pegando o ip do node
kubectl get nodes -o wide
```

## Exemplo do portal de noticias

### Criando os configmaps

```bash
# criando o pod
kubectl apply -f ./minikube/k8s/db-noticias-cm.yaml
kubectl apply -f ./minikube/k8s/portal-noticias-cm.yaml
kubectl apply -f ./minikube/k8s/sistema-noticias-cm.yaml

# pegando os servicos
kubectl get configmap
```

### Criando os pods

```bash
# criando o pod
kubectl apply -f ./minikube/k8s/db-noticias-pod.yaml
kubectl apply -f ./minikube/k8s/portal-noticias-pod.yaml
kubectl apply -f ./minikube/k8s/sistema-noticias-pod.yaml

kubectl get pods
```

### Criando os servicos

```bash
# criando os servicos
kubectl apply -f ./minikube/k8s/db-noticias-svc.yaml
kubectl apply -f ./minikube/k8s/portal-noticias-svc.yaml
kubectl apply -f ./minikube/k8s/sistema-noticias-svc.yaml

# pegando os servicos
kubectl get svc
```

### Expondo o servico sistema-noticias

```bash
# eu preciso expor o meu serviço para o meu host (macos)
# para isso vamos usar o kubectl port-forward
kubectl port-forward service/svc-sistema-noticias 30001:80
```

### Testando o servico

```bash
# pegando o ip do node
kubectl get nodes -o wide

# Executando o servico
minikube service --all

# acessando pelo browser
http://<ip-node>:<porta-servico>

# acessando o banco de dados
kubectl exec -it pod/pod-db-noticias -- /bin/bash
```
