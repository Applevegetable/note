# Docker操作手册

1.拉取pha仓库中文件

2.将编译文件替换掉

- grid目录下的grid
- chain目录下的chain

3.构建镜像，在Dockerfile目录下

```bash
docker build -t chain:3.0 .
```

```
docker build -t grid:3.0 .
```

4.给镜像打标签

```
docker tag chain:3.0 huochain/chain:3.0
docker tag grid:3.0 huochain/grid:3.0
```

5.登陆docker

```
docker login
账号：huochain
密码：huochaindocker
```

6.上传镜像

```
docker push huochain/chain:3.0
docker push huochain/grid:3.0
```

7.修改docker-compose文件

修改chain和grid对应的镜像标签

8.测试

```
bash start.sh
```






