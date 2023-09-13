/**
@author : ikkk
@date   : 2023/9/5
@text   : nil
**/

package gee

//
//import (
//	"fmt"
//	"net/http"
//	"testing"
//)
//
//func TestHttp_Start(t *testing.T) {
//	h := NewHTTP()
//	h.GET("/:get", func(c *Context) {
//		fmt.Println("get动态")
//		c.response.WriteHeader(200)
//		c.response.Write([]byte("动态请求" + c.params["get"]))
//	})
//	h.GET("/ok", func(c *Context) {
//		fmt.Println("get静态")
//		c.response.WriteHeader(200)
//		c.response.Write([]byte("ok静态"))
//	})
//	h.GET("/ok/:get/*file", func(c *Context) {
//		fmt.Println("get动态")
//		c.response.WriteHeader(200)
//		c.response.Write([]byte("动态请求" + c.params["get"] + " " + c.Param("file")))
//	})
//	err := h.Run(":8080")
//	if err != nil {
//		panic(err)
//	}
//}
//
//func TestHTTP_Start1(t *testing.T) {
//	h := NewHTTP()
//	h.GET("/study/login", func(c *Context) {
//		c.String(http.StatusOK, fmt.Sprintf("静态路由: %s", c.Pattern))
//	})
//	h.GET("/study/:text", func(c *Context) {
//		text := c.Param("text")
//		c.String(http.StatusOK, fmt.Sprintf("动态路由参数: %s", text))
//	})
//	h.GET("/star/*text", func(c *Context) {
//		file := c.Param("text")
//		c.String(http.StatusOK, fmt.Sprintf("通配符路由参数: %s", file))
//	})
//	h.GET("/", func(c *Context) {
//		a := c.Query("a")
//		b := c.Query("b")
//		c.JSON(http.StatusOK, H{
//			"a": a,
//			"b": b,
//		})
//	})
//	h.GET("/html", func(c *Context) {
//		c.HTML(http.StatusOK, `<h1>从零到一实现一个Web框架</h1>`)
//	})
//	h.POST("/login", func(c *Context) {
//		mp := map[string]string{}
//		err := c.BindJSON(&mp)
//		if err != nil {
//			c.JSON(http.StatusNotFound, err.Error())
//			return
//		}
//		c.JSON(http.StatusOK, H{
//			"code":     http.StatusOK,
//			"msg":      "ok",
//			"account":  mp["account"],
//			"password": mp["password"],
//			"a":        mp["a"],
//		})
//	})
//	err := h.Run(":8080")
//	if err != nil {
//		panic(err)
//	}
//
//}
