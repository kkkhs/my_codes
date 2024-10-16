package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

// 自定义一个中间件
func myHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//通过自定义的中间件，设置的值，在后续处理只要调用了这个中间件的都可以拿到这里的参数
		ctx.Set("usersesion", "userid-1") // key value

		ctx.Next() //放行

		ctx.Abort() //阻断
	}
}

type XmlUser struct {
	Id   int64  `xml:"id"`
	Name string `xml:"name"`
}

func main() {
	//创建一个服务
	r := gin.Default()
	r.Use(favicon.New("./123.ico"))

	//加载静态页面
	r.LoadHTMLGlob("templates/*") //全局加载
	// r.LoadHTMLFiles("templates/index.html")

	//加载资源文件
	r.Static("/static", "./static")

	//响应一个页面给前端
	r.GET("/index", func(ctx *gin.Context) {
		//ctx.Json() json数据
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "go后台传递的数据",
		})
	})

	//接收前端传送过来的参数
	// 1、user?userid=xxx&username=khs
	r.GET("/user/info", myHandler(), func(ctx *gin.Context) { //指定使用中间件

		//取出中间件的值
		usersesion := ctx.MustGet("usersesion").(string)
		log.Println("=============>", usersesion)

		userid := ctx.Query("userid")
		username := ctx.Query("username")
		ctx.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})
	})

	// 2、 /user/info/1/khs
	r.GET("user/info/:userid/:username", func(ctx *gin.Context) {
		userid := ctx.Param("userid")
		username := ctx.Param("username")
		ctx.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})
	})

	//前端给后端传递 json：[]byte---->数据
	r.POST("/json", func(ctx *gin.Context) {
		//request.body
		data, _ := ctx.GetRawData()

		var m map[string]interface{}
		_ = json.Unmarshal(data, &m)
		ctx.JSON(http.StatusOK, m)

	})

	r.POST("/user/add", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")

		ctx.JSON(http.StatusOK, gin.H{
			"msg":      "ok",
			"username": username,
			"password": password,
		})
	})

	//路由
	r.GET("/test", func(ctx *gin.Context) {
		//重定向
		ctx.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")

	})

	//404	Noroute
	r.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404.html", nil)
	})

	//路由组 /user
	usergroup := r.Group("/user")
	{
		usergroup.GET("/add")
		usergroup.POST("/login")
		usergroup.POST("/logout")
	}
	order := r.Group("/order")
	{
		order.GET("/add")
		order.DELETE("/delete")
	}

	r.GET("/user/save", func(ctx *gin.Context) {
		ctx.Header("test", "headertest")
	})

	r.GET("/cookie", func(ctx *gin.Context) {
		ctx.SetCookie("MyCookie", "khsCookie", 3600, "/", "localhost", false, true)
	})

	r.GET("/read", func(ctx *gin.Context) {
		// 根据cookie名字读取cookie值
		data, err := ctx.Cookie("MyCookie")
		if err != nil { //未拿到cookie
			ctx.String(http.StatusOK, "Not Found!")
			return
		}
		// 直接返回cookie值
		ctx.String(http.StatusOK, data)
	})

	r.GET("/del", func(ctx *gin.Context) {
		ctx.SetCookie("MyCookie", "khsCookie", -1, "/", "localhost", false, true)
		ctx.String(http.StatusOK, "Cookie已删除")
	})

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

	//服务器端口
	r.Run(":8080") //默认为8080
}
