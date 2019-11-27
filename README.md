**c4-payment - docker**

`Docker Mongodb`

```sh
$ docker run --network host --name mongo -d mongo
```

`Docker Rabbitmq`

```sh
$ docker run --network host --name rabbit -d rabbitmq
```

`Docker build c4-payment`

```sh
$   docker build -t c4-payment .
```

`Docker c4-payment`

```sh
$   docker run -d --name c4-payment -p 8080:8080 c4-payment
```

**c4-payment - local**


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

`generate wire_gen.go`

```sh
$   wire
```

`generate build`

```sh
$   go build -o bin/application
```


```sh
$   ./bin/application
```