# Dockerfile文件
```Dockerfile
FROM golang:1.18 AS builder
WORKDIR /build/
copy . .
ENV GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=darwin \
    GOARCH=amd64 \
    CGO_ENABLED=0 \
    GOSUMDB=sum.golang.google.cn

RUN go build -o myhttpserver *.go


FROM alpine
WORKDIR /mygo/
COPY --from=builder /build/myhttpserver .
EXPOSE 8080
CMD ["./myhttpserver"]

```


# 编译过程
编译命令：
```shell
docker build . -t http-server:v0.0.1 
```
编译输出日志：
```
[+] Building 112.1s (13/13) FINISHED                                                                                                                    
 => [internal] load build definition from Dockerfile                                                                                               0.0s
 => => transferring dockerfile: 388B                                                                                                               0.0s
 => [internal] load .dockerignore                                                                                                                  0.0s
 => => transferring context: 2B                                                                                                                    0.0s
 => [internal] load metadata for docker.io/library/alpine:latest                                                                                   0.0s
 => [internal] load metadata for docker.io/library/golang:1.18                                                                                     0.0s
 => [builder 1/4] FROM docker.io/library/golang:1.18                                                                                               0.0s
 => [stage-1 1/3] FROM docker.io/library/alpine                                                                                                    0.0s
 => [internal] load build context                                                                                                                  0.0s
 => => transferring context: 553B                                                                                                                  0.0s
 => CACHED [builder 2/4] WORKDIR /build/                                                                                                           0.0s
 => [builder 3/4] COPY . .                                                                                                                         0.1s
 => [builder 4/4] RUN go build -o myhttpserver *.go                                                                                              111.5s
 => CACHED [stage-1 2/3] WORKDIR /mygo/                                                                                                            0.0s
 => CACHED [stage-1 3/3] COPY --from=builder /build/myhttpserver .                                                                                 0.0s
 => exporting to image                                                                                                                             0.0s
 => => exporting layers                                                                                                                            0.0s
 => => writing image sha256:1d2987689a8679a88f55ee4f9926b02859b359d99eabade09a497aa671f039cf                                                       0.0s
 => => naming to docker.io/library/httpserver:0.0.1                
```

# 运行容器
命令：
```shell
 docker run -d http-server:v0.0.1  
```
查看
```shell
benny•~» docker ps                                                                                                               [10:35:06]
CONTAINER ID   IMAGE                COMMAND            CREATED         STATUS         PORTS      NAMES
94faaf30248b   http-server:v0.0.1   "./myhttpserver"   4 seconds ago   Up 4 seconds   8080/tcp   laughing_bell
```

# 推送镜像到dockerhub
登录
```shell
benny•goCode/mycncamp/module02(main⚡)» docker login                                                                                         [10:09:48]
Login with your Docker ID to push and pull images from Docker Hub. If you don't have a Docker ID, head over to https://hub.docker.com to create one.
Username: zhoucheng45
Password: 
Login Succeeded

Logging in with your password grants your terminal complete access to your account. 
For better security, log in with a limited-privilege personal access token. Learn more at https://docs.docker.com/go/access-tokens/

```
推送镜像
```shell
benny•goCode/mycncamp/module02(main⚡)» docker tag http-server:v0.0.1 zhoucheng45/test:http-servicev0.0.1                                    [10:35:10]
benny•goCode/mycncamp/module02(main⚡)» docker push zhoucheng45/test:http-servicev0.0.1                                                      [10:36:00]
The push refers to repository [docker.io/zhoucheng45/test]
c5458d73bcd3: Pushed 
8ba731c5e2bb: Pushed 
994393dc58e7: Mounted from library/alpine 
http-servicev0.0.1: digest: sha256:71fe201946875dff526516bfb85cd40c05cdc36402eb73415b70ae04320e29df size: 946
```
镜像地址： https://hub.docker.com/repository/docker/zhoucheng45/test
# 查看容器网络
mac上不知道怎么安装nsenter工具。借用debian容器以特权启动，共享主机的namespace。
```shell
benny•~» docker run -it --privileged --pid=host debian nsenter -t 1 -m -u -n -i sh                                               [10:59:06]
/ # nsenter  -t 6884 -n ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN group default qlen 1000
    link/ipip 0.0.0.0 brd 0.0.0.0
3: ip6tnl0@NONE: <NOARP> mtu 1452 qdisc noop state DOWN group default qlen 1000
    link/tunnel6 :: brd ::
32: eth0@if33: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
/ #
```
得到容器ip： 172.17.0.2/16
