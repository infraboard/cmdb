# 操作审计

[事件查询控制台](https://actiontrail.console.aliyun.com/cn-hangzhou/event-list)

## ECS常用事件

+ Release: 实例释放
+ DeleteInstances (调用DeleteInstances释放一台或多台按量付费ECS实例或者到期的包年包月ECS实例)
+ RunInstances (调用RunInstances创建一台或多台按量付费或者包年包月ECS实例)
+ ModifyInstanceAttribute: 修改实例属性


## RDS常用事件

+ ModifyDBInstanceSpec 调用ModifyDBInstanceSpec接口变更RDS实例的（包括常规实例和只读实例，不包括灾备...)
+ DeleteDBInstance (调用DeleteDBInstance接口释放RDS实例)
+ CreateDBInstance (调用CreateDBInstance接口创建RDS实例)

## Redis常用事件

资源类型: ACS::Redis::DBInstance

+ ModifyInstanceSpec
+ Create
+ RemainRefund