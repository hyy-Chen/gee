/**
 @author : ikkk
 @date   : 2023/9/5
 @text   : nil
**/

package gee

import (
	"fmt"
	"testing"
)

func TestHttp_Start(t *testing.T) {
	h := NewHTTP()
	h.GET("/", func(c *Context) {
		fmt.Println("get")
		c.response.WriteHeader(200)
		c.response.Write([]byte("ok"))
	})
	h.GET("/", func(c *Context) {
		fmt.Println("post")
		c.response.WriteHeader(200)
		c.response.Write([]byte("ok请求成功"))
	})
	err := h.Start(":8080")
	if err != nil {
		panic(err)
	}

}
