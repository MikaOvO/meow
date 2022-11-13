# 11.13

## 做法

```
service docker start
docker pull mongo:4.0  
docker run --name mymongo -v $PWD/mymongo/data:/data/db -p 27017:27017 -d mongo:4.0 

docker ps -a # 查看是否启动成功
```

```--name``` 代表启动的容器名字

```-d``` 代表启动的镜像名字

```-p``` 代表端口映射

```-v``` 代表宿主机和docker的文件路径映射，即使停止了也可以从宿主机读取，确保$PWD加入了docker的文件共享路径

复制```./api```下的内容，利用goctl生成代码，解决需要的包（goland自动）
```
goctl api go -api ./api/main.api -dir .     
```

添加./internal/logic下的逻辑

启动服务
```
go run user.go -f etc/user-api.yaml      
```

测试（需要curl）

<img width="574" alt="image" src="https://user-images.githubusercontent.com/93330615/201518297-9b4a5f83-f89e-4eac-9cb8-e212c4765af1.png">

注意mongo中小写存储，bson.D查询需要小写key

## 问题

* 代码相当冗余
* 如何把mongo客户端连接部分拆分出来，在user.go中只启动k次（单例/连接池？）
* 鉴权
