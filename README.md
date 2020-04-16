## 组成部分
* 给UI展示端的API
* 接收日志的push api

## 依赖服务  
Elasticsearch 6.8.6  
Redis   
Mysql

## push Api
### 说明
- 每次请求都需要在请求头添加授权token
```
X-TOKEN: {token}
```
- token在UI展示端添加应用以获取token  

### push api
- 请求地址
```
/push/v1/logs
```

- 方法 POST  
- 请求
```
POST /push/v1/logs
Content-Type：application/json
X-TOKEN： {token}

{"time":1586843267,"level":"error","content":"日志内容"}
```

- 返回信息
```
HTTP/1.1 200 OK
{
    "code": 0,
    "errmsg": ""
}

```



## push api压测 
> 服务器基础配置：    
> 阿里云 ecs.g6.2xlarge   
> 8 vCPU 32GiB  
> 本机部署ES

c 100 n 10000
![pic](https://easywen.oss-cn-beijing.aliyuncs.com/eccang/log/%E6%80%A7%E8%83%BD2.png)


c 5000 n 100000
![pic](https://easywen.oss-cn-beijing.aliyuncs.com/eccang/log/%E6%80%A7%E8%83%BD1.png)


