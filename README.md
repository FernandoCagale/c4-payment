# c4-payment [Flow](https://github.com/FernandoCagale/c4-kustomize)

### Docker

`running docker multi-stage builds and publish c4-payment`

```sh
$   ./scripts/publish.sh
```

### Kubernetes and Istio - [YAML](https://github.com/FernandoCagale/c4-kustomize/tree/master/c4-payment/base)

    *   deployment.yaml
    *   service.yaml
    *   virtualservice.yaml

# Running local

### Dependencies [docker-compose](https://github.com/FernandoCagale/c4-kustomize/blob/master/docker-compose.yml)

`Docker Mongodb`

```sh
$ docker run --network host --name mongo -d mongo
```

`Docker Rabbitmq`

```sh
$ docker run --network host --name rabbit -d rabbitmq
```

### Standard Go Project [Layout](https://github.com/golang-standards/project-layout)

```sh
$   go mod download
```

```sh
$   go mod vendor
```

`download "dependency injection"` [wire](https://github.com/google/wire)

```sh
$   go get -u github.com/google/wire/cmd/wire
```

```sh
$   ./scripts/start.sh
```