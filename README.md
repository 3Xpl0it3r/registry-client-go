# registry-client-go
a client api for docker registry



## 安装/获取
> 安装
```bash
#git clone https://github.com/3Xpl0it3r/registry-client-go.git
# cd registry-client-go/cli
# go build .
```

## 命令行使用
>命令行使用方式
```bash
./cli [-u http://registry_url:port|https://registry_url:port] [-k username:password] [-ca /path/ca_pem] [--insecure] [commnd] [args]
```
>参数详解
>> -u 指定registry仓库的地址
>> -k 基于basic认证，格式用户名：密码
>> -ca 如果采用https认证，并且没有设置--insecure这个参数，则需要设置ca证书来验证服务器证书的合法性
>> --insecure 忽略服务器证书验证
>查看镜像
```bash
./cli -u http://192.168.1.10:5000 image list
```
>查看镜像tag
```bash
./cli -u http://192.168.1.10:5000 tag list 镜像名称
```
> 删除镜像
```bash
./cli -u http://192.168.1.10:5000 image rm 镜像名称:镜像tag
```

## sdk使用
```go
import "l0calh0st.cn/docker-registry/registry"
```

### 创建一个客户端
```go
// insecure ,
clientInsecure := registry.NewInsecureRegistryApi("http://10.23.6.90:5000")
// secure client
clientTls := registry.NewTlsRegistryApi("10.23.6.90:5000", "/root/cert/ca.pem")
// if need some basic,
client.SetBasicAuth("username", "password")
```


### 查看所有的仓库
```go
repo,err := client.RepositoriesList()
```

### 列举所有tag的名称
```go
tags,err := client.TagsList("image_name")
```

### 获取一个具体镜像的digest
```go
digest,err := client.Digest("imagename", "reference")
// arg
// imagename string
// reference can be a tag or a digest
```

### 删除一个镜像
```go
err := client.ImageDelete("image name", "reference")
// reference must be digest
```
