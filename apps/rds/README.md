# Rds管理


## 统一状态

### 各厂商状态

阿里云: [实例状态表](https://help.aliyun.com/document_detail/26315.html?spm=api-workbench.API%20Explorer.0.0.84161e0fyL4SoF)
+ Creating	创建中
+ Running	使用中
+ Deleting	删除中
+ Released	已释放实例
+ Rebooting	重启中
+ Restoring	备份恢复中
+ TRANSING	迁移中
+ DBInstanceClassChanging	升降级中
+ GuardSwitching	容灾切换中
+ GuardDBInstanceCreating	生产灾备实例中
+ Importing	数据导入中
+ INS_CLONING	实例克隆中
+ EngineVersionUpgrading	迁移版本中
+ DBInstanceNetTypeChanging	内外网切换中

+ TransingToOthers	迁移数据到其他RDS中
+ ImportingFromOthers	从其他RDS实例导入数据中


腾讯云: [DescribeDBInstancesAPI文档](https://console.cloud.tencent.com/api/explorer?Product=cdb&Version=2017-03-20&Action=DescribeDBInstances&SignVersion=)

实例状态，可取值：
+ 0 - 创建中
+ 1 - 运行中
+ 4 - 正在进行隔离操作
+ 5 - 隔离中（可在回收站恢复开机）

实例任务状态，可能取值：
+ 0 - 没有任务

+ 1 - 升级中
+ 2 - 数据导入中
+ 6 - 回档中
+ 10 - 重启中
+ 12 - 自建迁移中
+ 14 - 灾备实例创建同步中
+ 15 - 升级待切换
+ 16 - 升级切换中
+ 17 - 升级切换完成
+ 4 - 外网访问开通中
+ 7 - 外网访问关闭中

+ 3 - 开放Slave中
+ 5 - 批量操作执行中
+ 8 - 密码修改中
+ 9 - 实例名修改中
+ 13 - 删除库表中
+ 19 - 参数设置待执行

华为云: [API 文档](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=RDS&api=ListInstances)

+ 值为“BUILD”,表示实例正在创建。
+ 值为“ACTIVE”,表示实例正常。
+ 值为“FAILED”,表示实例异常。
+ 值为“FROZEN”,表示实例冻结。
+ 值为“REBOOTING”,表示实例正在重启。
+ 值为“MIGRATING”,表示实例正在迁移。
+ 值为“RESTORING”,表示实例正在恢复。
+ 值为“MODIFYING”,表示实例正在扩容。
+ 值为“SWITCHOVER”,表示实例正在主备切换。
+ 值为“STORAGE FULL”,表示实例磁盘空间满
+ 值为“BACKING UP”,表示实例正在进行备份。

+ 值为“MODIFYING INSTANCE TYPE”,表示实例正在转主备。
+ 值为“MODIFYING DATABASE PORT”,表示实例正在修改数据库端口。

### 统一状态

+ PENDING：表示创建中
+ RUNNING：表示运行中
+ ERROR: 运行异常
+ REBOOTING：表示重启中
+ DELETING：表示销毁中。
+ DESTROYED: 已销毁
+ ISOLATIONING: 隔离中
+ ISOLATIONED: 已隔中
+ RESTORING: 备份恢复中
+ TRANSING: 迁移中
+ SWITCHOVER: 表示实例正在主备切换
+ GUARD_CREATING: 灾备实例创建同步中
+ IMPORTING:  数据导入中
+ STORAGE_FULL: 表示实例磁盘空间满
+ CLONING: 实例克隆中
+ UPGRADING: 迁移版本中
+ BACKING_UP: 表示实例正在进行备份
+ NET_CHANGING: 内外网切换中
+ MODIFYING: 实例配置变更生效中