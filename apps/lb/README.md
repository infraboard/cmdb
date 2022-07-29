# 负载均衡管理


## 统一状态


### 各云商状态

阿里云: [API 文档](https://next.api.aliyun.com/api/Slb/2014-05-15/DescribeLoadBalancers?params={}&lang=GO&tab=DOC)
+ inactive: 实例已停止，此状态的实例监听不会再转发流量。
+ active: 实例运行中，实例创建后，默认状态为active。
+ locked: 实例已锁定。当负载均衡实例到期后，但到期时间未超过7天时，负载均衡实例进入锁定状态。此种状态下，您不能对负载均衡实例进行任何操作，并且实例不再会进行流量转发，但会保留实例的IP和其它配置。

腾讯云: [LoadBalancer数据结构](https://cloud.tencent.com/document/api/214/30694#LoadBalancer)
0：创建中
1：正常运行。

华为云: [API 文档](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=ELB&api=ListLoadbalancers&version=v2)
+ ONLINE
+ FROZEN

### 统一后状态

+ PENDING：表示创建中
+ RUNNING：表示运行中
+ LOCKED: 实例已锁定