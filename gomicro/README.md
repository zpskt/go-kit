手动添加服务


curl \
--request PUT \
--data @myservice.json \
localhost:8500/v1/agent/service/register

手动取消注册
curl \
--request PUT \
localhost:8500/v1/agent/service/deregister/注册时候的ID

下载consul相关api 
        go getgithub.com/hashicorp/consul