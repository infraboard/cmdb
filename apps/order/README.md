# 订单管理

## 统一订单分类

### 各厂商分类

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

### 统一分类

+ purchase 购买
+ renew 续费
+ upgrade 升配
+ downgrade 降配
+ refund 退款
+ modify 费用调整


## 统一订单状态


### 各厂商状态

阿里云: [订单API文档](https://next.api.aliyun.com/api/BssOpenApi/2017-12-14/QueryOrders?params={}&lang=GO&tab=DOC)
+ Unpaid：未支付。
+ Paid：已支付。
+ Cancelled：已作废。

腾讯云: [腾讯云API文档](https://console.cloud.tencent.com/api/explorer?Product=billing&Version=2018-07-09&Action=DescribeDealsByCond&SignVersion=)
+ 1：未支付 
+ 2：已支付
+ 3：发货中 
+ 4：已发货 
+ 5：发货失败 
+ 6：已退款 
+ 7：已关单 
+ 8：订单过期 
+ 9：订单已失效 
+ 10：产品已失效 
+ 11：代付拒绝 
+ 12：支付中

华为云: [API文档](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=BSS&api=ListCustomerOrders)
+ 1:待审核
+ 3:处理中
+ 4:已取消
+ 5:已完成
+ 6:待支付
+ 9:待确认

### 统一状态

+ 1：未支付 
+ 8：订单过期 
+ 2：支付中
+ 4: 已取消
+ 5：已支付
+ 6：已退款 
+ 7：已关单 



## 资源续费订单

// RegionId实例所属的地域ID
// ResourceId查询续费价格的资源ID
// Period指定续费时长
// PriceUnit指定续费周期
// ExpectedRenewDay统一到期日

// InstanceIds
// Period 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36, 48, 60。
/ RenewFlag  自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费<br><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费<br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费<br><br>默认取值：NOTIFY_AND_MANUAL_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。