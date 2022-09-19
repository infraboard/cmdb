# 账号管理

配置

// 存储
Rds       instance  db    username   password
Redis     instance  db    username   password
Mongodb   instance  db    username   password
                    ref@xxx 

// 使用应用凭证访问
Bucket              bucket

// 功能
Sms                 模版Id    name:value
Vms                 模版Id    value

注入的应用凭证:  app_id, app_secret

cluster

A      rds    ins_ref    ref_db       ref_account
       redis  ins_ref    ref_db
       mongo  inf_ref    ref_db       ref_account
       bucket bucket_ref 

关联映射:
RDS       RDS_0_ADDRESS   RDS_0_DB    RDS_0_USERNAME   RDS_0_PASSWORD      bind    RDS  instance_a, db01, acc_01
REDIS     REDIS_0_ADDRESS RDS_0_DB    RDS_0_PASSWORD
MONGO     MONGO_0_ADDRESS MONGO_0_DB  MONGO_0_USERNAME MONGO_0_PASSWORD
