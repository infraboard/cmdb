# 主机资源管理


## 统一状态

### 各厂商状态

阿里云: [API 文档](https://next.api.aliyun.com/api/Ecs/2014-05-26/DescribeInstances?params={})
+ Pending：创建中。
+ Running：运行中。
+ Starting：启动中。
+ Stopping：停止中。
+ Stopped：已停止。

腾讯云: [实例状态表](https://cloud.tencent.com/document/api/213/15753#InstanceStatus)
+ PENDING：表示创建中
+ LAUNCH_FAILED：表示创建失败
+ RUNNING：表示运行中
+ STOPPED：表示关机
+ STARTING：表示开机中
+ STOPPING：表示关机中
+ REBOOTING：表示重启中
+ SHUTDOWN：表示停止待销毁
+ DELETING：表示销毁中。

华为云: [云服务器状态](https://support.huaweicloud.com/api-ecs/ecs_08_0002.html)

+ BUILD: 创建实例后，在实例状态进入运行中之前的状态。
+ REBOOT: 实例正在进行重启操作。
+ HARD_REBOOT: 实例正在进行强制重启操作
+ REBUILD: 实例正在重建中。
+ MIGRATING: 实例正在热迁移中。
+ RESIZE: 实例接收变更请求，开始进行变更操作。
+ ACTIVE: 实例正常运行状态。
+ SHUTOFF: 实例被正常停止。
+ REVERT_RESIZE: 实例正在回退变更规格的配置。
+ VERIFY_RESIZE: 实例正在校验变更完成后的配置。
+ ERROR: 实例处于异常状态。
+ DELETED: 实例已被正常删除。

AWS: [InstanceState](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_InstanceState.html)
+ 0 : pending
+ 16 : running
+ 32 : shutting-down
+ 48 : terminated
+ 64 : stopping
+ 80 : stopped

Vmware:
+ poweredOff: 开机
+ poweredOn: 关机
+ suspended: 暂停

### 统一状态

则中选中腾讯云作为统一状态, 其他状态映射到改状态
+ PENDING：表示创建中
+ LAUNCH_FAILED：表示创建失败
+ RUNNING：表示运行中
+ ERROR: 运行异常
+ STOPPED：表示关机
+ STARTING：表示开机中
+ STOPPING：表示关机中
+ REBOOTING：表示重启中
+ SHUTDOWN：表示停止待销毁
+ DELETING：表示销毁中。
+ DESTROYED: 已销毁, 
