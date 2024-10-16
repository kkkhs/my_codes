# Gin

[TOC]

## 一、gin介绍

### 1、gin特性

#### 1、快速

> 基于 Radix 树的路由，小内存占用。没有反射。可预测的 API 性能。

#### 2、支持中间件

> 传入的 HTTP 请求可以由一系列中间件和最终操作来处理。 例如：Logger，Authorization，GZIP，最终操作 DB。

#### 3、Crash 处理

> Gin 可以 catch 一个发生在 HTTP 请求中的 panic 并 recover 它。这样，你的服务器将始终可用。例如，你可以向 Sentry 报告这个 panic！

#### 4、JSON 验证

> Gin 可以解析并验证请求的 JSON，例如检查所需值的存在。

#### 5、路由组

> 更好地组织路由。是否需要授权，不同的 API 版本…… 此外，这些组可以无限制地嵌套而不会降低性能。

#### 6、错误管理

> Gin 提供了一种方便的方法来收集 HTTP 请求期间发生的所有错误。最终，中间件可以将它们写入日志文件，数据库并通过网络发送。

#### 7、内置渲染

> Gin 为 JSON，XML 和 HTML 渲染提供了易于使用的 API。

#### 8、可扩展性

> 新建一个中间件非常简单。

### 2、gin快速入门

第一步：go get github.com/gin-gonic/gin

示例程序：

```go
package main

import "github.com/gin-gonic/gin"

func main() {
    //创建一个路由服务
	route := gin.Default()
    //发送数据给前端
	route.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	route.Run() // 监听并在 127.0.0.1:8080 上启动服务
}
```

## 二、路由

路由是UR到函数的映射。

一个URI含: `http://localhost:8080/user/find?id=11`

> - <font color='red'>协议</font>，比如http，https等
> - <font color='red'>ip端口或者域名</font>，比如127.0.0.1:8080或者`www.test.com`
> - <font color='red'>path</font>，比如 /path
> - <font color='red'>query</font>，比如 ?query

同时访问的时候，还需要指明<font color='red'>HTTP METHOD</font>，比如：

> - GET：请求一个指定资源的表示形式. 使用GET的请求应该只被用于获取数据.
> - POST：用于将实体提交到指定的资源，通常会导致在服务器上的状态变化
> - HEAD：请求一个与GET请求的响应相同的响应，但没有响应体.
> - PUT：用请求有效载荷替换目标资源的所有当前表示
> - DELETE：删除指定的资源
> - CONNECT：建立一个到由目标资源标识的服务器的隧道。
> - OPTIONS：用于描述目标资源的通信选项。
> - TRACE：沿着到目标资源的路径执行一个消息环回测试。
> - PATCH：用于对资源应用部分修改。

> 使用的时候，应该尽量遵循其语义

### 1. RESTful API规范

RESTful API 的规范建议我们使用特定的HTTP方法来对服务器上的资源进行操作。

比如：

> 1. GET，表示<font color='red'>获取</font>服务器上的资源
> 2. POST，表示在服务器上<font color='red'>创建</font>资源
> 3. PUT,表示<font color='red'>更新</font>或者<font color='red'>替换</font>服务器上的资源
> 4. DELETE，表示<font color='red'>删除</font>服务器上的资源
> 5. PATCH，表示<font color='red'>更新/修改</font>资源的<font color='red'>一部分</font>

### 2. 请求方法

**比如：**

```go
	r.GET("/get", func(ctx *gin.Context) { //Get方法
		ctx.JSON(200, "get")
	})
	r.POST("/post", func(ctx *gin.Context) { //Post方法
		ctx.JSON(200, "post")
	})
	r.DELETE("/delete", func(ctx *gin.Context) { //Delete方法
		ctx.JSON(200, "delete")
	})
	r.PUT("/put", func(ctx *gin.Context) { //Put方法
		ctx.JSON(200, "put")
	})
```

如果想要支持<font color='red'>所有</font>：

```go
r.Any("/any", func(ctx *gin.Context) { //Any方法
		ctx.JSON(200, "any")
	})
```



如果想要支持其中的<font color='red'>几种</font>：

```go
   //统一URI下绑定多种方法
	r.GET("/hello", func(ctx *gin.Context) {
		//数组 map list 结构体
		ctx.JSON(200, gin.H{
			"name": "hello world",
		})
	})
	r.POST("/hello", func(ctx *gin.Context) {
		//数组 map list 结构体
		ctx.JSON(200, gin.H{
			"name": "hello world",
		})
	})
```

### 3. URI

URI书写的时候，我们不需要关心scheme和authority这两部分，我们主要通过path和query两部分的书写来进行资源的定位。

- 静态url，比如`/hello`，`/user/find`

	```go
	r.POST("/user/find", func(ctx *gin.Context) {
	})
	```

- 路径参数，比如`/user/find/:id`

	```go
	r.POST("/user/find/:id", func(ctx *gin.Context) {
			param := ctx.Param("id")
			ctx.JSON(200, param)
		})
	```

- 模糊匹配，比如`/user/*path`

	```go
	r.POST("/user/*path", func(ctx *gin.Context) {
			param := ctx.Param("path")
			ctx.JSON(200, param)
	})
	```

### 4. 处理函数

**定义：**

```go
type HandlerFunc func(ctx *Context)
```

通过上下文的参数，获取http的请求参数，响应http请求等。

### 5. 分组路由

> 在进行开发的时候，我们往往要进行<font color='red'>模块</font>的划分，<font color='red'>方便统一管理</font>，如用户模块，以user开发，商品模块，以goods开头。
>
> 或者进行多版本开发，不同版本之间路径是一致的，这种时候，就可以用到分组路由了。
>

如：以下的路由前面统一加上api的前缀

```go
   ug := r.Group("/user")
	{
		ug.GET("/find", func(ctx *gin.Context) {
			ctx.JSON(200, "user find")
		})
		ug.POST("/save", func(ctx *gin.Context) {
			ctx.JSON(200, "user save")
		})
	}
	gg := r.Group("/goods")
	{
		gg.GET("/find", func(ctx *gin.Context) {
			ctx.JSON(200, "goods find")
		})
		gg.POST("/save", func(ctx *gin.Context) {
			ctx.JSON(200, "goods save")
		})
	}
```

## 三、请求参数

### 1. Get请求参数

使用Get请求传参时，类似于这样 `http://localhost:8080/user/save?id=11&name=zhangsan`。

如何获取呢？

#### 1.1 普通参数

<font color='red'>Query </font>----> request url: `http://localhost:8080/user/save?id=11&name=zhangsan`

```go
r.GET("/user/save", func(ctx *gin.Context) {
		id := ctx.Query("id")
		name := ctx.Query("name")
		ctx.JSON(200, gin.H{
			"id":   id,
			"name": name,
		})
	})
```

<font color='red'>DefaultQuery </font>----> 如果参数不存在，就给一个默认值：

```go
r.GET("/user/save", func(ctx *gin.Context) {
		id := ctx.Query("id")
		name := ctx.Query("name")
		address := ctx.DefaultQuery("address", "北京")
		ctx.JSON(200, gin.H{
			"id":      id,
			"name":    name,
			"address": address,
		})
	})
```

<font color='red'>GetQuery </font>----> 判断参数是否存在：

```go
r.GET("/user/save", func(ctx *gin.Context) {
		id, ok := ctx.GetQuery("id")
		address, aok := ctx.GetQuery("address")
		ctx.JSON(200, gin.H{
			"id":      id,
			"idok":    ok,
			"address": address,
			"aok":     aok,
		})
	})
```

<font color='red'>BindQuery </font>----> id是数值类型，上述获取的都是string类型，根据类型获取：

```go
type User struct {
	Id   int64  `form:"id"`
	Name string `form:"name"`
}
r.GET("/user/save", func(ctx *gin.Context) {
		var user User
		err := ctx.BindQuery(&user)
		if err != nil {
			log.Println(err)
		}
		ctx.JSON(200, user)
})
```

<font color='red'>ShouldBindQuery </font>----> 也可以：

```go
r.GET("/user/save", func(ctx *gin.Context) {
		var user User
		err := ctx.ShouldBindQuery(&user)
		if err != nil {
			log.Println(err)
		}
		ctx.JSON(200, user)
	})
```

**区别：**

```go
type User struct {
	Id      int64  `form:"id"`
	Name    string `form:"name"`
	Address string `form:"address" binding:"required"`
}
```

当bind是必须的时候，ShouldBindQuery会报错，开发者自行处理，<font color='red'>状态码不变</font>。

BindQuery则报错的同时，会将<font color='red'>状态码改为400</font>。所以一般建议是使用Should开头的bind。

#### 1.2 数组参数

<font color='red'>QueryArray</font> ——> 请求url：`http://localhost:8080/user/save?address=Beijing&address=shanghai`

```go
r.GET("/user/save", func(ctx *gin.Context) {
		address := ctx.QueryArray("address")
		ctx.JSON(200, address)
	})
```

<font color='red'>GetQueryArray</font> ——> 判断参数是否存在

```go
r.GET("/user/save", func(ctx *gin.Context) {
		address, ok := ctx.GetQueryArray("address")
		fmt.Println(ok)
		ctx.JSON(200, address)
	})
```

<font color='red'>ShouldBindQuery</font> ——> 根据类型获取

```go
r.GET("/user/save", func(ctx *gin.Context) {
		var user User
		err := ctx.ShouldBindQuery(&user)
		fmt.Println(err)
		ctx.JSON(200, user)
	})
```

#### 1.3 map参数

<font color='red'>QueryMap</font> ——> 请求url：`http://localhost:8080/user/save?addressMap[home]=Beijing&addressMap[company]=shanghai`

```go
r.GET("/user/save", func(ctx *gin.Context) {
		addressMap := ctx.QueryMap("addressMap")
		ctx.JSON(200, addressMap)
	})
```

<font color='red'>GetQueryMap</font> ——> 判断参数是否存在

```go
r.GET("/user/save", func(ctx *gin.Context) {
		addressMap, _ := ctx.GetQueryMap("addressMap")
		ctx.JSON(200, addressMap)
	})
```

map参数 bind并<font color='red'>没有</font>支持

#### 1.4 动态参数

<font color='red'>Param</font> ——>  请求url：`http://localhost:8080/user/save/111`

```go
r.GET("/user/save/:id", func(ctx *gin.Context) {
		ctx.JSON(200, ctx.Param("id"))
	})
```



### 2. Post请求参数

post请求一般是<font color='red'>表单参数</font>和<font color='red'>json参数</font>

#### 2.1 表单参数

![image-20221121163103752](https://www.mszlu.com/assets/image-20221121163103752.9bd84f58.png)

<font color='red'>PostForm</font> ——> 

```go
r.POST("/user/save", func(ctx *gin.Context) {
		id := ctx.PostForm("id")
		name := ctx.PostForm("name")
		address := ctx.PostFormArray("address")
		addressMap := ctx.PostFormMap("addressMap")
		ctx.JSON(200, gin.H{
			"id":         id,
			"name":       name,
			"address":    address,
			"addressMap": addressMap,
		})
	})
```

<font color='red'>ShouldBind</font> ——> 根据类型获取：

<font color='red'>GetPostFormMap</font> ———> 获取表单Map数据

```go
r.POST("/user/save", func(ctx *gin.Context) {
		var user User
		err := ctx.ShouldBind(&user)
		addressMap, _ := ctx.GetPostFormMap("addressMap")
		user.AddressMap = addressMap
		fmt.Println(err)
		ctx.JSON(200, user)
	})
```

#### 2.2 json参数

```json
{
    //json数据准备
    "id":1111,
    "name":"zhangsan",
    "address": [
        "beijing",
        "shanghai"
    ],
    "addressMap":{
        "home":"beijing"
    }
}
```

<font color='red'>GetRawData </font>——> 根据行获取

<font color='red'>Unmarshal </font>——> Json反序类化

```go
r.POST("/json", func(ctx *gin.Context) {
		//request.body
		data, _ := ctx.GetRawData()

		var m map[string]interface{}
		_ = json.Unmarshal(data, &m)
		ctx.JSON(http.StatusOK, m)
	})
```

### 3. 路径参数

<font color='red'>Param</font> ——>  请求url：`http://localhost:8080/user/save/111`

```go
r.POST("/user/save/:id", func(ctx *gin.Context) {
		ctx.JSON(200, ctx.Param("id"))
	})
```

### 4. 文件参数

```go
r.POST("/user/save", func(ctx *gin.Context) {
		form, err := ctx.MultipartForm()
		if err != nil {
			log.Println(err)
		}
		files := form.File
		for _, fileArray := range files {
			for _, v := range fileArray {
				ctx.SaveUploadedFile(v, "./"+v.Filename)
			}

		}
		ctx.JSON(200, form.Value)
	})
```



## 四、响应

### 1. 字符串方式

```go
r.GET("/user/save", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "this is a %s", "ms string response")
	})
```

### 2. JSON方式

```go
r.GET("/user/save", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})
```

### 3. XML方式

```go
type XmlUser struct {
	Id   int64  `xml:"id"`
	Name string `xml:"name"`
}
r.GET("/user/save", func(ctx *gin.Context) {
		u := XmlUser{
			Id:   11,
			Name: "zhangsan",
		}
		ctx.XML(http.StatusOK, u)
	})
```

### 4. 文件格式

<font color='red'>FileAttachment </font>——>

```go
r.GET("/user/save", func(ctx *gin.Context) {
		//ctx.File("./1.png")
		ctx.FileAttachment("./1.png", "2.png")
	})
```

### 5. 设置http响应头

HTTP响应头是HTTP响应消息头的组成部分之一，可<font color='red'>携带特定响应参数</font>并传递给客户端

<font color='red'>Header</font> ——> 

```go
r.GET("/user/save", func(ctx *gin.Context) {
		ctx.Header("test", "headertest")
	})
```

### 6. 重定向

<font color='red'>Redirect</font> ——> 网页跳转

```go
r.GET("/user/save", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})
```

### 7. YAML方式

```go
r.GET("/user/save", func(ctx *gin.Context) {
		ctx.YAML(200, gin.H{"name": "ms", "age": 19})
})
```

## 五、模板渲染

模板是golang语言的一个标准库，使用场景很多，gin框架同样支持<font color='red'>模板</font>

### 1. 基本使用

定义一个存放模板文件的`templates`文件夹

```html
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>gin_templates</title>
</head>
<body>
    
获取后端的数据为:
{{.msg}}
</body>
</html>
```

**后端代码：**

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 模板解析LoadHTMLFiles 加载静态页面
	r.LoadHTMLGlob("templates/*") //全局加载
	// r.LoadHTMLFiles("templates/index.html")

    //响应一个页面给前端
	r.GET("/index", func(c *gin.Context) {
		// HTML请求
        // 模板的渲染
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"msg": "hello 模板",
		})
	})

	r.Run(":9090") // 启动server
}
```



### 2. 多个模板渲染

如果有多个模板，可以统一进行渲染

```go
// 模板解析,解析templates目录下的所有模板文件
	r.LoadHTMLGlob("templates/**")
```

如果目录为`templates/post/index.tmpl`和`templates/user/index.tmpl`这种，可以

```go
	// **/* 代表所有子目录下的所有文件
	router.LoadHTMLGlob("templates/**/*")
```

### 3. 自定义模板函数

```go
   // gin框架给模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	// 模板解析,解析templates目录下的所有模板文件
	r.LoadHTMLGlob("templates/**")

	r.GET("/index", func(c *gin.Context) {
		// HTML请求
		// 模板的渲染
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "<a href='http://baidu.com'>跳转到其他地方</a>",
		})
	})
```



```go
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>gin_templates</title>
</head>
<body>
{{.title | safe}}
</body>
</html>
```



### 4. 静态文件处理

如果在模板中引入静态文件，比如样式文件

./static/style.css

```css
body{
    background-color: aqua;
}
```

./static/commen.js

```go
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>gin_templates</title>
    <link rel="stylesheet" href="/css/index.css"> //引入css
	<script src="/static/js/commen.js"></script>  //引入js
</head>
<body>
{{.title}}
</body>
</html>
```



```go
//加载静态资源文件
	r.Static("/static", "./static")
```

## 六、会话

`会话控制涉及到cookie和session的使用`

###  1. cookie

> 1. HTTP是<font color='red'>无状态协议</font>，服务器不能记录浏览器的访问状态，也就是说服务器不能区分两次请求是否由同一个客户端发出
> 2. Cookie就是<font color='red'>解决HTTP协议无状态</font>的方案之一
> 3. Cookie实际上就是服务器保存在浏览器上的一段信息。浏览器有了Cookie之后，每次向服务器发送请求时都会同时将该信息发送给服务器，服务器收到请求后，就可以根据该信息处理请求
> 4. Cookie由<font color='red'>服务器创建</font>，并发送给浏览器，最终由浏览器保存

#### 1.1 设置cookie的函数

```go
func (c *Context) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
```

**参数说明：**

| 参数名   | 类型   | 说明                                                         |
| -------- | ------ | ------------------------------------------------------------ |
| name     | string | cookie名字                                                   |
| value    | string | cookie值                                                     |
| maxAge   | int    | 有效时间，单位是秒，MaxAge=0 忽略MaxAge属性，MaxAge<0 相当于删除cookie, 通常可以设置-1代表删除，MaxAge>0 多少秒后cookie失效 |
| path     | string | cookie路径                                                   |
| domain   | string | cookie作用域                                                 |
| secure   | bool   | Secure=true，那么这个cookie只能用https协议发送给服务器       |
| httpOnly | bool   | 设置HttpOnly=true的cookie不能被js获取到                      |

```go
r.GET("/cookie", func(c *gin.Context) {
		// 设置cookie
		c.SetCookie("site_cookie", "cookievalue", 3600, "/", "localhost", false, true)
	})
```

#### 1.2 读取cookie

```go
r.GET("/read", func(c *gin.Context) {
		// 根据cookie名字读取cookie值
		data, err := c.Cookie("site_cookie")
		if err != nil {  //未拿到Cookie
			c.String(200,"not found!")
			return
		}
    	// 直接返回cookie值
		c.String(200,data)
	})
```

#### 1.3 删除cookie

通过将cookie的MaxAge设置为-1, 达到删除cookie的目的。

```go
r.GET("/del", func(c *gin.Context) {
		// 设置cookie  MaxAge设置为-1，表示删除cookie
		c.SetCookie("site_cookie", "cookievalue", -1, "/", "localhost", false, true)
		c.String(200,"删除cookie")
	})
```



### 2. Session

- **什么是session**

> ​    session在网络应用中称为“<font color='red'>会话控制</font>”，是服务器为了保存<font color='red'>用户状态</font>而创建的一个特殊的对象。简而言之，session就是一个对象，<font color='red'>用于存储信息</font>。 

- **session有什么用**

>   	  我们先来想一个问题，这个问题就是我们在游览购物网站时，我们<font color='red'>并没有登录</font>，但是我们任然可以将商品加入购物车，并且进行查看，当我们退出游览器后再打开游览器进行查看时，购物车中依然有我们选择的商品，这该怎么实现呢？
>
> ​	    当然，我们可以使用cookie，但是cookie能存放<font color='red'>大量数据</font>吗？这时，我们就需要一种新的技术，Session。session是存储于服务器端的特殊对象，服务器会为每一个游览器(客户端)创建一个<font color='red'>唯一</font>的session。这个session是服务器端共享，每个游览器(客户端)独享的。我们可以在session存储数据，实现<font color='red'>数据共享</font>

- **session的存储形式**

> ​    session类似于一个Map，里面可以存放多个键值对，是以key-value进行存放的。key必须是一个字符串，value是一个对象。



#### 2.1 单个session

在Gin框架中，我们可以依赖[gin-contrib/sessions](https://github.com/gin-contrib/sessions)中间件处理session。

- **安装session包：**

```go
go get github.com/gin-contrib/sessions
```

```go
package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 创建基于cookie的存储引擎，secret 参数是用于加密的密钥
	store := cookie.NewStore([]byte("secret"))
	// 设置session中间件，参数mysession，指的是session的名字，也是cookie的名字
	// store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/hello", func(c *gin.Context) {
		// 初始化session对象
		session := sessions.Default(c)
		// 通过session.Get读取session值
		// session是键值对格式数据，因此需要通过key查询数据
		if session.Get("hello") != "world" {
			fmt.Println("没读到")
			// 设置session数据
			session.Set("hello", "world")
			session.Save()
		}
		c.JSON(200, gin.H{"hello": session.Get("hello")})
	})
	r.Run(":8080")
}
```



#### 2.2 多session

```go
package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	sessionNames := []string{"a", "b"}
	r.Use(sessions.SessionsMany(sessionNames, store))

	r.GET("/hello", func(c *gin.Context) {
		sessionA := sessions.DefaultMany(c, "a")
		sessionB := sessions.DefaultMany(c, "b")

		if sessionA.Get("hello") != "world!" {
			sessionA.Set("hello", "world!")
			sessionA.Save()
		}

		if sessionB.Get("hello") != "world?" {
			sessionB.Set("hello", "world?")
			sessionB.Save()
		}

		c.JSON(200, gin.H{
			"a": sessionA.Get("hello"),
			"b": sessionB.Get("hello"),
		})
	})
	r.Run(":8080")
}
```

#### 2.3 基于redis存储引擎的session

如果我们想将session数据保存到redis中，只要将session的存储引擎改成redis即可。

使用redis作为存储引擎的例子：

- 首先<font color='red'>安装</font>redis存储引擎的包：

```go
go get github.com/gin-contrib/sessions/redis
```

- 代码：

```go
package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 初始化基于redis的存储引擎
	// 参数说明：
	//    第1个参数 - redis最大的空闲连接数
	//    第2个参数 - 数通信协议tcp或者udp
	//    第3个参数 - redis地址, 格式，host:port
	//    第4个参数 - redis密码
	//    第5个参数 - session加密密钥
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/incr", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})
	r.Run(":8080")
}
```

### 3. session和cookie的比较

> - cookie保存在<font color='red'>客户端</font>，session保存在<font color='red'>服务端</font>
> - cookie作用于他所表示的path中(url中要包含path)，范围较小。session代表客户端和服务器的一次会话过程，web页面跳转时也可以共享数据，范围是本次会话，客户端关闭也不会消失。会持续到我们设置的session生命周期结束(默认30min)
> - 我们使用session需要cookie的配合。cookie用来携带JSESSIONID
> - cookie存放的数据量较小，session可以存储更多的信息。
> - cookie由于存放在客服端，相对于session更不<font color='red'>安全</font>
> - 由于session是存放于服务器的，当有很多客户端访问时，肯定会产生大量的session，这些session会对<font color='red'>服务端的性能</font>造成影响。

## 七、中间件

Gin框架允许开发者在处理请求的过程中，加入用户自己的钩子（Hook）函数。这个钩子函数就叫中间件

**Gin中的中间件必须是一个gin.HandlerFunc类型**

在Gin框架中，**中间件**（Middleware）指的是可以<font color='red'>拦截</font>**http请求-响应**生命周期的特殊函数，在请求-响应生命周期中可以注册多个中间件，每个中间件执行不同的功能，一个中间执行完再轮到下一个中间件执行。

**中间件的常见应用场景如下：**

> - 请求限速
> - api接口签名处理
> - 权限校验
> - 统一错误处理

Gin支持设置<font color='red'>全局中间件</font>和<font color='red'>针对路由分组</font>设置中间件，设置全局中间件意思就是会拦截所有请求，针对分组路由设置中间件，意思就是仅对这个分组下的路由起作用。

### 1. 单独注册中间件

```go

// 定义一个中间件
func m1(ctx *gin.Context) {
	fmt.Println("1")
    ctx.JSON(200, gin.H{"msg": "m1"})
	
}

func main() {
	router := gin.Default()

    //m1处于func前面，获得请求后，先走m1后走func
	router.GET("/", m1, func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "index"})
	})
	router.Run()
}
```

### 2. 多个中间件

router.GET，后面可以跟很多HandlerFunc方法，这些方法其实都可以叫中间件

```go

func m1(c *gin.Context) {
  fmt.Println("m1 ...in")
}
func m2(c *gin.Context) {
  fmt.Println("m2 ...in")
}

func main() {
  router := gin.Default()

  router.GET("/", m1, func(c *gin.Context) {
    fmt.Println("index ...")
    c.JSON(200, gin.H{"msg": "响应数据"})
  }, m2)

  router.Run(":8080")
}

/*
m1  ...in
index ...
m2  ...in
*/
```

### 3. 中间件拦截响应

c.Abort()拦截，后续的HandlerFunc就不会执行了

```go

func m1(c *gin.Context) {
  fmt.Println("m1 ...in")
  c.JSON(200, gin.H{"msg": "第一个中间件拦截了"})
  c.Abort()
}
func m2(c *gin.Context) {
  fmt.Println("m2 ...in")
}

func main() {
  router := gin.Default()

  router.GET("/", m1, func(c *gin.Context) {
    fmt.Println("index ...")
    c.JSON(200, gin.H{"msg": "响应数据"})
  }, m2)

  router.Run(":8080")
}
```

###  4. 中间件放行

c.Next()，Next前后形成了其他语言中的<font color='red'>请求中间件</font>和<font color='red'>响应中间件</font>

```go

func m1(c *gin.Context) {
  fmt.Println("m1 ...in")
  c.Next()
  fmt.Println("m1 ...out")
}
func m2(c *gin.Context) {
  fmt.Println("m2 ...in")
  c.Next()
  fmt.Println("m2 ...out")
}

func main() {
  router := gin.Default()

  router.GET("/", m1, func(c *gin.Context) {
    fmt.Println("index ...in")
    c.JSON(200, gin.H{"msg": "响应数据"})
    c.Next()
    fmt.Println("index ...out")
  }, m2)

  router.Run(":8080")
}

/*
m1 ...in
index ...in
m2 ...in   
m2 ...out  
index ...out
m1 ...out
*/
```

![img](http://python.fengfengzhidao.com/pic/20221210220434.png)

如果其中一个中间件响应了c.Abort()，后续中间件将<font color='red'>不再执行</font>，直接按照顺序走完所有的响应中间件

使用Use可以使用gin自带的中间件或者其他第三方中间件，也可以自己开发中间件

###  5. 全局注册中间件

```go

func m10(c *gin.Context) {
  fmt.Println("m1 ...in")
  c.Next()
  fmt.Println("m1 ...out")
}

func main() {
  router := gin.Default()

  //Use使用全局中间件
  router.Use(m10)
    
  router.GET("/", func(c *gin.Context) {
    fmt.Println("index ...in")
    c.JSON(200, gin.H{"msg": "index"})
    c.Next()
    fmt.Println("index ...out")
  })

  router.Run(":8080")

}
CopyErrorOK!
```

使用Use去注册全局中间件，Use接收的参数也是<font color='red'>多个</font>HandlerFunc

###  6. 中间件传递数据

使用<font color='red'>Set设置</font>一个key-value,

在后续中间件中使用<font color='red'>Get接收</font>数据

```go

func m10(c *gin.Context) {
  fmt.Println("m1 ...in")
  c.Set("name", "fengfeng")
}

func main() {
  router := gin.Default()

  router.Use(m10)
  router.GET("/", func(c *gin.Context) {
    fmt.Println("index ...in")
    name, _ := c.Get("name")
    fmt.Println(name)
    
    c.JSON(200, gin.H{"msg": "index"})
  })

  router.Run(":8080")

}
CopyErrorOK!
```

value的类型是any类型，所有我们可以用它传任意类型，在接收的时候做好断言即可

```go

type User struct {
  Name string
  Age  int
}

func m10(c *gin.Context) {
  fmt.Println("m1 ...in")
  c.Set("name", User{"枫枫", 21})
  c.Next()
  fmt.Println("m1 ...out")
}

func main() {
  router := gin.Default()

  router.Use(m10)
  router.GET("/", func(c *gin.Context) {
    fmt.Println("index ...in")
    name, _ := c.Get("name")
    user := name.(User)
    fmt.Println(user.Name, user.Age)
    c.JSON(200, gin.H{"msg": "index"})
  })

  router.Run(":8080")

}
```

### 7.分组路由注册中间件

可以指定哪一些分组下可以使用中间件

```go
func middle(c *gin.Context) {
  fmt.Println("middle ...in")
}

func main() {
  router := gin.Default()

  r := router.Group("/api").Use(middle)  // 可以链式，也可以直接r.Use(middle)
  r.GET("/index", func(c *gin.Context) {
    c.String(200, "index")
  })
  r.GET("/home", func(c *gin.Context) {
    c.String(200, "home")
  })

  router.Run(":8080")
}
```

当然，中间件还有一种写法，就是使用<font color='red'>函数加括号</font>的形式,可以<font color='red'>传参</font>：

```go

func middle(c *gin.Context) {
  fmt.Println("middle ...in")
}
func middle1() gin.HandlerFunc {
  // 这里的代码是程序一开始就会执行 
  return func(c *gin.Context) {
    // 这里是请求来了才会执行
    fmt.Println("middle1 ...inin")
  }
}

func main() {
  router := gin.Default()

  r := router.Group("/api").Use(middle, middle1()) //加()
  r.GET("/index", func(c *gin.Context) {
    c.String(200, "index")
  })
  r.GET("/home", func(c *gin.Context) {
    c.String(200, "home")
  })

  router.Run(":8080")
}
CopyErrorOK!
```

### 8. gin.Default

```go
func Default() *Engine {
  debugPrintWARNINGDefault()
  engine := New()
  engine.Use(Logger(), Recovery())
  return engine
}
CopyErrorOK!
```

> gin.Default()<font color='red'>默认使用</font>了Logger和Recovery中间件，其中：
>
> - Logger中间件将日志写入gin.DefaultWriter，即使配置了GIN_MODE=release。 
>
> - Recovery中间件会recover任何panic。如果有panic的话，会写入500响应码。 

如果不想使用上面两个默认的中间件，可以使用gin.New()新建一个没有任何默认中间件的路由。

使用gin.New，如果不指定日志，那么在控制台中就不会有日志显示

## 八、Bind绑定参数

gin中的bind可以很方便地将 前端传递 来的数据与<font color='red'>结构体</font>进行<font color='red'>参数绑定</font>以及<font color='red'>参数校验</font>

### 1、参数绑定

在使用此功能时，需要给结构体加上Tag 

> `json	form	uri	   xml	 yaml`

### 2、MustBind

一般不用，因为校验失败会<font color='red'>改变状态码</font>

### 3、ShouldBind

可以绑定json、query、param、yaml、xml

如果校验不通过会返回错误

#### 3.1 ShouldBindJSON

绑定Json参数

```go
package main

import "github.com/gin-gonic/gin"

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func main() {
	router := gin.Default()

	router.POST("/", func(ctx *gin.Context) {
		var userInfo UserInfo
		err := ctx.ShouldBindJSON(&userInfo)
		if err != nil {
			ctx.JSON(200, gin.H{"msg": "wrong"})
			return
		}
		ctx.JSON(200, userInfo)
	})

	router.Run(":8080")
}
```

#### 3.2 ShouldBindQuery

绑定查询参数   `/query?name=khs&age=32&sex=man`

- tag对应为form

```go
package main

import "github.com/gin-gonic/gin"

type UserInfo struct {
	Name string `json:"name" form:"name"`
	Age  int    `json:"age" form:"age"`
	Sex  string `json:"sex" form:"sex"`
}

func main() {
	router := gin.Default()

	router.POST("/query", func(ctx *gin.Context) {
		var userInfo UserInfo
		err := ctx.ShouldBindQuery(&userInfo)
		if err != nil {
			ctx.JSON(200, gin.H{"msg": "wrong"})
			return
		}
		ctx.JSON(200, userInfo)
	})

	router.Run(":8080")
}

```

#### 3.3 ShouldBindUri

绑定动态参数 	`/uri/khs/18/man`

- tag对应为uri

```go
package main

import "github.com/gin-gonic/gin"

type UserInfo struct {
	Name string `json:"name" form:"name" uri:"name"`
	Age  int    `json:"age" form:"age" uri:"age"`
	Sex  string `json:"sex" form:"sex" uri:"sex"`
}

func main() {
	router := gin.Default()

	router.POST("/uri/:name/:age/:sex", func(ctx *gin.Context) {
		var userInfo UserInfo
		err := ctx.ShouldBindUri(&userInfo)
		if err != nil {
			ctx.JSON(200, gin.H{"msg": "wrong"})
			return
		}
		ctx.JSON(200, userInfo)
	})

	router.Run(":8080")
}
```

#### 3.4 ShouldBind

> - 会根据请求头中的content-type去<font color='red'>自动绑定</font>
> - form-data的参数也用这个， Tag用from
> - 默认的Tag就是form
> - **绑定from-data、x-ww-form-urlencode**:
>

```go
package main

import "github.com/gin-gonic/gin"

type UserInfo struct {
	Name string `form:"name"`
	Age  int    `form:"age"`
	Sex  string `form:"sex"`
}

func main() {
	router := gin.Default()

	router.POST("/form", func(ctx *gin.Context) {
		var userInfo UserInfo
		err := ctx.ShouldBind(&userInfo)
		if err != nil {
			ctx.JSON(200, gin.H{"msg": "wrong"})
			return
		}
		ctx.JSON(200, userInfo)
	})

	router.Run(":8080")
}

```

## 九、Bind绑定器

需要使用参数验证功能,需要加<font color='red'>binding</font> Tag

### 1、常用验证器

```go
required: 必填字段 , 如 binding:"required" //不能为空且不能不填

//针对字符串的长度
min 最小长度, 如：binding:"min=5"
max 最大长度, 如：binding:"max=10"
len 长度, 如：binding:"len=6"

//针对数字大小
eq  等于, 如：binding:"eq=3"
ne  不等于, 如：binding:"ne=12"
gt 	大于, 如：binding:"gt=10"
gte 大于等于, 如：binding:"gt=10"
lt  小于, 如：binding:"lt=5"
lte 小于等于, 如：binding:"lte=5"

//针对同级字段， 如密码确认
eqfield 等于其他字段的值,如：binding:"eqfield=ComfiemPassword"
neqfield 不等于其他字段的值
-   忽略字段, 如：binding:"-"
```

### 	2、Gin内置验证器枚举

```go
//枚举， 只能为man or oman
oneof=man woman

//字符串
contains=khs    //包含khs的字符串
excludes=		//不包含
startswith=		//字符串前缀
endswith=		//字符串后缀

//数组
dive	//dive后面的验证就是针对数组中的每一个元素

//网络验证
ip
ipv4
ipv6
uri 	//在I(Identifier)是统一资源标示符，可以唯一标识一个资源
url		//在Locater，是统一资源定位符，提供找到该资源的确切路径
```



### 3、自定义错误信息

当验证不通过时，会给出错误的信息，但是原始的错误信息不太友好，不利于用户查看

```go
router.POST("/mistake", func(ctx *gin.Context) {
		type User struct {
			Name string `json:"name" binding:"required" msg:"用户名校验失败"` //用户名
			Age  int    `json:"age" binding:"required"  msg:"请输入年龄"`   //年龄
		}
		var user User
		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			ctx.JSON(200, gin.H{"msg": GetValidMsg(err, &user)})
			return
		}
		ctx.JSON(200, gin.H{"msg": user})
	})
```

当出现错误时，就可以来获取出错字段上的msg

```go
// 获取结构体中的msg参数
func GetValidMsg(err error, obj any) string {
	//使用时，需要穿obj的指针
	getObj := reflect.TypeOf(obj)
	//1、将err接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		//2、断言成功,循环每一个错误信息
		for _, e := range errs {
			//根据报错字段名 获取结构体的具体字段
			if field, exist := getObj.Elem().FieldByName(e.Field()); exist {
				msg := field.Tag.Get("msg")
				return msg
			}
		}
	}
	return err.Error()
}
```

### 4、自定义验证器

1、注册验证器函数

```go
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("sign", signValid)
	}
```

2、编写函数

```go
//若用户名已在List中则校验失败
func signValid(fl validator.FieldLevel) bool {
	var nameList []string = []string{"khs", "khx", "hsh"}
	for _, nameStr := range nameList {
		name := fl.Field().Interface().(string)
		if name == nameStr {
			return false
		}
	}
	return true
}
```

3、使用

```go
type User struct {
	Name string `json:"name" binding:"required,sign" msg:"用户名校验失败"` //用户名
	Age  int    `json:"age" binding:"required"  msg:"请输入年龄"`        //年龄
}
```

## 十、文件上传和下载

### 01、文件上传

> 注意需要使用POST中`form-data`的文件上传机制进行上传文件

- #### 单文件

	```go
	//单文件上传
		router.POST("/upload", func(ctx *gin.Context) {
			file, _ := ctx.FormFile("file") //file即form-data中上传文件的key名
			fmt.Println(file.Filename)
			fmt.Println(file.Size / 1024)                    //Size单位是字节
			ctx.SaveUploadedFile(file, "../uploads/123.png") //文件保存
			ctx.JSON(200, gin.H{"msg": "上传成功"})
		})
	```

	服务端保存文件的几种方式 

	- `SaveUploadedFile`

	```go
	ctx.SaveUploadedFile(file,dst) //文件对象  文件路径
	```

	- `Create + Copy`

	> file.Open()的第一个返回值就是我们文件对象中的那个文件(只读的)，我们可以用这个去直接读取文件内容

	```go
			file, _ := ctx.FormFile("file") //file即form-data中上传文件的key名
			readerFile, _ := file.Open() 
			//创建一个文件
			writerFile, _ := os.Create("../uploads/1.png")
			defer writerFile.Close()
			//复制文件
			n, _ := io.Copy(writerFile, readerFile)
			fmt.Println(n)
	
			ctx.JSON(200, gin.H{"msg": "上传成功"})
	```

	- `读取上传的文件`

	```go
			file, _ := ctx.FormFile("file")
			readerFile, _ := file.Open()
			data, _ := io.ReadAll(readerFile)
			fmt.Println(string(data))
	```

- #### 多文件

```go
router.POST("/uploads", func(ctx *gin.Context) {
		form, _ := ctx.MultipartForm()
		files := form.File["upload[]"] //upload[]即form-data中上传文件的key名
		for _, file := range files {
			ctx.SaveUploadedFile(file, "../uploads/"+file.Filename)
		}
		ctx.JSON(200, gin.H{"msg": fmt.Sprintf("成功上传%d个文件", len(files))})
	})
```

### 02、文件下载

- **直接响应一个路径下的文件**

```go
ctx.File("../uploads/12.png")
```

​	有些响应，比如图片，浏览器会显示这个图片，而不是下载，所以我们需要使浏览器`唤起下载行为`

```go
		ctx.Header("Content-Type", "application/octet-stream")                //设置响应头为文件流
		ctx.Header("Content-Disposition", "attachment; filename="+"good.png") //指定下载的文件名
		ctx.Header("Content-Transfer-Encoding", "binary")                     //传输过程中的编码形式
		ctx.File("../uploads/12.png")
```

注意，文件下载浏览器可能有缓存，解决办法就是加查询参数

- 前后端模式下的文件下载

在前后端模式下，后端只需要响应一个文件数据，文件名和其他信息就写在请求头中:

```go
ctx.Header("fileName", "xxx.png")
ctx.Header("msg","文件下载成功")
ctx.File("../uploads/12.png")
```

前端写法:

```js
async downloadFile(row) {
   this.$http({
      method: 'post',
      url: 'file/upload',
      data:postData,
      responseType: "blob"
   }).then(res => {
      const _res = res.data
      let blob = new Blob([_res], {
            type: 'application/png'
          });
      let downloadElement = document.createElement("a");
      let href = window.URL.createObjectURL(blob); //创建下载的链接
      downloadElement.href = href;
      downloadElement.download = res.headers["fileName"]; //下载后文件名
      document.body.appendChild(downloadElement);
      downloadElement.click(); //点击下载
      document.body.removeChild(downloadElement); //下载完成移除元素
      window.URL.revokeObjectURL(href); //释放掉blob对象
    })}
```

## 十一、日志

> 1. 记录用户操作，猜测用户行为
> 2. 记录bug

### 1、Gin内置日志系统

#### 01 输出日志到log文件

```go
package main

import (
  "github.com/gin-gonic/gin"
  "io"
  "os"
)

func main() {
  // 输出到文件
  f, _ := os.Create("gin.log")
  //gin.DefaultWriter = io.MultiWriter(f)
  // 如果需要同时将日志写入文件和控制台，请使用以下代码。
  gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
  router := gin.Default()
  router.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{"msg": "/"})
  })
  router.Run()
}
```

#### 02 定义路由格式

启动gin，它会显示所有的路由，默认格式如下

> [GIN-debug] POST   /foo    --> main.main.func1 (3 handlers)
> [GIN-debug] GET    /bar    --> main.main.func2 (3 handlers)
> [GIN-debug] GET    /status --> main.main.func3 (3 handlers)

```go
//自定义路由格式：
gin.DebugPrintRouteFunc = func(
  httpMethod,
  absolutePath,
  handlerName string,
  nuHandlers int) {
  log.Printf(
    "[ feng ] %v %v %v %v\n",
    httpMethod,
    absolutePath,
    handlerName,
    nuHandlers,
  )
}
/*  输出如下
2022/12/11 14:10:28 [ feng ] GET / main.main.func3 3
2022/12/11 14:10:28 [ feng ] POST /index main.main.func4 3
2022/12/11 14:10:28 [ feng ] PUT /haha main.main.func5 3
2022/12/11 14:10:28 [ feng ] DELETE /home main.main.func6 3
*/
```

#### 03 查看路由

```go
router.Routes()  // 它会返回已注册的路由列表
```

#### 04 环境切换

![img](http://python.fengfengzhidao.com/pic/20221211142056.png)

如果不想看到这些debug日志，那么我们可以改为<font color='red'>release模式</font>

```go
gin.SetMode(gin.ReleaseMode)
router := gin.Default()
```

#### 05 修改log的显示

默认的是这样的

```go
[GIN] 2022/12/11 - 14:22:00 | 200 |  0s |  127.0.0.1 | GET  "/"CopyErrorOK!
```

如果觉得不好看，我们可以<font color='red'>自定义log</font>：

```go
package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
)

func LoggerWithFormatter(params gin.LogFormatterParams) string {
  return fmt.Sprintf(
    "[ khs ] %s  | %d | \t %s | %s | %s \t  %s\n",
    params.TimeStamp.Format("2006/01/02 - 15:04:05"),
    params.StatusCode,  // 状态码
    params.ClientIP,  // 客户端ip
    params.Latency,  // 请求耗时
    params.Method,  // 请求方法
    params.Path,  // 路径
  )
}
func main() {
  router := gin.New()
  router.Use(gin.LoggerWithFormatter(LoggerWithFormatter))
  router.Run()
}
```

也可以这样

```go
func main() {
  router := gin.New()
  router.Use(gin.LoggerWithConfig(
      gin.LoggerConfig{Formatter: LoggerWithFormatter},
    ),
  )
  router.Run()
}
```

但是你会发现自己这样输出之后，没有颜色了，不太好看，我们可以<font color='red'>输出有颜色的log</font>

```go
func LoggerWithFormatter(params gin.LogFormatterParams) string {
  var statusColor, methodColor, resetColor string
  statusColor = params.StatusCodeColor()
  methodColor = params.MethodColor()
  resetColor = params.ResetColor()
  return fmt.Sprintf(
    "[ feng ] %s  | %s %d  %s | \t %s | %s | %s %-7s %s \t  %s\n",
    params.TimeStamp.Format("2006/01/02 - 15:04:05"),
    statusColor, params.StatusCode, resetColor,
    params.ClientIP,
    params.Latency,
    methodColor, params.Method, resetColor,
    params.Path,
  )
}
```

### 2、Logrus
