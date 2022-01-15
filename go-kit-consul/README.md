## go-kit-consul

## 环境设置
GOPATH是一个环境变量，用来表明你写的go项目的存放路径
GOPATH路径最好只设置一个，所有的项目代码都放到GOPATH的src目录下。
bin：存放编译后生成的可执行文件
pkg：存放编译后生成的归档文件
src：存放源码文件
>在进行Go语言开发的时候，我们的代码总是会保存在$GOPATH/src目录下。在工程经过go build、go install或go get等指令后，会将下载的第三方包源代码文件放在$GOPATH/src目录下， 产生的二进制可执行文件放在 $GOPATH/bin目录下，生成的中间缓存文件会被保存在 $GOPATH/pkg 下  

>废弃(此为go version 1.11之前)
>>设置你的GOPATH路径
>>>      cd go-kit-consul
>>>      pwd
>>>        /此时你的工作目录是你的工作目录/
>>>        export GOPATH=/Users/zp/mygit/go-kit/go-kit-consul
>>>        go env 

新版本 使用go.mod形式管理包  

        go mod init go-kit-consul
此时如果你用Go的IDEA编译环境，请你打开根目录，比如我的根目录应该是go-kit-consul
此时导入自建包就可以直接用下列语句
>import（  
>  "go-kit-consul/EndPoint"  
） 
 

此时输出go的一些参数，GOPATH现在应该就是你的工作目录  

        go get -u github.com/gorilla/mux
        go get -u github.com/go-kit/kit/transport/http
## 使用
在根目录下
        go run main.go  
之后访问浏览器可以看到我们的demo了  
http://127.0.0.1:8000/?name=zp
http://127.0.0.1:8001/?name=zp
