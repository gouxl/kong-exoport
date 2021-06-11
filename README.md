#### 1使用方法
##### 1.1启动插件
该工具用来自定义kong监控指标，在kong中启动httplog插件，并且把日志输出路径设置为：kong-expoter-server/api/kong,例如：
```
curl -X POST http://192.168.75.129:8001/plugins/ \
    --data "name=http-log"  \
    --data "config.http_endpoint=http://192.168.1.9:8080/kong" \   #kong-exporter接受日志地址
    --data "config.method=POST" \
    --data "config.timeout=1000" \
    --data "config.keepalive=1000"
```