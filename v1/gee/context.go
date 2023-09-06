/**
 @author : ikkk
 @date   : 2023/9/5
 @text   : nil
**/

package gee

import "net/http"

type Context struct {
	response http.ResponseWriter
	request  *http.Request
	Method   string
	Pattern  string
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		response: w,
		request:  r,
		Method:   r.Method,
		Pattern:  r.URL.Path,
	}
}
