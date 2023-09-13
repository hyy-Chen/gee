/**
 @author : ikkk
 @date   : 2023/9/12
 @text   : nil
**/

package gee

import "testing"

func TestRouterGroupUse(t *testing.T) {
	r := New()
	r.GET("/", func(c *Context) {
		c.String(200, "/")
	})
	group := r.Group("/")
	group.GET("/hello", func(c *Context) {
		c.String(200, "hello")
	})
	group.GET("/hello/", func(c *Context) {
		c.String(200, "hello")
	})
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}

}
