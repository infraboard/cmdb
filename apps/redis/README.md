# Redis管理

## 统一状态


### 各厂商状态

阿里云: [API 文档](https://next.api.aliyun.com/api/R-kvstore/2015-01-01/DescribeInstances?params={}&lang=GO&tab=DOC)
+ Normal：正常。
+ Creating：创建中。
+ Error：创建失败。
+ Flushing：清除中。
+ Released：已释放。
+ BackupRecovering：备份恢复中。
+ MinorVersionUpgrading：小版本升级中。
+ MajorVersionUpgrading：大版本升级中，可正常访问。
+ NetworkModifying：网络变更中。
+ Inactive：被禁用。
+ Changing：修改中。
+ Transforming：转换中。
+ Migrating：迁移中。
+ Unavailable：服务停止。
+ SSLModifying：SSL变更中。

腾讯云: [API 文档](https://console.cloud.tencent.com/api/explorer?Product=redis&Version=2018-04-12&Action=DescribeInstances&SignVersion=)
+ 0-待初始化
+ 1-流程中
+ 2-运行中

+ -2-已隔离
+ -3-待删除

华为云: [缓存实例状态说明](https://support.huaweicloud.com/api-dcs/dcs-api-0312047.html)
+ CREATING: 申请缓存实例后，在缓存实例状态进入运行中之前的状态。
+ RUNNING: 缓存实例正常运行状态。在这个状态的实例可以运行您的业务。
+ CREATEFAILED: 缓存实例处于创建失败的状态。
+ ERROR: 缓存实例处于故障的状态。
+ RESTARTING: 缓存实例正在进行重启操作。
+ FROZEN: 缓存实例处于已冻结状态，用户可以在“我的订单”中续费开启冻结的缓存实例。
+ FLUSHING: 缓存实例数据清空中的状态。
+ RESTORING: 缓存实例数据恢复中的状态。
+ EXTENDING: 缓存实例处于正在扩容的状态。

### 统一状态

+ PENDING：表示创建中
+ RUNNING：表示运行中
+ ERROR: 缓存实例处于故障的状态。
+ REBOOTING：表示重启中
+ ISOLATIONING: 隔离中
+ ISOLATIONED: 已隔中
+ FLUSHING: 缓存实例数据清空中的状态。
+ RESTORING: 缓存实例数据恢复中的状态。
+ DESTROYED: 已销毁
+ RESTORING: 备份恢复中
+ UPGRADING: 迁移版本中
+ NET_CHANGING: 内外网切换中
+ EXTENDING: 缓存实例处于正在扩容的状态。
+ MODIFYING: 实例配置变更生效中
+ TRANSING: 迁移中