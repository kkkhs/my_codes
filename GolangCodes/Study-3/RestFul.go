package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ArticleModel struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

// 文章列表页面
func _getList(c *gin.Context) {
	//包含搜索，分页，关键词
	articleList := []ArticleModel{
		{Title: "GO语言入门", Content: "这篇文章是《Gin》"},
		{Title: "Python语言入门", Content: "这篇文章是《Django》"},
		{Title: "Java语言入门", Content: "这篇文章是《Spring》"},
	}

	c.JSON(200, Response{0, articleList, "成功"})
}

// 文章详情页面
func _getDetail(c *gin.Context) {
	//获取parem中的id
	fmt.Println("id =", c.Param("id"))
	article := ArticleModel{
		Title: "GO语言入门", Content: "这篇文章是《Gin》",
	}

	c.JSON(200, Response{0, article, "成功"})
}

// 创建文章接口
func _create(c *gin.Context) {
	//接收前端传递来的Json数据
	data, err := c.GetRawData()
	if err != nil {
		panic(err)
	}

	var m map[string]interface{} //创建m接受Json反序列化结果
	_ = json.Unmarshal(data, &m)

	c.JSON(200, Response{0, m, "成功"})
}

// 编辑文章
func _uptate(c *gin.Context) {
	//获取parem中的id 找到数据库中原数据
	fmt.Println("id =", c.Param("id"))
	//找到数据库中原数据....

	//接收前端传递来的Json新数据
	data, err := c.GetRawData()
	if err != nil {
		panic(err)
	}
	var m map[string]interface{} //创建m接受Json反序列化结果
	_ = json.Unmarshal(data, &m)
	//修改数据库中原数据....

	c.JSON(200, Response{0, m, "修改成功"})
}

func _delete(c *gin.Context) {
	//获取parem中的id 找到数据库中原数据并删除
	fmt.Println("id =", c.Param("id"))

	c.JSON(200, Response{0, map[string]string{}, "删除成功"})
}

func main() {
	router := gin.Default()
	router.GET("/articles", _getList)       //文章列表
	router.GET("/articles/:id", _getDetail) //文章详情
	router.POST("/articles", _create)       //添加文字
	router.PUT("/articles/:id", _uptate)    //编辑文章
	router.DELETE("/articles/:id", _delete) //删除文章

	router.Run(":8080")
}
