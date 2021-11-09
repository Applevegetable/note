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

## &集群初始化（只在master节点操作)

```
kubeadm init \
  --apiserver-advertise-address=116.63.185.25  \
  --image-repository registry.aliyuncs.com/google_containers \
  --kubernetes-version v1.18.0 \
  --service-cidr=10.1.0.0/16 \
  --pod-network-cidr=10.244.0.0/16
```





```
kubeadm init \
  --apiserver-advertise-address=192.168.0.175  \
  --image-repository registry.aliyuncs.com/google_containers \
  --kubernetes-version v1.18.0 \
  --service-cidr=10.1.0.0/16 \
  --pod-network-cidr=10.244.0.0/16
```

```bash
#坑位：使用公网IP地址在华为云上会出现bug,所以使用的是内网IP,插眼
```

```
--apiserver-advertise-addresss是master节点的IP地址
--image-repository是使用阿里镜像
--kubernetes-version指定拉取的镜像
```

###########################################

#### 出错时可以使用

```
kubeadm reset
```

###########################################

## 初始完毕获得token和密钥(保存)

```
kubeadm join 192.168.0.175:6443 --token wrdrup.mu2cay5w1g7no75p \
    --discovery-token-ca-cert-hash sha256:c869724300e708240d16b09a3b53295cab13a015ee2f931b162b766fc1254f08
```

### 备份配置（master节点）

```
mkdir -p $HOME/.kube
cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
#这里输入yes指令
chown $(id -u):$(id -g) $HOME/.kube/config
```



## 集群网络部署（flannel)

### 部署flannel(master节点)

```
kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
```

### 集群加入节点（node节点）

```
kubeadm join 192.168.0.175:6443 --token wrdrup.mu2cay5w1g7no75p \
    --discovery-token-ca-cert-hash sha256:c869724300e708240d16b09a3b53295cab13a015ee2f931b162b766fc1254f08
```



## 验证集群(master节点)

```
kubectl get nodes
```

可以看到其他的两个节点已经加入



## 集群测试（master节点）

在kubernetes集群中创建一个pod,验证是否正常工作

```
kubectl create deployment nginx --image=nginx
kubectl expose deployment nginx --port=80 --type=NodePort
kubectl get pod,srv
```

**访问地址就是节点的地址/port  ,**

```
NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)        AGE
service/kubernetes   ClusterIP   10.1.0.1     <none>        443/TCP        18h
service/nginx        NodePort    10.1.33.98   <none>        80:30002/TCP   12s

```



## 安装dashboard(master节点)

```
wget https://raw.githubusercontent.com/kubernetes/dashboard/v1.10.1/src/deploy/recommended/kubernetes-dashboard.yaml
```

### 换源

找到image位置，将k8s.grc.io换位lizhenliang

```
lizhenliang/kubernetes-dashboard-amd64:v1.10.1
```

运行

```
kubectl apply -f kubernetes-dashboard.yaml
```

运行

```bash
kubectl get pod -A -o wide |grep dash
kubectl get svc -A -o wide |grep dash
```

查看服务情况

```
kubectl -n kube-system describe pod
```

查看docker启动情况

```
docker ps
```

## 创建登陆用户

创建用户

```
kubectl create serviceaccount dashboard-admin -n kube-system
kubectl create clusterrolebinding dashboard-admin --clusterrole=cluster-admin --serviceaccount=kube-system:dashboard-admin
```

生成登陆token

```
kubectl describe secrets -n kube-system $(kubectl -n kube-system get secret | awk '/dashboard-admin/{print $1}')
```

```
token:      eyJhbGciOiJSUzI1NiIsImtpZCI6IjBWX2F6ODk3eXU0cU9JUmh0MVNDelk1eGxnbUE1V3NVRFB4aEZVbEVfeVEifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJkYXNoYm9hcmQtYWRtaW4tdG9rZW4tejd4anIiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoiZGFzaGJvYXJkLWFkbWluIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiZTQ2YzE5MGYtNmEwNy00OTdkLWIyMDktYTkyY2FkNzExYTM4Iiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50Omt1YmUtc3lzdGVtOmRhc2hib2FyZC1hZG1pbiJ9.cJQEwEiXoq4LbY4q70fZr7sM-_tGyoBuF7wt8UfNXHgnW6Nm2SvggLXj-MrmAjoBELKDNh8rJNDqTX4Eji-X9Fi_uYGeRIibaCCj93j2ZfDIjuI35JxGJxatlueuGnsXkikArHud-4fsO3gnUG_85E_Bxk_TXr4T0OeNaku08fTwv7NNbDsc--G5-EvmSu2Qs6CimeAThPzLKvP9CNIjAxAslTjJ3fWVQXIb9Nnx9TKaU_Juji7gZoWnG1ATfBG39rnIZNSL-6AK3Q_cA1J5yxZ13VdAalxUO8cVWFRxMJqWulbvoaB24V8yWN8wtH28blpL1XkFwUCiCvrrCQvKMA
```







