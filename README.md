# <img src=https://github.com/Gambya/zincsearchstash/blob/main/logo.png> Zinc Search Stash

Serviço para consumo de fila em rabbitmq para inserir dados no [ZincSearch](https://docs.zincsearch.com/).

# Run

### Docker

```sh
docker run -it -e ZINC_EXCHANGE_NAME=exchangename -e ZINC_ROUTING_KEY=routingkey -e ZINC_QUEUE=queuename -e ZINC_BROKER_URL=amqp://user:password@host:port/ -e ZINC_URL=http://host_zincsearch/api/%s/_doc -e ZINC_INDEX=index -e ZINC_USER=user -e ZINC_PASS=password -e ZINC_LOG_LEVEL=-1 gambya/zincsearchstash:0.0.1-alpine
```

### Docker Compose

```sh
git clone https://github.com/Gambya/zincsearchstash.git
cd zincsearchstash
docker-compose up -d
```

### Terminal

```sh
go mod tidy
export ZINC_EXCHANGE_NAME=exchangename
export ZINC_ROUTING_KEY=routingkey
export ZINC_QUEUE=queuename
export ZINC_BROKER_URL=amqp://user:password@host:port/
export ZINC_URL=http://host_zincsearch/api/%s/_doc
export ZINC_INDEX=index
export ZINC_USER=user
export ZINC_PASS=pass
export ZINC_LOG_LEVEL=-1
go run cmd/main.go
```
