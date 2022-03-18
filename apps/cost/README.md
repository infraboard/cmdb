基于实例聚合后的账单为准


类型         value   year  month  cost  delta_cost delta_percent task_id
domain(总账)        


namespace(项目)  domain
env
vendor 
account
  
  

总帐: 0xxxxx


1月   2月   3月   4月

namesapce|vendor|account|env|resource_type

---------
ns1  xxxx  +192(200%)   
ns2  xxxx  -110(100%)
ns3  xxxx  +100(100%)   

分摊处理

项目账单: 0xxxxxxx


1月   2月   3月   4月


----------
app1 xxxx +192(200%)
app2 xxxx -100(100%)

ins1  xxxxx     app   weight   cost   
   ns-app-c1   app1     1      10
   ns-app-c2   app2     1      10
   ns-app-c2   app3     1      10

ins2  xxxxx     30      
   app   app1   10
   app   app2   10
   app   app3   10

