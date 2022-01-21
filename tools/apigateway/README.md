## 官网网址链接
https://konghq.com//  
安装说明：https://docs.konghq.com/gateway/2.7.x/install-and-run/macos/
## 介绍
kong可以用来做什么？
1. 使用 Service 和 Route 对象公开您的服务
2. 设置速率限制和代理缓存
3. 使用密钥身份验证保护服务
4. 负载均衡流量
## 安装
### mac安装
#### 安装postgresql
    brew install postgresql
查看安装路径  

    which psql
    ----/usr/local/bin/psql
查看pg版本

    pg_ctl -V

修改配置文件

    nano /usr/local/var/postgres/postgresql.conf
设置host和port,其他使用默认值
>设置host和port,其他使用默认值
把port = 5432和listen_addresses = 'localhost'这两句取消前面的注释。

启动数据库

    pg_ctl -D /usr/local/var/postgres -l /usr/local/var/postgres/server.log start
查看数据库访问日志

    cat /usr/local/var/postgres/server.log
查看数据库运行状态

    pg_ctl -D /usr/local/var/postgres -l /usr/local/var/postgres/server.log status 
停止数据库服务

    pg_ctl -D /usr/local/var/postgres -l /usr/local/var/postgres/server.log stop -s -m fast
查看数据库运行进程

    ps -ef |grep postgres 或 ps auxwww | grep postgres 

创建数据库用户-kong

    createuser kong -P  
之后会提示你输密码，我用kong作为kong的密码。
创建数据库/密码-kong/kong

    createdb kong -O kong -E UTF8 -e
然后就可以通过navicat图形化连接pg数据库kong
#### 安装kong
    brew tap kong/kong
    brew install kong
按照上述postgresql安装，已经准备好存储，现在需要执行kong migrations来初始化数据库表

    kong migrations bootstrap
此时你用navicat图形化连接看数据库就会看到很多新增的数据库表了
此时我们就是按照默认配置进行安装配置完毕，如果要更加的个性化，可以看官方的文档。
#### 安装kong-dashboard
UI管理界面

项目地址：https://github.com/PGBI/kong-dashboard

此项目停止更新了只到2.0.0以下版本
项目地址2:https://github.com/m1h43l/kong-dashboard  
这个也是只支持2.x.x版本。
我是2.7.0版本没成功，虽然跑起来了
## 使用说明
### 公开服务  
\<admin-hostname\>是你的kong服务ip
开启kong

    kong start [-c /path/to/kong.conf]
验证开启

    curl -i http://localhost:8001/
注册服务

    curl -i -X POST http://<admin-hostname>:8001/services \
    --data name=example_service \
    --data url='http://mockbin.org'
如果服务创建成功，您将收到 201 成功消息  
验证服务端点

    curl -i http://<admin-hostname>:8001/services/example_service
添加route  

    curl -i -X POST http://<admin-hostname>:8001/services/example_service/routes \
    --data 'paths[]=/mock' \
    --data name=mocking
一条201消息指示路由已成功创建  
验证路由是否将请求转发到服务  
默认情况下，Kong Gateway 处理 port 上的代理请求:8000
直接访问  
http://\<admin-hostname\>:8000/mock 

或者

    curl -i -X GET http://<admin-hostname>:8000/mock/request
我们都做了什么？  
1. 添加了一个以example_serviceURL命名的服务http://mockbin.org。
2. 添加了一个名为/mock.
3. 这意味着如果一个 HTTP 请求被发送到端口8000（代理端口）上的 Kong Gateway 节点并且它与 route 匹配/mock，那么该请求将被发送到http://mockbin.org.
4. 抽象了后端/上游服务，并将您选择的路由放在前端，您现在可以将其提供给客户端以发出请求
### 保护服务
Kong 的Rate Limiting 插件可让您限制上游服务从 API 消费者接收的请求数量，或者每个用户调用 API 的频率。  
**为什么使用速率限制？**  
速率限制可防止 API 意外或恶意过度使用。在没有速率限制的情况下，每个用户都可以根据自己的喜好频繁地请求，这可能会导致请求高峰，从而使其他消费者挨饿。启用速率限制后，API 调用被限制为每秒固定数量的请求。  
**设置速率限制**  
在端口上调用 Admin API8001并配置插件，以在节点上启用每分钟五 (5) 个请求的限制，这些请求存储在本地和内存中

    curl -i -X POST http://<admin-hostname>:8001/plugins \
    --data name=rate-limiting \
    --data config.minute=5 \
    --data config.policy=local
**验证速率限制**  
输入\<admin-hostname>:8000/mock并刷新浏览器六次。在第 6 次请求之后，您将收到一条错误消息。
### 提高性能

