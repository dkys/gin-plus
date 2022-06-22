package gplus

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strings"
)

const (
	Get     = "Get"
	Head    = "Head"
	Post    = "Post"
	Put     = "Put"
	Patch   = "Patch" // RFC 5789
	Delete  = "Delete"
	Connect = "Connect"
	Options = "Options"
	Trace   = "Trace"
)

type Engine struct {
	*gin.Engine
}

func (e *Engine) Handle(relativePath string, dest interface{}) {
	//elog.Debug("||||||||")
}

func (e *Engine) GET(relativePath string, handlers ...gin.HandlerFunc) {
	e.Engine.GET(relativePath, handlers...)
}

func (e *Engine) Handles(relativePath string, dest ...interface{}) {
	if len(dest) > 0 {
		for _, item := range dest {
			ty := reflect.TypeOf(item)
			switch ty.Kind() {
			case reflect.Ptr:
				v := reflect.ValueOf(item)
				n := ty.NumMethod()
				var path strings.Builder
				for i := 0; i < n; i++ {
					path.WriteString(relativePath)
					path.WriteByte('/')
					httpMethod, funcName := GetMethod(ty.Method(i).Name)
					path.WriteString(strings.ToLower(funcName))
					e.RouterGroup.Handle(httpMethod, path.String(), v.Method(i).Interface().(func(ctx *gin.Context)))
					path.Reset()
				}
			}
		}
	}

}

func GetMethod(name string) (string, string) {
	switch {
	case strings.HasPrefix(name, Get):
		return http.MethodGet, strings.Replace(name, Get, "", 1)
	case strings.HasPrefix(name, Post):
		return http.MethodPost, strings.Replace(name, Post, "", 1)
	case strings.HasPrefix(name, Put):
		return http.MethodPut, strings.Replace(name, Put, "", 1)
	case strings.HasPrefix(name, Delete):
		return http.MethodDelete, strings.Replace(name, Delete, "", 1)
	case strings.HasPrefix(name, Options):
		return http.MethodOptions, strings.Replace(name, Options, "", 1)
	default:
		return http.MethodGet, name

	}

}

func Default() *Engine {
	engine := New()
	engine.Use(gin.Logger(), gin.Recovery())
	return engine
}

func New() *Engine {
	return &Engine{gin.New()}
}
