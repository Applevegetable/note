# 安装Kubernetes





### 为了防止开机自动挂载swap分区，可以注释/etc/fstab中的条目：

```
sudo sed -i '/ swap / s/^\(.*\)$/#\1/g' /etc/fstab
```

查看交换分区状态

```
free -mh
```







## Kubernetes学习

### 创建一个Kubernetes集群

**Kubernetes 以更高效的方式跨集群自动分发和调度应用容器**

一个 Kubernetes 集群包含两种类型的资源:

- **Master** 调度整个集群
- **Nodes** 负责运行应用

**Master 负责管理整个集群。** Master 协调集群中的所有活动，例如调度应用、维护应用的所需状态、应用扩容以及推出新的更新。

**Node 是一个虚拟机或者物理机，它在 Kubernetes 集群中充当工作机器的角色** 每个Node都有 Kubelet , 它管理 Node 而且是 Node 与 Master 通信的代理。 Node 还应该具有用于处理容器操作的工具，例如 Docker 或 rkt 。处理生产级流量的 Kubernetes 集群至少应具有三个 Node 。

 **Node 使用 Master 暴露的 Kubernetes API 与 Master 通信。**

终端用户也可以使用 Kubernetes API 与集群交互。

Kubernetes 既可以部署在物理机上也可以部署在虚拟机上。

#### 使用minikube创建环境

查看版本

```shell
minikube version
```

启动集群

```
minikube start
```

#### kubectl：命令行界面

查看版本

```
kubectl version
```

会显示两个版本，一个是客户端版本，一个是服务器端版本，

客户端版本是kubectl的版本，服务器端版本是Kubernetes的版本

#### 查看集群的细节

```
kubectl cluster-info
```

可以看到部署的IP地址等信息

```bash
Kubernetes control plane is running at https://172.17.0.38:8443
KubeDNS is running at https://172.17.0.38:8443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy
```

#### 查看node状态

```
kubectl get nodes
```

可以看到node的状态

### 部署应用程序

一旦运行了 Kubernetes 集群，就可以在其上部署容器化应用程序。

为此，需要创建 Kubernetes **Deployment** 配置。Deployment 指挥 Kubernetes 如何创建和更新应用程序的实例。创建 Deployment 后，Kubernetes master 将应用程序实例调度到集群中的各个节点上。

创建应用程序实例后，Kubernetes Deployment 控制器会持续监视这些实例。 如果托管实例的节点关闭或被删除，则 Deployment 控制器会将该实例替换为群集中另一个节点上的实例。 **这提供了一种自我修复机制来解决机器故障维护问题。**

在没有 Kubernetes 这种编排系统之前，安装脚本通常用于启动应用程序，但它们不允许从机器故障中恢复。通过创建应用程序实例并使它们在节点之间运行， Kubernetes Deployments 提供了一种与众不同的应用程序管理方法。

![img](https://d33wubrfki0l68.cloudfront.net/8700a7f5f0008913aa6c25a1b26c08461e4947c7/cfc2c/docs/tutorials/kubernetes-basics/public/images/module_02_first_app.svg)

可以使用 Kubernetes 命令行界面 **Kubectl** 创建和管理 Deployment。Kubectl 使用 Kubernetes API 与集群进行交互。

创建 Deployment 时，您需要指定应用程序的容器映像以及要运行的副本数。

```
kubectl action resource
```

查看帮助

```
kubectl get nodes --help
```

**部署应用**,需要提供部署名称以及app镜像位置，及dockerhub的镜像名称

```
kubectl create deployment kubernetes-bootcamp --image=gcr.io/google-samples/kubernetes-bootcamp:v1
```

**列举出部署**

```
kubectl get deployments
```

**创建代理**,创建一个代理，将通信转发到集群范围的专用网络。在主机和集群之间有了一个连接

```
kubectl proxy
```

### 查看Pod和工作节点

Kubernetes 添加了一个 **Pod** 来托管你的应用实例

Pod 是 Kubernetes 抽象出来的，表示一组一个或多个应用程序容器（如 Docker），以及这些容器的一些共享资源。这些资源包括:

- 共享存储，当作卷
- 网络，作为唯一的集群 IP 地址
- 有关每个容器如何运行的信息，例如容器映像版本或要使用的特定端口。

Pod 为特定于应用程序的“逻辑主机”建模，并且可以包含相对紧耦合的不同应用容器。例如，Pod 可能既包含带有 Node.js 应用的容器，也包含另一个不同的容器，用于提供 Node.js 网络服务器要发布的数据。Pod 中的容器共享 IP 地址和端口，始终位于同一位置并且共同调度，并在同一工作节点上的共享上下文中运行。

Pod是 Kubernetes 平台上的原子单元。 当我们在 Kubernetes 上创建 Deployment 时，该 Deployment 会在其中创建包含容器的 Pod （而不是直接创建容器）。每个 Pod 都与调度它的工作节点绑定，并保持在那里直到终止（根据重启策略）或删除。 如果工作节点发生故障，则会在群集中的其他可用工作节点上调度相同的 Pod。

![img](https://d33wubrfki0l68.cloudfront.net/fe03f68d8ede9815184852ca2a4fd30325e5d15a/98064/docs/tutorials/kubernetes-basics/public/images/module_03_pods.svg)

一个 pod 总是运行在 **工作节点**node。工作节点是 Kubernetes 中的参与计算的机器，可以是虚拟机或物理计算机，具体取决于集群。每个工作节点由主节点管理。工作节点可以有多个 pod ，Kubernetes 主节点master会自动处理在群集中的工作节点上调度 pod 。 主节点的自动调度考量了每个工作节点上的可用资源。

每个 Kubernetes 工作节点至少运行:

- Kubelet，负责 Kubernetes 主节点和工作节点之间通信的过程; 它管理 Pod 和机器上运行的容器。

- 容器运行时（如 Docker）负责从仓库中提取容器镜像，解压缩容器以及运行应用程序。![img](https://d33wubrfki0l68.cloudfront.net/5cb72d407cbe2755e581b6de757e0d81760d5b86/a9df9/docs/tutorials/kubernetes-basics/public/images/module_03_nodes.svg)

  **常见的操作：**

  - **kubectl get** - 列出资源
  - **kubectl describe** - 显示有关资源的详细信息
  - **kubectl logs** - 打印 pod 和其中容器的日志
  - **kubectl exec** - 在 pod 中的容器上执行命令

查看有哪些pod

```
kubectl get pods
```

查看pod中有哪些容器以及使用什么镜像构建的这些容器

```
kubectl describe pods
```

Pod创建并跑起来之后，就可以直接执行命令

```
kubectl exec $POD_NAME -- env
```

与pod内的容器进行交互

```bash
kubectl exec -it $POD_NAME -- bash
```

检查容器内的app已经启动，使用

```
curl localhost:8080
```

### 使用Service暴露应用

Kubernetes [Pod](https://kubernetes.io/zh/docs/concepts/workloads/pods/) 是转瞬即逝的。 Pod 实际上拥有 [生命周期](https://kubernetes.io/zh/docs/concepts/workloads/pods/pod-lifecycle/)。 当一个工作 Node 挂掉后, 在 Node 上运行的 Pod 也会消亡。

[ReplicaSet](https://kubernetes.io/zh/docs/concepts/workloads/controllers/replicaset/) 会自动地通过创建新的 Pod 驱动集群回到目标状态，以保证应用程序正常运行。 换一个例子，考虑一个具有3个副本数的用作图像处理的后端程序。这些副本是可替换的; 前端系统不应该关心后端副本，即使 Pod 丢失或重新创建。也就是说，Kubernetes 集群中的每个 Pod (即使是在同一个 Node 上的 Pod )都有一个唯一的 IP 地址，因此需要一种方法自动协调 Pod 之间的变更，以便应用程序保持运行。

Kubernetes 中的服务(Service)是一种抽象概念，它定义了 Pod 的逻辑集和访问 Pod 的协议。Service 使从属 Pod 之间的松耦合成为可能。 和其他 Kubernetes 对象一样, Service 用 YAML [(更推荐)](https://kubernetes.io/zh/docs/concepts/configuration/overview/#general-configuration-tips) 或者 JSON 来定义. Service 下的一组 Pod 通常由 *LabelSelector* (请参阅下面的说明为什么您可能想要一个 spec 中不包含`selector`的服务)来标记。

尽管每个 Pod 都有一个唯一的 IP 地址，但是如果没有 Service ，这些 IP 不会暴露在集群外部。Service 允许您的应用程序接收流量。Service 也可以用在 ServiceSpec 标记`type`的方式暴露

- *ClusterIP* (默认) - 在集群的内部 IP 上公开 Service 。这种类型使得 Service 只能从集群内访问。
- *NodePort* - 使用 NAT 在集群中每个选定 Node 的相同端口上公开 Service 。使用`<NodeIP>:<NodePort>` 从集群外部访问 Service。是 ClusterIP 的超集。
- *LoadBalancer* - 在当前云中创建一个外部负载均衡器(如果支持的话)，并为 Service 分配一个固定的外部IP。是 NodePort 的超集。
- *ExternalName* - 通过返回带有该名称的 CNAME 记录，使用任意名称(由 spec 中的`externalName`指定)公开 Service。不使用代理。这种类型需要`kube-dns`的v1.7或更高版本。

另外，需要注意的是有一些 Service 的用例没有在 spec 中定义`selector`。 一个没有`selector`创建的 Service 也不会创建相应的端点对象。这允许用户手动将服务映射到特定的端点。没有 selector 的另一种可能是您严格使用`type: ExternalName`来标记。

![img](https://d33wubrfki0l68.cloudfront.net/cc38b0f3c0fd94e66495e3a4198f2096cdecd3d5/ace10/docs/tutorials/kubernetes-basics/public/images/module_04_services.svg)

Service 通过一组 Pod 路由通信。Service 是一种抽象，它允许 Pod 死亡并在 Kubernetes 中复制，而不会影响应用程序。在依赖的 Pod (如应用程序中的前端和后端组件)之间进行发现和路由是由Kubernetes Service 处理的。

Service 匹配一组 Pod 是使用 [标签(Label)和选择器(Selector)](https://kubernetes.io/zh/docs/concepts/overview/working-with-objects/labels), 它们是允许对 Kubernetes 中的对象进行逻辑操作的一种分组原语。标签(Label)是附加在对象上的键/值对，可以以多种方式使用:

- 指定用于开发，测试和生产的对象
- 嵌入版本标签
- 使用 Label 将对象进行分类

从集群中列举出现存的服务

```bash
kubectl get services
```



### 应用可扩展

当流量增加时，我们需要扩容应用程序满足用户需求

**扩缩** 是通过改变 Deployment 中的副本数量来实现的。

![img](https://d33wubrfki0l68.cloudfront.net/30f75140a581110443397192d70a4cdb37df7bfc/b5f56/docs/tutorials/kubernetes-basics/public/images/module_05_scaling2.svg)

扩展 Deployment 将创建新的 Pods，并将资源调度请求分配到有可用资源的节点上，收缩 会将 Pods 数量减少至所需的状态。

运行应用程序的多个实例需要在它们之间分配流量。服务 (Service)有一种负载均衡器类型，可以将网络流量均衡分配到外部可访问的 Pods 上。服务将会一直通过端点来监视 Pods 的运行，保证流量只分配到可用的 Pods 上。

查看ReplicaSet的状态

```bash
kubectl get rs
```

伸缩ReplicaSet的数量

```bash
kubectl scale deployments/kubernetes-bootcamp --replicas=4
```

### 滚动更新

用户希望应用程序始终可用，而开发人员则需要每天多次部署它们的新版本。在 Kubernetes 中，这些是通过滚动更新（Rolling Updates）完成的。 **滚动更新** 允许通过使用新的实例逐步更新 Pod 实例，零停机进行 Deployment 更新。新的 Pod 将在具有可用资源的节点上进行调度。

在前面的模块中，我们将应用程序扩展为运行多个实例。这是在不影响应用程序可用性的情况下执行更新的要求。默认情况下，更新期间不可用的 pod 的最大值和可以创建的新 pod 数都是 1。这两个选项都可以配置为（pod）数字或百分比。 在 Kubernetes 中，更新是经过版本控制的，任何 Deployment 更新都可以恢复到以前的（稳定）版本。

与应用程序扩展类似，如果公开了 Deployment，服务将在更新期间仅对可用的 pod 进行负载均衡。可用 Pod 是应用程序用户可用的实例。

滚动更新允许以下操作：

- 将应用程序从一个环境提升到另一个环境（通过容器镜像更新）
- 回滚到以前的版本
- 持续集成和持续交付应用程序，无需停机

那么容易，听我一一道来：

- image—Springboot项目一般是以jar包的形式跑在像centos等服务器上，运行nohup java -jar xxx.jar &命令就能启动起来。但是在k8s中，运行起来的的并不是jar,而是image，因此我们需要把jar打包成image;
- 自动扩缩—最基础的image有了，接下来就要考虑的是自动扩缩：顾名思义，比如说就是在服务访问量大的时候，我可以添加实例，来减少每个服务实例的压力，在访问量小的时候，我可以删除一部分实例，来实现资源的高效利用。
- 负载均衡—当我们的实例越来越多，我并不希望把所有的请求都落在一个实例上，如若不然，我们自动扩缩也就没了意义，传统方法我们可以用Nginx等实现负载均衡，待会来看看K8S能做些什么
- 域名绑定—这个就没什么好说的了。
  

K8S拉取镜像的时候是需要集群中配置CA证书

deployment:管理pod集群

service:管理pod中的服务

我们传统的集群负载均衡做法是在一台机器上安装Nginx,把我们的服务配置在Nginx上，外部直接访问你Nginx,它帮我们做负载，做限流





------

# 使用虚拟机搭建K8S集群（三台ubuntu)

### 虚拟机配置

|    虚拟机IP     |  虚拟机名  |
| :-------------: | :--------: |
| 192.168.116.136 | k8s-master |
| 192.168.116.137 | k8s-node1  |
| 192.168.116.138 | k8s-node2  |

### 修改各操作系统的host文件(每一台都必须配置)

```shell
vim /etc/hosts
```

将虚拟机IP地址和名称输进去,主机名可以随便起，这只是一个别名

```shell
192.168.116.136 k8s-master
192.168.116.137 k8s-node1
192.168.116.138 k8s-node2
```

检验效果

```
ping k8s-master
ping k8s-node1
ping k8s-node2
```

### 安装docker（略）

docker:20.10.8

查看docker运行状态

```shell
service docker status
```

设置开机启动docker

```shell
systemctl enable docker && systemctl restart docker && service docker status
```

### 关闭防火墙

```shell
systemctl stop firewalld
systemctl disable firewalld
```

### 关闭swap

```
swapoff -a
```

检查

```
free -h
```

### 配置iptable管理ipv4/6请求(配置内核转发)

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



### 卸载kebeadm组件

首先清空k8s集群设置

```
kubeadm reset
```

卸载管理组件

```
apt remove -y  kubelet kubectl kubeadm kubernetes-cni
```



### 安装kebeadm套件(ubuntu)1.18.0版本

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

### 集群初始化（只在master节点)

```
kubeadm init --apiserver-advertise-address=192.168.116.136\
			 --image-repository registry.aliyuncs.com/google_containers \
     		 --kubernetes-version v1.18.0 \
			 --service-cidr=10.1.0.0/16\
			 --pod-network-cidr=10.244.0.0/16
```

pod-network-cdir

指明 pod 网络可以使用的 IP 地址段。如果设置了这个参数，控制平面将会为每一个节点自动分配 CIDRs。

--service-cidr string   默认值："10.96.0.0/12"

为服务的虚拟 IP 地址另外指定 IP 地址段

**出现报错可以使用清空指令**

```
kubeadm reset
```

### 生成token和密钥

```bash
kubeadm join 192.168.116.136:6443 --token iltknh.8cyytyozugm3toy7 \
    --discovery-token-ca-cert-hash sha256:abef1c40a8fca877337eecfce88bff451448b147cd471e321d7b534ba1892495 
```

切换回普通用户，执行

```
  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

### 集群网络部署

有多种选择，

Calio,ACI,Flannel,Knitter,Multus

这里选择flannel

#### 部署flannel(master)

```
kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
```

### 集群加入节点

首先修改主机名称（node节点操作）

```
hostnamectl set-hostname k8s-node1
```

在其他两个Node上，以root用户权限执行

```
kubeadm join 192.168.116.136:6443 --token iltknh.8cyytyozugm3toy7 \
    --discovery-token-ca-cert-hash sha256:abef1c40a8fca877337eecfce88bff451448b147cd471e321d7b534ba1892495 

```

## 验证集群

在master上进行操作，之前证书存放的位置是普通用户，如果使用root用户会出现x509错误

```bash
kubectl get nodes
```

```
NAME        STATUS   ROLES    AGE     VERSION
k8s-node1   Ready    <none>   5m27s   v1.18.0
k8s-node2   Ready    <none>   2m40s   v1.18.0
ubuntu18    Ready    master   46m     v1.18.0
```

### 安装dashboard(在master节点上)

```bash
wget https://raw.githubusercontent.com/kubernetes/dashboard/v1.10.1/src/deploy/recommended/kubernetes-dashboard.yaml
```

修改image的值为，换源

```
lizhenliang/kubernetes-dashboard-amd64:v1.10.1
```

接下来运行

```bash
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

### 创建登陆用户

创建用户

```
kubectl create serviceaccount dashboard-admin -n kube-system
kubectl create clusterrolebinding dashboard-admin --clusterrole=cluster-admin --serviceaccount=kube-system:dashboard-admin
```

生成登陆token

```bash
kubectl describe secrets -n kube-system $(kubectl -n kube-system get secret | awk '/dashboard-admin/{print $1}')
```



```
token:eyJhbGciOiJSUzI1NiIsImtpZCI6InRSbl9ZQk52T2VFWkRZRVI5TkhONVZXZUZRNzBVU1RaNGxUNVBLcjJFZkEifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJkYXNoYm9hcmQtYWRtaW4tdG9rZW4tOHpxYjQiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoiZGFzaGJvYXJkLWFkbWluIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiYTJjZGNhMWMtYTA2Yy00ODQ4LWJjM2YtNmJlOTA0NGZlM2RmIiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50Omt1YmUtc3lzdGVtOmRhc2hib2FyZC1hZG1pbiJ9.i7dNrXOOFahIhB1ox0kNhz96bzcxT_FYfBrKzyXv6QnCrPJPX6Ev9LQJIKiUVjCrX5vRclRfIEI7JuWeQnE9vTR5Sreh-oUHNFNt0upJp5lCvccb9zd7nPtq58tADXzZk7ApN4ktr_7228lacJvmlnQvnJoSbvbjEpwPlTrN1JYvFJv4ojUBP__LXeu2HM4hWKfrzAtGRexAohzWey2N2bAkkbA0qmEUOUpTHEhk2dp8sI73dX_LPA0PrTS7xNDjcPwtYjfkgaJrsFeVRujfZBB1tHTPDjTieA5J0N45zSbML0o6S3LRHzw7Xe0ZReSlpOOApOXv6EWl05Rd18QKsw
```

开启代理

```
kubectl proxy
```

