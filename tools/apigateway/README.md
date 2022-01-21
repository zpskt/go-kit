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
**什么是代理缓存？**  
Kong Gateway 通过缓存提供快速性能。代理缓存插件使用反向代理缓存实现提供了这种快速性能。它根据请求方法、可配置的响应代码、内容类型缓存响应实体，并且可以缓存每个消费者或每个 API。

缓存实体存储一段可配置的时间。当达到超时时，网关将请求转发到上游，缓存结果并从缓存中响应，直到超时。该插件可以将缓存数据存储在内存中，或者为了提高性能，在 Redis 中。  
**为什么使用代理缓存？**  
使用代理缓存，以便上游服务不会因重复请求而陷入困境。使用代理缓存，Kong Gateway 可以响应缓存结果以获得更好的性能。  
**设置代理缓存插件**  

    curl -i -X POST http://<admin-hostname>:8001/plugins \
    --data name=proxy-cache \
    --data config.content_type="application/json; charset=utf-8" \
    --data config.cache_ttl=30 \
    --data config.strategy=memory

**验证代理缓存**  
使用 Admin API访问/mock路由并记下响应标头：  
    curl -i -X GET http:\<admin-hostname>:8000/mock/request
特别要密切注意X-Cache-Status、X-Kong-Proxy-Latency和的值X-Kong-Upstream-Latency：
>HTTP/1.1 200 OK  
...  
X-Cache-Key: d2ca5751210dbb6fefda397ac6d103b1  
X-Cache-Status: Miss  
X-Content-Type-Options: nosniff  
...  
X-Kong-Proxy-Latency: 25  
X-Kong-Upstream-Latency: 37  

接下来，再次访问/mock路由  
这一次，请注意X-Cache-Status的X-Kong-Proxy-Latency差异X-Kong-Upstream-Latency。缓存状态是 hit，这意味着 Kong Gateway 直接从缓存响应请求，而不是将请求代理到上游服务。

此外，请注意响应中的最小延迟，这使 Kong Gateway 能够提供最佳性能：
>HTTP/1.1 200 OK  
...  
X-Cache-Key: d2ca5751210dbb6fefda397ac6d103b1  
X-Cache-Status: Hit  
...  
X-Kong-Proxy-Latency: 0  
X-Kong-Upstream-Latency: 1  

为了更快地测试，可以通过调用 Admin API 来删除缓存：  
    
    curl -i -X DELETE http://<admin-hostname>:8001/proxy-cache
### 安全服务  
API 网关身份验证、设置密钥身份验证插件并添加使用者  
**什么是身份验证？**  
API 网关身份验证是控制允许使用您的 API 传输的数据的重要方法。基本上，它使用一组预定义的凭据检查特定消费者是否有权访问 API。

Kong Gateway 有一个插件库，提供简单的方法来实现最知名和最广泛使用的 API 网关身份验证方法。以下是一些常用的：

1. 基本认证
2. 密钥认证
3. OAuth 2.0 身份验证
4. LDAP 身份验证高级
5. OpenID 连接  
   
可以将身份验证插件配置为应用于 Kong 网关中的服务实体。反过来，服务实体与它们所代表的上游服务一对一映射，本质上意味着身份验证插件直接应用于这些上游服务。  
**为什么使用 API 网关身份验证？**  
启用身份验证后，Kong Gateway 不会代理请求，除非客户端首先成功进行身份验证。这意味着上游 (API) 不需要对客户端请求进行身份验证，也不会浪费验证凭据的关键资源。

Kong Gateway 可以查看所有身份验证尝试（成功、失败等），从而能够对这些事件进行编目和仪表板，以证明正确的控制措施到位，并实现合规性。身份验证还使您有机会确定如何处理失败的请求。这可能意味着简单地阻止请求并返回错误代码，或者在某些情况下，您可能仍希望提供有限的访问权限。

在此示例中，您将启用密钥身份验证插件。API 密钥认证是进行 API 认证最流行的方式之一，可以根据需要实现创建和删除访问密钥。  

**设置密钥验证插件**  
在端口上调用 Admin API8001并配置插件以启用密钥身份验证。对于此示例，将插件应用于您创建的/mock路由：

    curl -X POST http://<admin-hostname>:8001/routes/mocking/plugins \
    --data name=key-auth
再次尝试访问服务：

    curl -i http://<admin-hostname>:8000/mock  
由于您添加了密钥身份验证，您应该无法访问它：
>HTTP/1.1 401 Unauthorized  
...  
{  
    "message": "No API key found in request"  
}  

在 Kong 代理请求此路由之前，它需要一个 API 密钥。对于此示例，由于您安装了 Key Authentication 插件，因此您需要先创建一个具有关联密钥的使用者。  
**设置消费者和凭证**  
要创建消费者，请调用 Admin API 和消费者的端点。以下创建了一个名为consumer的新消费者：

    curl -i -X POST http://<admin-hostname>:8001/consumers/ \
    --data username=consumer \
    --data custom_id=consumer

配置后，调用 Admin API 为上面创建的使用者配置密钥。对于此示例，将密钥设置为apikey 

    curl -i -X POST http://<admin-hostname>:8001/consumers/consumer/key-auth \
    --data key=apikey
如果没有输入密钥，Kong 会自动生成密钥.  
结果：
>HTTP/1.1 201 Created  
...  
{  
    "consumer": {  
        "id": "2c43c08b-ba6d-444a-8687-3394bb215350"  
    },  
    "created_at": 1568255693,  
    "id": "86d283dd-27ee-473c-9a1d-a567c6a76d8e",  
    "key": "apikey"  
}

您现在拥有一个配置 API 密钥的消费者来访问该路由。  
**验证密钥身份验证**  

用浏览器：  
要验证密钥身份验证插件，请通过浏览器访问您的路由，方法是附加?apikey=apikey到 url：

http://\<admin-hostname>:8000/mock/?apikey=apikey

用api:  
要验证密钥身份验证插件，请再次访问模拟路由，使用apikey密钥值为apikey.

    curl -i http://<admin-hostname>:8000/mock/request \
    -H 'apikey:apikey'  

**（可选）禁用插件**
如果您按主题遵循此入门指南，则需要在以后的任何请求中使用此 API 密钥。如果您不想继续指定密钥，请在继续之前禁用插件。  
找到插件 ID 并复制它：

        curl -X GET http://<admin-hostname>:8001/routes/mocking/plugins/
或者访问浏览器  
http://localhost:8001/routes/mocking/plugins  
会输出id  
禁用插件  

    curl -X PATCH http://<admin-hostname>:8001/routes/mocking/plugins/{<plugin-id>} \
    --data enabled=false

总结：
1. 启用密钥验证插件。
2. 创建了一个名为consumer.
3. 给消费者一个 API 密钥，apikey以便它可以/mock通过身份验证访问路由。

### 设置智能负载均衡
**什么是上游？**  
上游对象是指位于 Kong 网关后面的上游 API/服务，客户端请求被转发到该服务。在 Kong Gateway 中，Upstream Object 代表一个虚拟主机名，可用于对多个服务（目标）的传入请求进行健康检查、断路和负载平衡。

在本主题中，您将配置之前创建的服务 ( example_service) 以指向上游而不是主机。出于我们示例的目的，上游将指向两个不同的目标，httpbin.org并且mockbin.org. 在真实环境中，上游将指向在多个系统上运行的同一个服务。  
**为什么要跨上游目标进行负载平衡？**  
在以下示例中，您将使用跨两个不同服务器或上游目标部署的应用程序。Kong Gateway 需要在两台服务器之间进行负载平衡，以便如果其中一台服务器不可用，它会自动检测问题并将所有流量路由到工作服务器。
**配置上游服务**  
在本节中，您将创建一个名为 Upstream 的上游upstream并向其添加两个目标。
在端口上调用 Admin API8001并创建一个名为的上游upstream：

    curl -X POST http://<admin-hostname>:8001/upstreams \
    --data name=upstream
更新您之前创建的服务以指向此上游：  

    curl -X PATCH http://<admin-hostname>:8001/services/example_service \
    --data host='upstream'  
向上游添加两个目标，每个目标都有端口 80:mockbin.org:80和 httpbin.org:80:

    curl -X POST http://<admin-hostname>:8001/upstreams/upstream/targets \
    --data target='mockbin.org:80'
    curl -X POST http://<admin-hostname>:8001/upstreams/upstream/targets \
    --data target='httpbin.org:80'

您现在有一个带有两个目标的 Upstreamhttpbin.org和mockbin.org，以及一个指向该 Upstream 的服务。  
**验证上游服务**  
http://\<admin-hostname>:8000/mock配置 Upstream 后，通过使用 Web 浏览器或 CLI访问路由来验证它是否正常工作。
继续点击端点，站点应该从httpbin变为mockbin.  
总结：  
创建了一个名为 Upstream 的对象并将upstreamServiceexample_service指向它。  
向上游添加了两个具有相同权重的目标httpbin.org和mockbin.org。