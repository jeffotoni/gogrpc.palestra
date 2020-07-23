# Poc Palestra

Este repo são exemplos que foram demostrados na palestra Grpc e GO.

Temos exemplos de APIS, RPC e GRPC.

 - No diretório Apis temos um rEST e seu Dockerfile.
 - No diretório gRPC temos nosso famigerado exemplo rodando gRPC.
 - No diretório RPC temos nosso server e client RPC.

Grpc é um framework.

  - [helloworld](https://godoc.org/google.golang.org/grpc/examples/helloworld/helloworld)
  - [Quick Start](https://grpc.io/docs/languages/go/quickstart/)
  - [Basics Tutorial](https://grpc.io/docs/languages/go/basics/)
  - [Grpc Go pkg](https://godoc.org/google.golang.org/grpc)

## Instalar proto

Primeiro passo é instalar nosso protobuf e entender como ele funciona.

[proto v1.3]
 ```bash
 $ go get github.com/golang/protobuf/protoc-gen-go@v1.3
```

[proto gen]
 ```bash
 $ go get -u github.com/golang/protobuf/protoc-gen-go
```

## Ferramentas para Client grpc

 - [grpcurl](https://github.com/fullstorydev/grpcurl)
 - [bloomrpc](https://github.com/uw-labs/bloomrpc)

### API

```bash

$ curl -i -X POST localhost:8080/auth

```

```bash

$ curl -i -X POST localhost:8080/user

```

```bash

$ curl -i -X GET localhost:8080/user/uuid

```

#### Executando API
```go
$ go run main.go
```

#### Compilando API
```go

$ CGO_ENABLED=0 go build -ldflags="-s -w" -o api.rest api-rest.go

```

```go

$ ./api.rest

```

### Subindo em Docker

```bash

$ docker build --no-cache -f Dockerfile -t jeffotoni/api.rest:latest .

```

```bash
$ docker run -p 8080:8080 --rm --name api.rest jeffotoni/api.rest 

```

### RPC

Primeiramente suba o server, logo depois execute o client

```go

$ go run main.go

```

```go

$ CGO_ENABLED=0 go build -ldflags="-s -w" -o rpc.server main.go

```

### gRPC

Primeiramente suba o server, logo depois execute o client.

```go

$ go run main.go

```

```go

$ CGO_ENABLED=0 go build -ldflags="-s -w" -o grpc.server main.go

```
### HTTP/3

Aqui está passo a passo como configurar o nginx com http/3.

[http/3 docker](https://hub.docker.com/r/ymuski/nginx-quic)
```bash
$ docker run --name nginx -d -p 80:80 -p 443:443/tcp -p 443:443/udp \
-v /etc/letsencrypt/:/opt/nginx/certs/ \
-v /opt/nginx/conf/example.nginx.conf:/etc/nginx/nginx.conf ymuski/nginx-quic
```

### Implementação quic-go

Aqui está uma implementação em Go sob o protocolo quic para que possa tirar proveito do http/3

[start guic-go](https://github.com/lucas-clemente/quic-go)