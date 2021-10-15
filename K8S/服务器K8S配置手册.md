# 服务器K8S配置手册



|    IP地址     |  主机名称  |   操作系统   |
| :-----------: | :--------: | :----------: |
| 116.36.185.25 | k8s-master | Ubuntu 18.04 |
| 139.9.247.94  | k8s-node1  | Ubuntu 18.04 |
| 122.9.144.124 | k8s-node2  | Ubuntu 18.04 |



绑定主机名称

```
hostnamectl --static set-hostname k8s-master 
hostnamectl --static set-hostname k8s-node1
hostnamectl --static set-hostname k8s-node2
```

## 修改各操作系统的host文件(每一台都必须配置)

```shell
vim /etc/hosts
```

将虚拟机IP地址和名称输进去,主机名可以随便起，这只是一个别名

```shell
116.63.185.25  k8s-master
139.9.247.94   k8s-node1
122.9.144.124  k8s-node2
```

检验效果（互相测试）

```
ping k8s-master
ping k8s-node1
ping k8s-node2
```



## 安装docker(生产环境使用18.6.3版本)

### 配置docker安装环境

```
apt-get install apt-transport-https ca-certificates curl gnupg-agent software-properties-common
```

### 添加阿里云的docker GPG密钥

```
curl -fsSL http://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg | sudo apt-key add -
```

### 添加阿里镜像源

```
sudo add-apt-repository "deb [arch=amd64] http://mirrors.aliyun.com/docker-ce/linux/ubuntu $(lsb_release -cs) stable"

#更新
sudo apt-get update
```

### 安装指定版本的docker-ce

```
apt-get install -y docker-ce=18.06.3~ce~3-0~ubuntu
```

### 添加阿里源

```

vim /etc/docker/daemon.json

{
   "registry-mirrors": ["https://alzgoonw.mirror.aliyuncs.com"]
}

```

### 重启docker

```
systemctl restart docker
```

### 查看docker版本

```
docker version
```



## 查看交换分区是否关闭

```
free -h
```

## 查看防火墙状态

```bash
ufw status
```

## 配置iptable管理ipv4/6请求(配置内核转发)

```
vim /etc/sysctl.d/k8s.conf
```

写入

```bash
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
```

使配置生效

```
sysctl --system
```

## 安装kebeadm套件(ubuntu)1.18.0版本

- `kubelet`: k8s 的核心服务
- `kubeadm`: 这个是用于快速安装 k8s 的一个集成工具，我们在`master1`和`worker1`上的 k8s 部署都将使用它来完成。
- `kubectl`: k8s 的命令行工具，部署完成之后后续的操作都要用它来执行

```shell
# 使得 apt 支持 ssl 传输
apt-get update && apt-get install -y apt-transport-https
# 下载 gpg 密钥
curl https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | apt-key add - 
# 添加 k8s 镜像源
cat <<EOF >/etc/apt/sources.list.d/kubernetes.list
deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main
EOF
# 更新源列表
apt-get update
# 下载 kubectl，kubeadm以及 kubelet
apt-get install kubelet=1.18.0-00 kubeadm=1.18.0-00 kubectl=1.18.0-00
#apt-mark 用于将软件包标记/取消标记为自动安装。 hold 选项用于将软件包标记为保留，以防止软件包被自动安装、升级或删除。
sudo apt-mark hold kubelet kubeadm kubectl
```

### 校验版本

```
kubelet --version
kubeadm version
kubectl version --client
```

## 启动kubectl服务

```bash
systemctl enable kubelet && systemctl start kubelet
```



## 集群初始化（只在master节点操作)

```
kubeadm init --apiserver-advertise-address=116.36.185.25 --image-repository registry.aliyuncs.com/google_containers  --kubernetes-version v1.18.0  --service-cidr=10.1.0.0/16 --pod-network-cidr=10.244.0.0/16
```

pod-network-cdir

指明 pod 网络可以使用的 IP 地址段。如果设置了这个参数，控制平面将会为每一个节点自动分配 CIDRs。

--service-cidr string   默认值："10.96.0.0/12"

为服务的虚拟 IP 地址另外指定 IP 地址段

都是默认地址

#### 出错时可以使用

```
kubeadm reset
```

初始化时会出现问题，主要是代理的原因

```bash
This error is likely caused by:
		- The kubelet is not running
		- The kubelet is unhealthy due to a misconfiguration of the node in some way (required cgroups disabled)
```



### 初始完毕获得token和密钥

```
systemctl status kubelet
journalctl -xeu kebelet
```





```
kubeadm init \
  --apiserver-advertise-address=116.63.185.25  \
  --image-repository registry.aliyuncs.com/google_containers \
  --kubernetes-version v1.18.0 \
  --service-cidr=10.1.0.0/16 \
  --pod-network-cidr=10.244.0.0/16

```

解决方案：

```
#加标签
#docker images查看镜像，其中pause镜像有问题
docker tag registry.aliyuncs.com/google_containers/pause:3.2 k8s.gcr.io/pause:3.2
```

```
images=(  
    kube-apiserver:v1.18.0
    kube-controller-manager:v1.18.0
    kube-scheduler:v1.18.0
    kube-proxy:v1.18.0
    pause:3.2
    etcd:3.4.3-0
    coredns:1.6.7
)

for imageName in ${images[@]} ; do
    docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/$imageName
    docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/$imageName k8s.gcr.io/$imageName
    docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/$imageName
done
```

