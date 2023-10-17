# Imersão Full Stack && Full Cycle 15 - Microsserviço CodePix

## Tecnologias

- Docker
- Go Lang
- Apache Kafka
- PostgreSQL
- gRPC

## Docker

- Rodar o docker compose: docker compose up -d
- Entrar no container app: docker exec -it codepix-app-1 bash 

## Go

- Iniciar app Go: go mod init github.com/rodolfoHOk/fullcycle.imersao15/tree/main/codepix
- Baixar dependências: go mod tidy

## gRPC

- generate go codes: protoc --go_out=application/grpc/pb --go-grpc_out=application/grpc/pb --proto_path=application/grpc/protofiles application/grpc/protofiles/*.proto

## Cobra

- cobra-cli init
- cobra-cli add grpc
