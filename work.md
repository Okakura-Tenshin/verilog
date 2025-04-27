### 本项目是作为go后端开发的上手项目，其目的是使用go ＋gin ＋ grom ＋mysql做一个前端页面

router是本项目的路由管理，我们引用的是gin中的中间件gin.Engine，gin还包括

 Gin 是一个用于构建 Go 应用程序的 web 框架,它提供了许多中间件,包括:

1. 静态文件服务中间件:用于提供静态文件服务,如 CSS、JavaScript 和图片等。

```go
r := gin.Default()
r.Use(gin.Static("/static", "./static"))
```

2. 模板渲染中间件:用于渲染 HTML 模板文件。

```go
r := gin.Default()
r.LoadHTMLGlob("templates/*")
r.Use(gin.HTML(gin.Default()))
```

3. 请求参数解析中间件:用于解析请求参数,如 JSON、XML 和表单数据等。

```go
r := gin.Default()
r.Use(gin.BindJSON(MyStruct{}))
```

4. 路由分组中间件:用于对路由进行分组,可以对一组路由应用相同的中间件。

```go
r := gin.Default()
v1 := r.Group("/v1")
v1.Use(gin.Logger())
v1.GET("/hello", func(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "Hello, World!",
    })
})
```

5. 中间件栈:可以组合多个中间件,形成一个中间件栈,依次应用这些中间件。

```go
r := gin.Default()
r.Use(gin.Logger(), gin.Recovery())
```

这些中间件可以单独使用,也可以组合使用,根据实际需求进行选择和配置。

#### 数据库配置

```
docker run --name mariadb1231 -e MYSQL_ROOT_PASSWORD=123456  -d mariadb
```

要给高权限：docker没配置好

```
docker run --name mysql123 --privileged -e MYSQL_ROOT_PASSWORD=123456 -d mysql:latest

docker exec -it mysql123 mysql -u root -p 

docker exec -it mysql mysql -u root -p123456
```

#### 后台管理

```
使用 tmux：
创建一个新的 tmux 会话：

bash
复制代码
tmux new -s mysession
在 tmux 会话中运行你希望持续运行的命令：

bash
复制代码
your-command
要将 tmux 会话从当前终端分离（后台运行），按下：

bash
复制代码
Ctrl + B 然后按 D
重新连接到 tmux 会话：

bash
复制代码
tmux attach -t mysession
退出 tmux 会话： 输入 exit，或按下 Ctrl + B 然后按 X。
```
