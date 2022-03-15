# 资源管理

## 资源的统一搜索


## 资源标签管理


## 资源的监控发现

### prometheus 自动发现

我们使用2个标签:
+ prometheus.io/scrape = "true" 
+ prometheus.io/port = "9100"

比如搜索 prometheus.io/scrape = "true"

```json
[
    {
        "targets": ["10.0.10.2:9100", "10.0.10.3:9100", "10.0.10.4:9100", "10.0.10.5:9100"],
        "labels": {
            "domain": "admin",
            "namespace": "default",
            "env": "prod",
            "accout": "acount11",
            "vendor": "ali_yun",
            "region": "hangzou",
            "instance_id":"ins-xxxxx"
        }
    },
    ...
]
```