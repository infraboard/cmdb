# 订单管理

## 统一订单状态

### 各厂商状态

阿里云: [订单API文档](https://next.api.aliyun.com/api/BssOpenApi/2017-12-14/QueryOrders?params={}&lang=GO&tab=DOC)
+ New：新购。
+ Renew：续费。
+ Upgrade：升级。
+ Refund：退款。

腾讯云: [订单对象数据结构](https://cloud.tencent.com/document/api/555/19183#Deal)
+ modifyNetworkMode 调整带宽模式
+ modifyNetworkSize 调整带宽大小
+ refund 退款
+ downgrade 降配
+ upgrade 升配
+ renew 续费
+ purchase 购买
+ preMoveOut 包年包月迁出资源
+ preMoveIn 包年包月迁入资源
+ preToPost 预付费转后付费
+ postMoveOut 按量计费迁出资源
+ postMoveIn 按量计费迁入资源

华为云: [订单对象order_info](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=BSS&api=ShowCustomerOrderDetails)
+ 1:开通
+ 2:续订
+ 3:变更
+ 4:退订
+ 11:按需转包年/包月
+ 13:试用
+ 14:转商用
+ 15:费用调整

### 统一状态

+ purchase 购买
+ renew 续费
+ upgrade 升配
+ downgrade 降配
+ refund 退款
+ modify 费用调整