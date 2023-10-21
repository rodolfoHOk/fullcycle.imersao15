# Imersão Full Cycle & Full Stack 15 - CodePix

> Projetos da Imersão Full Cycle & Full Stack 15 CodePix da Full Cycle / Code Edu

## Sobre o Projeto

### Geral

![Sobre](/files/sobre.png)

### Bancos

![Sobre os Bancos](/files/sobre-bancos.png)

### CodePix

![Sobre os CoedPix](/files/sobre-codepix.png)

### Cadastro e consulta de chave Pix

![Cadastro de chave Pix](/files/sobre-cadastro-pix.png)

### Dinâmica do Processo

![Dinâmica do Processo](/files/sobre-processo.png)

### Principais Desafios

![Principais Desafios](/files/sobre-desafios.png)

### Wire Frame

![Wire Frame 1](/files/wireframe-1.png)

![Wire Frame 2](/files/wireframe-2.png)

![Wire Frame 3](/files/wireframe-3.png)

![Wire Frame 4](/files/wireframe-4.png)

![Wire Frame 5](/files/wireframe-5.png)

## CodePix (Go lang)

### Detalhes do CodePix

![Detalhes do CodePix](/files/sobre-detalhes-codepix.png)

### Arquitetura

![Arquitetura](/files/arquitetura.png)

### Estrutura do CodePix

![Estrutura do CodePix](/files/estrutura-codepix.png)

### Principais Tecnologias utilizadas no CodePix

- Docker
- Go Lang
- Apache Kafka
- PostgreSQL
- gRPC

## Bank (Nest Js)

### Visão geral dos projetos

![Visão geral](/files/geral.png)

### Operações

- Criar chave Pix
- Consultar chave Pix
- Realizar transferência por Pix
- Consultar o saldo

### Tecnologias

- Typescript / Javascript
- Nest.js
- TypeORM
- PostgreSQL
- Rest
- gRPC
- Kafka

## Bank Frontend (Next Js)

### Tecnologias

- Typescript / Javascript
- Next Js
- React Js
- React Server components
- Material UI

### Páginas

- Lista de contas do banco
- Home com informações da conta
- Criar e listar chaves Pix da conta
- Fazer transferência por Pix utilizando a conta atual

## Kubernetes

### Build docker images

- build codepix-go:

`docker build -t rodolfohok/codepix-go:latest -f codepix/Dockerfile.prod codepix`

- build bank-api:

`docker build -t rodolfohok/bank-api:latest -f bank/Dockerfile.prod bank`

### Kind

- website: https://kind.sigs.k8s.io/docs/user/quick-start
- create cluster: kind create cluster --name=codepix

### Helm

- website: https://helm.sh/docs/intro/install/
- charts: https://helm.sh/docs/topics/charts/
- bitnami charts: https://charts.bitnami.com/
- postgres helm chart: https://github.com/bitnami/charts/tree/main/bitnami/postgresql
- add repo: helm repo add bitnami https://charts.bitnami.com/bitnami
- install postgres: helm install postgres bitnami/postgresql
- env postgres password: export POSTGRES_PASSWORD=$(kubectl get secret --namespace default postgres-postgresql -o jsonpath="{.data.postgres-password}" | base64 -d)
- get postgres password: echo $POSTGRES_PASSWORD -> lg11Q0kmFY
- view postgres pode: kubectl get pods
- access postgres pod: kubectl run postgres-postgresql-client --rm --tty -i --restart='Never' --namespace default --image docker.io/bitnami/postgresql:16.0.0-debian-11-r13 --env="PGPASSWORD=$POSTGRES_PASSWORD" \
   --command -- psql --host postgres-postgresql -U postgres -d postgres -p 5432
- create databases: create database codepix; create database bank001; create database bank002;
- postgres host: postgres-postgresql.default.svc.cluster.local
- install kafka: helm install kafka bitnami/kafka

### kubectl

- website: https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/
- cluster info: kubectl cluster-info --context kind-codepix
- get nodes: kubectl get nodes
- docker list running containers: docker ps
- apply config map: kubectl apply -f configmap.yaml
- apply deployment: kubectl apply -f deploy.yaml
- get pods: kubectl get pods
- describe pod: kubectl describe pod codepix-66ddd97d75-ccfm4
- get pod logs: kubectl logs codepix-66ddd97d75-ccfm4
- apply secret: kubectl apply -f secret.yaml
- get services: kubectl get services
- delete pod: kubectl delete pod codepix-57964dd5d9-pwtqf
- apply service: kubectl apply -f service.yaml
- access pod: kubectl exec -it bankapi-8df94c6bc-6g87k bash
- port forward: kubectl port-forward svc/bankapi-service 8080:3000
