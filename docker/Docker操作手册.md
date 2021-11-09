# Docker操作手册

1.替换文件

2.修改Dockerfile文件

3.构建镜像

```bash
docker build -t chain:3.0 .
```

4.给镜像打标签

```
docker tag chain:3.0 huochain/chain:3.0
```

5.上传镜像

```
docker push huochain/chain:3.0
```



查看镜像

```
docker search huochain
```



修改docker-compose文件对应的镜像版本号

