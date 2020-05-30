# c4-payment [Flow](https://github.com/FernandoCagale/c4-kustomize)

## Dependencies

`Docker Mongodb`

```sh
$ docker run --network host --name mongo -d mongo
```

`Docker Rabbitmq`

```sh
$ docker run --network host --name rabbit -d rabbitmq
```

## Build Docker

`build and publish c4-payment`

```sh
$   ./scripts/publish.sh
```

## Kubernetes [YAML](https://github.com/FernandoCagale/c4-kustomize/tree/master/c4-payment/base)

    *   deployment.yaml
    *   service.yaml
    *   virtualservice.yaml

## Running local

```sh
$   go mod download
```

```sh
$   go mod vendor
```

`download wire "dependency injection"`

```sh
$   go get -u github.com/google/wire/cmd/wire
```

```sh
$   ./scripts/start.sh
```