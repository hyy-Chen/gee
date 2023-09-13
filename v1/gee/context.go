/**
 @author : ikkk
 @date   : 2023/9/5
 @text   : nil
**/

package gee

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// H 提供一个map,便于使用
type H map[string]interface{}

// Context 上下文
type Context struct {
	response http.ResponseWriter
	request  *http.Request
	Method   string
	// 存储路径
	Pattern string
	// 存储查询参数
	params map[string]string
	// 存储处理责任链
	handlers HandlersChain

	index int

	// cacheQuery 内部维护一份查询参数数据
	cacheQuery url.Values
	// cacheBody 内部维护一份请求体参数
	cacheBody io.ReadCloser

	// 状态码
	status int
	// 响应头
	header map[string]string
	// 响应体
	data []byte
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		response: w,
		request:  r,
		Method:   r.Method,
		Pattern:  r.URL.Path,
		index:    -1,
		params:   map[string]string{},
		header:   map[string]string{},
		data:     []byte{},
	}
}

func (c *Context) Next() {
	c.index++
	n := len(c.handlers)
	for c.index < n {
		c.handlers[c.index](c)
		c.index++
	}
}

// Param 获取请求参数
func (c *Context) Param(key string) string {
	return c.params[key]
}

// Query 获取查询参数
func (c *Context) Query(key string) string {
	if c.cacheQuery == nil {
		c.cacheQuery = c.request.URL.Query()
	}
	value := c.cacheQuery.Get(key)
	return value
}

// Form 获取请求体内参数
func (c *Context) Form(key string) string {
	if c.cacheBody == nil {
		c.cacheBody = c.request.Body
	}
	return c.request.FormValue(key)
}

// BindJSON 解析json信息请求参数
func (c *Context) BindJSON(val any) error {
	if c.cacheBody == nil {
		c.cacheBody = c.request.Body
	}
	decoder := json.NewDecoder(c.cacheBody)
	return decoder.Decode(val)
}

// SetStatusCode 设置状态码
func (c *Context) SetStatusCode(code int) {
	c.status = code
}

// SetHeader 设置响应头部
func (c *Context) SetHeader(key string, value string) {
	if value == "" {
		c.DelHeader(key)
		return
	}
	c.header[key] = value
}

func (c *Context) DelHeader(key string) {
	delete(c.header, key)
}

// SetData 设置响应体
func (c *Context) SetData(data []byte) {
	c.data = data
}

// JSON serializes the given struct as JSON into the response body.
// It also sets the Content-Type as "application/json".
func (c *Context) JSON(code int, data any) {
	c.SetStatusCode(code)
	c.SetHeader("Content-Type", "application/json")
	if res, err := json.Marshal(data); err != nil {
		c.SetStatusCode(http.StatusInternalServerError)
		c.DelHeader("Content-Type")
		panic(err)
	} else {
		c.SetData(res)
	}
}

// HTML 设置HTML响应数据
func (c *Context) HTML(code int, html string) {
	c.SetStatusCode(code)
	c.SetHeader("Content-Type", "text/html")
	c.SetData([]byte(html))
}

// 设置字符串响应数据
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.SetStatusCode(code)
	c.SetData([]byte(fmt.Sprintf(format, values...)))
}

// Data 设置Data数据
func (c *Context) Data(code int, data []byte) {
	c.SetStatusCode(code)
	c.SetData(data)
}

// 临时测试方法
func (c *Context) flashDataToResponse() {
	c.response.WriteHeader(c.status)
	for key, value := range c.header {
		c.response.Header().Set(key, value)
	}
	c.response.Write(c.data)
}
