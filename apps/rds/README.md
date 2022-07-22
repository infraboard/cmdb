# Rds管理


## 统一状态

### 各厂商状态

阿里云: [实例状态表](https://help.aliyun.com/document_detail/26315.html?spm=api-workbench.API%20Explorer.0.0.84161e0fyL4SoF)
+ Creating	创建中
+ Running	使用中
+ Deleting	删除中
+ Rebooting	重启中
+ DBInstanceClassChanging	升降级中
+ TRANSING	迁移中
+ EngineVersionUpgrading	迁移版本中
+ TransingToOthers	迁移数据到其他RDS中
+ GuardDBInstanceCreating	生产灾备实例中
+ Restoring	备份恢复中
+ Importing	数据导入中
+ ImportingFromOthers	从其他RDS实例导入数据中
+ DBInstanceNetTypeChanging	内外网切换中
+ GuardSwitching	容灾切换中
+ INS_CLONING	实例克隆中
+ Released	已释放实例


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
+ 3 - 开放Slave中
+ 4 - 外网访问开通中
+ 5 - 批量操作执行中
+ 6 - 回档中
+ 7 - 外网访问关闭中
+ 8 - 密码修改中
+ 9 - 实例名修改中
+ 10 - 重启中
+ 12 - 自建迁移中
+ 13 - 删除库表中
+ 14 - 灾备实例创建同步中
+ 15 - 升级待切换
+ 16 - 升级切换中
+ 17 - 升级切换完成
+ 19 - 参数设置待执行