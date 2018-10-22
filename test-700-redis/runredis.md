# 开启docker
# 下载redis镜像

`docker search redis` // 搜索redis镜像

`docker pull redis` // 下载last版本redis

`docker images redis` // 列出本地的redis镜像


# 启动redis镜像

`docker run -p6379:6379 -v $PWD/data:/data -d redis redis-server --appendonly yes`

参数说明：

- -p 6379:6379 将容器的6379端口映射到主机的6379端口
- -v $PWD/data:/data 将主机中当前目录下的data挂载到容器的/data
- redis-server --appendonly yes  在容器执行redis-server启动命令，并打开redis持久化配置

# 查看容器
`docker ps | grep redis`

# 连接，查看容器
使用redis镜像执行redis-cli 命令连接到刚启动的容器，主机ip为127.0.0.1

`docker exec -it 43f7a65ec7f8 redis-cli`


