#云原生作业

##模块二作业
```
编写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把 1 都做完。
1. 接收客户端 request，并将 request 中带的 header 写入 response header
2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4. 当访问 localhost/healthz 时，应返回 200
```
###作业地址
[模块二作业](https://github.com/501miles/cnstudy/tree/master/week2)


##模快三作业
###作业要求
```
·构建本地镜像。
·编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化（请思考有哪些最佳实践可以引入到 Dockerfile 中来）。
·将镜像推送至 Docker 官方镜像仓库。
·通过 Docker 命令本地启动 httpserver。
·通过 nsenter 进入容器查看 IP 配置。
·作业需编写并提交 Dockerfile 及源代码。
```
###作业地址
[模块三作业](https://github.com/501miles/cnstudy/tree/master/week3)

