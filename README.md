# registry-client-go
a client api for docker registry





## Usage
```go
import "l0calh0st.cn/docker-registry/registry"
```

## create client
```go
// insecure ,
clientInsecure := registry.NewInsecureRegistryApi("http://10.23.6.90:5000")
// secure client
clientTls := registry.NewTlsRegistryApi("10.23.6.90:5000", "/root/cert/ca.pem")
// if need some basic,
client.SetBasicAuth("username", "password")
```


## list all repo
```go
repo,err := client.RepositoriesList()
```

## list all tags for a given image
```go
tags,err := client.TagsList("image_name")
```

### get a image digest
```go
digest,err := client.Digest("imagename", "reference")
// arg
// imagename string
// reference can be a tag or a digest
```

### delete a image
```go
err := client.ImageDelete("image name", "reference")
// reference must be digest
```
