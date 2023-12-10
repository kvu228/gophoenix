package engine

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PhoenixContext struct {
	res         http.ResponseWriter
	req         *http.Request
	Params      map[string]any
	Path        string
	Method      string
	index       int // -1 default
	middlewares []HandlerFunc
}

func (c *PhoenixContext) JSON(httpStatus int, message any) {
	c.setStatus(httpStatus)
	json.NewEncoder(c.res).Encode(message)
}

func (c *PhoenixContext) setStatus(httpStatus int) {
	c.res.WriteHeader(httpStatus)
}

func (c *PhoenixContext) GetParam(key string) any {
	if value, exist := c.Params[key]; exist {
		return value
	}
	return nil
}

func (c *PhoenixContext) Halt() {
	fmt.Println("Halted context")
}

func (c *PhoenixContext) Next() {
	c.index += 1
	if c.index < len(c.middlewares) {
		nextHandler := c.middlewares[c.index]
		nextHandler(c)
	}
}


// Cấp phát new context
func NewPhoenixContext(res http.ResponseWriter, req *http.Request) *PhoenixContext {
	ctx := &PhoenixContext{
		res:    res,
		req:    req,
		Params: make(map[string]any),
		index:  -1,
	}
	ctx.Path = req.URL.Path
	ctx.Method = req.Method
	return ctx
}
