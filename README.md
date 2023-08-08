# go-1.20.5
# Stay Hungry , Stay Foolish


#### 总结
docker run -idt --name mysql-hadoop -p 3306:3306 \
-v /etc/localtime:/etc/localtime \
-v /home/mysql/mysql:/var/lib/mysql \
-v /home/mysql/data:/data/mysql/data \
-e MYSQL_ROOT_PASSWORD=123456 \
-e MYSQL_USER=ambari \
-e MYSQL_PASSWORD=ambari \
mysql:5.7.41 \
--character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci --max_allowed_packet=2000M

docker run -d --name myredis -p 6379:6379 redis --requirepass "123456"

* 创建数据库go
* 建表
```sql
CREATE TABLE `blog_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `blog`.`blog_auth` (`id`, `username`, `password`) VALUES (null, 'test', 'test123456');
```
```sql

```
# 思否 文档
https://segmentfault.com/a/1190000013297683


1、 标准库：

fmt：实现了类似C语言printf和scanf的格式化I/O。格式化动作（'verb'）源自C语言但更简单
net/http：提供了HTTP客户端和服务端的实现
2、 Gin：

gin.Default()：返回Gin的type Engine struct{...}，里面包含RouterGroup，相当于创建一个路由Handlers，可以后期绑定各类的路由规则和函数、中间件等
router.GET(...){...}：创建不同的HTTP方法绑定到Handlers中，也支持POST、PUT、DELETE、PATCH、OPTIONS、HEAD 等常用的Restful方法
gin.H{...}：就是一个map[string]interface{}
gin.Context：Context是gin中的上下文，它允许我们在中间件之间传递变量、管理流、验证JSON请求、响应JSON请求等，在gin中包含大量Context的方法，例如我们常用的DefaultQuery、Query、DefaultPostForm、PostForm等等
3、 &http.Server和ListenAndServe？

例如main.go是否在gin-blog/的根目录下，参照一下文章最后的"当前目录结构"

go build 打包生成可执行程序时，只是打包go文件，而不包括其他的比如ini文件，所以ini文件可以随时修改达到热更新的效果
golangs：@golangs 太粗心了，仔细又去看了一下文档，go build 是只编译go文件的。

page, _ := com.StrTo(c.Query("page")).Int()这个为什么不直接用strconv.Atoi(string)

我们创建了一个Tag struct{}，用于Gorm的使用。并给予了附属属性json，这样子在c.JSON的时候就会自动转换格式，非常的便利
可能会有的初学者看到return，而后面没有跟着变量，会不理解；其实你可以看到在函数末端，我们已经显示声明了返回值，这个变量在函数体内也可以直接使用，因为他在一开始就被声明了
有人会疑惑db是哪里来的；因为在同个models包下，因此db *gorm.DB是可以直接使用的


这在构建镜像的时候，非常不利于调试。
解决方法
在 docker build 命令前面增加 DOCKER_BUILDKIT=0，即可跟之前的 Docker 版本一样打印命令输出。
DOCKER_BUILDKIT=0  docker build -t gin-blog-docker .


