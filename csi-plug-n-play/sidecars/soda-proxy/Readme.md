## soda-proxy

soda-proxy is a simple http server which helps to call the api's of SodaFoundation in a simpler way, it bypasses the authentication and directly helps to connect the osdsapi-server using the client provided by sodafoundation/api.

### Build
Please follow the below steps to build and run the proxy

```go
go get github.com/sodafoundation/nbp

cd $GOPATH/src/github.com/sodafoundation/nbp/csi-plug-n-play/sidecars/soda-proxy

go build cmd/proxy.go
```

Before running soda-proxy you need to export the below variables.(These are the same env variables which are required by osdsctl, for more reference you can see [this](https://docs.sodafoundation.io/soda-gettingstarted/installation-using-ansible/#how-to-test-soda-projects-cluster))

```go

export OPENSDS_ENDPOINT=http://{your_host_ip}:50040
export OPENSDS_AUTH_STRATEGY=keystone
export OS_AUTH_URL=http://{your_host_ip}/identity
export OS_USERNAME=admin
export OS_PASSWORD=opensds@123
export OS_TENANT_NAME=admin
export OS_PROJECT_NAME=admin
export OS_USER_DOMAIN_ID=default
```

```go
./proxy

```

By default soda-proxy runs on `0.0.0.0:50029`
