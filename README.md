# k8s_demo2

## go

```
$ go mod init k8s_demo2
```

# telepresence 

```
telepresence intercept demo2 --port 18081:8082
```

```
  - containerID: docker://55e01525535f73f311c7872e4af747d4a6a9ff77cbca2382e04d38a4c228b093
    image: datawire/tel2:2.4.0
    imageID: docker-pullable://datawire/tel2@sha256:f88e16070521acfaa592b048f7bb431f9ef14268f51b0b7fd83978c0ec800391
    lastState: {}
    name: traffic-agent
    ready: true
    restartCount: 0
    started: true
    state:
      running:
        startedAt: "2021-08-20T11:36:22Z"
```

```
telepresence intercept demo2 --port 18081:8082 --http-match=HTTP2_HEADER=test
```

```
  - containerID: docker://82c663fb56ed1fcf771bc14dfd80c0436ab74c5a916a3029efd80a3c533d3e95
    image: datawire/ambassador-telepresence-agent:1.10.0
    imageID: docker-pullable://datawire/ambassador-telepresence-agent@sha256:f6e7816d5774fd845de0746b7fd34bdb8c14edf85b67f227ee4b23c37f7bbbbf
    lastState: {}
    name: traffic-agent
    ready: true
    restartCount: 0
    started: true
    state:
      running:
        startedAt: "2021-08-20T11:39:32Z"
```

## name指定

```
telepresence intercept nginx --port 80:http
```

## OpenAPI

oapi-codegenによるRESTインターフェースの作成

### install

```
$ go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.0
```

### Go コード生成

- types

```
$ oapi-codegen -package domain -generate "types" -o src/domain/petstore.go openapi.yaml 
```

- spec

```
$ oapi-codegen -package domain -generate "spec" -o src/domain/spec.go openapi.yaml 
```

- server

```
$ oapi-codegen -package domain -generate "chi-server" -o src/domain/server.go openapi.yaml 
```