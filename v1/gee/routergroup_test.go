/**
 @author : ikkk
 @date   : 2023/9/12
 @text   : nil
**/

package gee

import (
	"fmt"
	"testing"
)

func TestRouterGroupUse(t *testing.T) {
	r := New()
	r.GET("", func(c *Context) {
		fmt.Println("1")
		c.String(200, "/")
	}, func(c *Context) {
		fmt.Println("2")
		c.String(200, "2")
	})
	group := r.Group("group")
	group.Use(func(c *Context) {
		fmt.Println(1)
	}, func(c *Context) {
		fmt.Println(2)
	})
	group.GET("/hello", func(c *Context) {
		c.String(200, "hello")
	})
	{
		v1 := group.Group("/v1")
		v1.Use(func(c *Context) {
			fmt.Println(3)
		}, func(c *Context) {
			fmt.Println(4)
		})
		v1.GET("/ok", func(c *Context) {
			c.String(200, "0")
		}, func(c *Context) {
			c.String(200, "ok")
		})
	}
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}

}
