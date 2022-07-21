# 资产秘文管理

配置

// 存储
Rds       instance  db    username   password
Redis     instance  db    username   password
Mongodb   instance  db    username   password
                    ref@xxx 

Bucket              bucket client_id client_credential

// 功能
Sms                 client_id client_credential
Vms                 client_id client_credential



cluster

A      rds    ins_ref    ref_db       ref_account
       redis  ins_ref    ref_db
       mongo  inf_ref    ref_db       ref_account
       bucket bucket_ref 